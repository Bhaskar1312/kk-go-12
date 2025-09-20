package main

import "fmt"

func main() {
        arr := [5]int{10, 20, 90, 70, 60}
        slice := arr[:3]
        fmt.Println(cap(slice))
        new_slice := append(slice, 100, 200, 300)
        fmt.Println(cap(new_slice))
		new_slice[0] = 0
		fmt.Println(arr)
		fmt.Println(slice)
		fmt.Println(new_slice)

		// Create a slice with length 5 and capacity 20 to demonstrate difference between length and capacity
		slice2 := make([]int, 5, 20)
		slice2 = append(slice2, 1, 2)
		slice2 = append(slice2, slice...)
		fmt.Println(slice2)

		arr4 := []int{10, 20, 90, 70, 60}
        slice4 := make([]int, 10)
        num := copy(slice4, arr4)
        fmt.Println(slice4)
        fmt.Println(num)
}