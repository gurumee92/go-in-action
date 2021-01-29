package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/golang-collections/collections/set"
)

func main() {
	counts := make(map[string]*set.Set)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			log.Fatalln(err)
		}

		for _, line := range strings.Split(string(data), "\n") {
			if counts[line] == nil {
				counts[line] = set.New()
			}

			counts[line].Insert(filename)
		}
	}

	for line, sets := range counts {
		if sets.Len() > 1 {
			fmt.Printf("%v: ", line)
			sets.Do(func(i interface{}) {
				fmt.Printf("%v ", i)
			})
			fmt.Println()
		}
	}
}
