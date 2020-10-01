package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalln("에러 로그 파일을 열 수 없습니다.", err)
	}

	Trace = log.New(ioutil.Discard, "Tracing: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("일반적인 로그 메시지")
	Info.Println("특별한 정보를 위한 로그 메시지")
	Warning.Println("경고성 로그 메시지")
	Error.Println("에러 로그 메시지")
}
