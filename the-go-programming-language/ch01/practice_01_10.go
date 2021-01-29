package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	defer resp.Body.Close()
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func main() {
	f, err := os.Create("practice_10.txt")

	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Fprintln(f, <-ch)
	}

	fmt.Fprintf(f, "%.2fs elapsed\n", time.Since(start).Seconds())
}
