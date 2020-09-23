package main

import (
	"log"
	"sync"
	"time"

	"github.com/gurumee92/go-in-action/ch07/work"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

type namePrinter struct {
	name string
}

func (p *namePrinter) Task() {
	log.Println(p.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(10)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	p.Shutdown()
}
