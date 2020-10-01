package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	defer r.Body.Close()
	file, err := os.Create(os.Args[2])

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	dest := io.MultiWriter(os.Stdout, file)
	io.Copy(dest, r.Body)

	err = os.Remove(os.Args[2])

	if err != nil {
		log.Fatalln(err)
	}
}
