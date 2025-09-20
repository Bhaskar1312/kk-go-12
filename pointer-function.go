package main

import "fmt"

func modify(s *string) {
	*s = "world"
}

func modify2(numbers ...int) {
	for i := range numbers {
			numbers[i] -= 5
	}
}

func main() {
	a := "hello"
	fmt.Println(a)
	modify(&a)
	fmt.Println(a)


	arr := []int{10, 20, 30}
	fmt.Println(arr)
	modify2(arr...)
	fmt.Println(arr)
}