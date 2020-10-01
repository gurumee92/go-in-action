package main

import "log"

func init() {
	log.SetPrefix("Tracing: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("Message")
	log.Fatalln("Critical Error Message")
	log.Panicln("Panic Message")
}
