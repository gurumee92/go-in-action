package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 버퍼 없는 채널 생성
	court := make(chan int)
	// 고루틴 당 하나씩 총 2개의 카운터
	wg.Add(2)

	go player("나딜", court)
	go player("죠코비치", court)

	court <- 1
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court

		if !ok {
			fmt.Printf("%s 선수가 승리했습니다.\n", name)
			return
		}

		n := rand.Intn(100)

		if n%13 == 0 {
			fmt.Printf("%s 선수가 공을 받아치지 못했습니다.\n", name)
			//채널을 닫는다.
			close(court)
			return
		}

		fmt.Printf("%s 선수가 %d 번째 공을 받아쳤습니다.\n", name, ball)
		ball++

		court <- ball
	}
}
