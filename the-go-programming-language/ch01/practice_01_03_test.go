package test

import (
	"fmt"
	"os"
	"testing"
)

func echo1() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo2() {
	for idx, arg := range os.Args[1:] {
		fmt.Printf("idx: %v, value: %v\n", idx, arg)
	}
}

func BenchmarkEcho1(b *testing.B) {
	b.ResetTimer()
	os.Args = []string{"echo1", "hello", "world"}

	for i := 0; i < b.N; i++ {
		echo1()
	}
}

func BenchmarkEcho2(b *testing.B) {
	b.ResetTimer()
	os.Args = []string{"echo2", "hello", "world"}

	for i := 0; i < b.N; i++ {
		echo2()
	}
}
