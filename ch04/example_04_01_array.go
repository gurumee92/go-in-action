package main

import "fmt"

func main() {
	fmt.Println("Go에서 배열은, [숫자]타입{ 값들 }로 만들 수 있습니다.")
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println("생성 후", a)
	fmt.Println("생성 후", a[2])

	fmt.Println("\n배열 원소 접근은 [] 연산자로 합니다.")
	a[1] = 3
	fmt.Println("a[1] 변경 후", a)
	fmt.Println("a[1] 변경 후", a[1])

	fmt.Println("\n배열 순회는 다음과 같이 할 수 있습니다.")
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	fmt.Println("\n배열 순회는 다음과 같이 이쁘게^^ 할 수 있습니다. idx는 _로 생략 가능합니다.")
	for idx, num := range a {
		fmt.Println(idx, num)
	}

	fmt.Println("\nGo에서 배열은 값으로 취급됩니다.")
	fmt.Println("함수 호출 전", a)
	SetIndexCallByValue(a)
	fmt.Println("함수 호출 후", a)

	fmt.Println("\n포인터로 넘기면, 메모리도 아끼고 좋지요")
	fmt.Println("함수 호출 전", a)
	SetIndexCallByPoint(&a)
	fmt.Println("함수 호출 후", a)

	fmt.Println("\n다차원 배열도 만들 수 있어요 다음 처럼 말이죠.")
	b := [3][3]int{
		{0, 0, 0},
		{2: 1},
		{1},
	}
	fmt.Println(b)
}

func SetIndexCallByValue(a [5]int) {
	a[2] = 6
}

func SetIndexCallByPoint(a *[5]int) {
	a[2] = 6
}
