package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

// 여러 개의 고루틴을 안전하게 공유하기 위한 리소스 집합을 표현
// 이 풀을 관리하기 위한 리소스는 io.Closer를 관리해야 한다.
type Pool struct {
	// 고루틴이 풀을 접근할 때, 안전하게 작업하기 위한 뮤텍스
	m sync.Mutex
	// io.Closer 채널, 버퍼가 있는 채널로 생성 리소스 공유가 목적
	resources chan io.Closer
	// 함수 타입.. 풀에 리소스 요청이 들어올 때 새로운 리소스 생성
	factory func() (io.Closer, error)
	// 풀이 닫혀있는지 확인하는 변수
	closed bool
}

var ErrPoolClosed = errors.New("Pool is closed.")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("Pool's size too small")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		log.Println("Resource acquire:", "Shared Resource")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Resource acquire:", "New Resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.resources <- r:
		log.Println("Resource Return: ", "Returned Resource Queue")
	default:
		log.Println("Resource Return: ", "Release Resource")
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)

	for r := range p.resources {
		r.Close()
	}
}
