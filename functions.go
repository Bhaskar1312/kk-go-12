package main

import "fmt"

func calcSquare(numbers []int) ([]int, bool) {
	squares := []int{}
	for _, v := range numbers {
		squares = append(squares, v*v)
	}
	return squares, true

}

func main() {
	nums := [3]int{10, 20, 15}
	fmt.Println(calcSquare(nums[:]))
}
