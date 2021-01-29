package main

import "fmt"

func main() {
	fmt.Println("맵은 map[타입]타입 으로 생성할 수 있습니다. 방법은 크게 2가지")
	m1 := make(map[string]string)
	fmt.Println(m1)

	m2 := map[string]string{
		"RED":    "#da1337",
		"ORANGE": "#e92a22",
		"GREEN":  "#a3ff47",
	}
	fmt.Println(m2)

	fmt.Println("\n원소 추가는 다음과 같이 할 수 있어요.")
	m2["BLACK"] = "#ffffff"
	fmt.Println(m2)

	fmt.Println("\n원소 접근은 이렇게 할 수 있어요. (값,존재 여부)를 반환하지요.")
	v, isExist := m2["WHITE"]
	fmt.Println(v, isExist)

	v, isExist = m2["BLACK"]
	fmt.Println(v, isExist)

	fmt.Println("\n삭제는 요렇게 할 수 있어요.")
	delete(m2, "BLACK")
	fmt.Println("삭제 후: ", m2)

	fmt.Println("\n순회는 요렇게 합니다. 근데 순서는 그때마다 달라요.")
	for k, v := range m2 {
		fmt.Println(k, v)
	}

	fmt.Println("\n슬라이스처럼 참조로 취급받습니다.")
	AddElem(m2)
	fmt.Println("함수 호출 후: ", m2)
}

func AddElem(m map[string]string) {
	m["WHITE"] = "#000000"
}
