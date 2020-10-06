package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberOfGorouines = 4
	taskLoad          = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	// 채널생성, 10개의 값을 보관 가능
	tasks := make(chan string, taskLoad)
	// 4개의 고루틴이 실행됨
	wg.Add(numberOfGorouines)

	// 4개의 고루틴 실행
	for gr := 1; gr <= numberOfGorouines; gr++ {
		go worker(tasks, gr)
	}

	// 10개의 작업을 채널에 송신
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("작업: %d", post)
	}

	close(tasks)
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		// 채널에서 값을 꺼냄.
		task, ok := <-tasks

		// 채널이 닫혔다면, 고루틴 종료
		if !ok {
			fmt.Printf("작업자 :%d : 종료합니다.\n", worker)
			return
		}

		// 작업 실행
		fmt.Printf("작업자 :%d : 작업 시작 :%s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)

		fmt.Printf("작업자 :%d : 작업 완료 :%s\n", worker, task)
	}
}
