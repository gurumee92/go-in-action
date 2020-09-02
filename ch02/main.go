package main

import (
	// 표준 패키지, log, os를 가져옴
	"log"
	"os"

	// 현재 모듈 matchers, search 가져옴.
	"github.com/gurumee92/go-in-action/ch02/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Sherlock Homes")
}
