package main

import "fmt"

type Point struct {
	X int
	Y int
}

type MyInt int

func main() {
	fmt.Println("구조체 정의 방법 2가지")
	p1 := Point{
		X: 1,
		Y: 2,
	}
	p2 := Point{4, 1}
	fmt.Println(p1, p2)

	fmt.Println("타입 재 정의")
	var m1 MyInt
	m1 = 5
	m2 := 5
	fmt.Println(m1, m2)
}
