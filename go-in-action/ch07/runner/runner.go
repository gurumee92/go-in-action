package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner is ..
type Runner struct {
	// OS로부터 전달되는 인터럽트를 수신하기 위한 채널
	interrupt chan os.Signal
	// 처리가 종료되었음을 알리는 채널
	complete chan error
	// 지정된 시간이 초과했음을 알리는 채널
	timeout <-chan time.Time
	// 태스크를 위한 슬라이스
	tasks []func(int)
}

// ErrTimeout is ...
var ErrTimeout = errors.New("시간을 초과했습니다.")

// ErrInterrupt is ...
var ErrInterrupt = errors.New("OS 인터럽트 신호를 수신했습니다.")

// New Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add is Runner's method
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.goInterrupt() {
			return ErrInterrupt
		}

		task(id)
	}

	return nil
}

func (r *Runner) goInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
