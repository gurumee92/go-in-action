package main

import (
	"log"
	"os"
	"time"

	"github.com/gurumee92/go-in-action/ch07/runner"
)

const timeout = 3 * time.Second

func main() {
	log.Println("작업 시작!")

	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrInterrupt:
			log.Println("인터럽트 에러 발생")
			os.Exit(2)
		case runner.ErrTimeout:
			log.Println("작업 시간 초과")
			os.Exit(1)
		}
	}

	log.Println("작업 종료!")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Process - JOB #%d Start.\n", id)
		time.Sleep(time.Duration(id) * time.Second)
		log.Printf("Process - JOB #%d Finish.\n", id)
	}
}
