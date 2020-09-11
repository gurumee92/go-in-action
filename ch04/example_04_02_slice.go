package main

import "fmt"

func main() {
	fmt.Println("슬라이스는 가변 배열입니다.")
	fmt.Println("만드는 방법 1) []타입{값..} []에 숫자가 없어야 합니다.")
	s := []int{1, 2, 3}
	fmt.Println(s)
	fmt.Println("만드는 방법 2) make 함수를 써서 만들 수도 있어요. 1, 크기를 초기화할 수 있습니다.")
	s1 := make([]int, 2)
	fmt.Printf("길이 %v, 용량 %v\n", len(s1), cap(s1))

	fmt.Println("만드는 방법 3) make 함수를 써서 만들 수도 있어요. 2, 용량도 초기화가 가능합니다.")
	s2 := make([]int, 2, 3)
	fmt.Printf("길이 %v, 용량 %v\n", len(s2), cap(s2))

	fmt.Println("\n원소 추가는 이렇게 할 수 있습니다.")
	fmt.Printf("추가 전 %v\n", s2)
	s2 = append(s2, 5)
	fmt.Printf("추가 후 %v\n", s2)

	fmt.Println("\n슬라이스는 잘라낼 수가 있어요")
	c1 := []string{"Red", "Blue", "Green", "Gray", "Pink"}
	c2 := c1[2:4]
	fmt.Printf("c1: %v, c2: %v\n", c1, c2)

	fmt.Println("\n원소를 변경하는 요령은 배열과 같아요.")
	fmt.Printf("원소 변경 전 c1: %v, c2: %v\n", c1, c2)
	c2[1] = "바뀌었엉"
	fmt.Printf("원소 변경 후 c1: %v, c2: %v\n", c1, c2)

	fmt.Println("\n슬라이스는 용량 크기까지 초기화해서 잘라낼 수가 있어요")
	c3 := c1[2:4:4]
	fmt.Printf("c1: %v\nc2: %v\nc3: %v\n", c1, c2, c3)
	fmt.Println("뭐가 다르냐면, 용량이 찼을 때 내부 배열이 바뀝니당")
	c2 = append(c2, "추가됐엄")
	c3 = append(c3, "추가됐엄2")
	fmt.Printf("c1: %v\nc2: %v\nc3: %v\n", c1, c2, c3)

	fmt.Println("\n다차원 슬라이스도 가능합니다.")
	d := [][]int{
		{},
		{1, 2},
	}
	fmt.Println(d)

	fmt.Println("\n슬라이스는 배열과 달리 참조로 취급받습니다.")
	fmt.Println("함수로 전달 받으면, 주소가 전달됩니다.")
	fmt.Println("호출 전: ", c3)
	SetElem(c3)
	fmt.Println("호출 후: ", c3)
}

func SetElem(s []string) {
	s[2] = "함수 호출 후 변경됨."
}
