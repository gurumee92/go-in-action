package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("생성 후", a)
	fmt.Println("생성 후", a[2])

	a[1] = 3
	fmt.Println("a[1] 변경 후", a)
	fmt.Println("a[1] 변경 후", a[1])

	fmt.Println("배열 순회")
	for idx, num := range a {
		fmt.Println(idx, num)
	}

	fmt.Println("배열은 값입니다.")
	fmt.Println("함수 호출 전", a)
	SetIndexCallByValue(a)
	fmt.Println("함수 호출 후", a)

	fmt.Println("포인터로 넘기면, 메모리도 아끼고 좋지요")
	fmt.Println("함수 호출 전", a)
	SetIndexCallByPoint(&a)
	fmt.Println("함수 호출 후", a)
}

func SetIndexCallByValue(a [5]int) {
	a[2] = 6
}

func SetIndexCallByPoint(a *[5]int) {
	a[2] = 6
}
