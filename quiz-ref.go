package main

import "fmt"

func calculate(a int, b int) []int {
	// your code goes here
	arr := [4]int{a + b, a - b, a * b, a / b}
	return arr[:]

}

func main() {
	fmt.Println(calculate(20, 10))
	fmt.Println(calculate(700, 70))
}
