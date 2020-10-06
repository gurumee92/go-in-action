package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 스케줄러가 사용할 하나의 논리 프로세스 할당
	runtime.GOMAXPROCS(1)

	// wg는 프로그램 종료를 대기하기 위해 사용
	// 고루틴 개수만큼 더해준다.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("고루틴 시작!")

	go func() {
		// main 함수에 종료를 알리기 위해 Done 함수 호출
		defer wg.Done()

		// 소문자를 3번 출력한다.
		for count := 0; count < 3; count++ {
			time.Sleep(3 * time.Second)
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	go func() {
		// main 함수에 종료를 알리기 위해 Done 함수 호출
		defer wg.Done()

		// 대문자를 3번 출력한다.
		for count := 0; count < 3; count++ {
			time.Sleep(3 * time.Second)
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	fmt.Println("고루틴 대기 중")
	wg.Wait()

	fmt.Println("고루틴 끝~!")
}
