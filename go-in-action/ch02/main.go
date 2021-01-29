package main

import (
	// 표준 패키지, log, os를 가져옴
	"log"
	"os"

	// 현재 모듈 matchers, search 가져옴.
	_ "github.com/gurumee92/golang-studies/go-in-action/ch02/matchers"
	"github.com/gurumee92/golang-studies/go-in-action/ch02/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("corona")
}
