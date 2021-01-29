package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gurumee92/golang-studies/the-go-programming-language/ch02/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\b", err)
			os.Exit(1)
		}

		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%v = %v, %v = %v\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
