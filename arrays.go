package main

import "fmt"

func main() {
	var arr2 [5]string = [5]string {}
	fmt.Println(arr2)
	fmt.Println(len(arr2))

	arr := [5]bool{true, true, true, true}
        for i := 0; i < len(arr); i++ {
                if arr[i] {
                        fmt.Println(i)
                }
        }

		for index, element := range arr {
                fmt.Println(index, "->", element)
        }
}