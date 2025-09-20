package main

import "fmt"

func main() {
	ascii := make(map[string]int)
	fmt.Println(ascii)

	ascii["A"] = 65
	ascii["F"] = 70
    ascii["K"] = 75
	fmt.Println(ascii)

	delete(ascii, "F")
	fmt.Println(ascii)

	ascii_codes := make(map[string]int, 10)
        ascii_codes["A"] = 65
        ascii_codes["F"] = 70
        ascii_codes["K"] = 75
        fmt.Println(len(ascii_codes))
        ascii_codes = make(map[string]int)
        ascii_codes["U"] = 85
        fmt.Println(len(ascii_codes))
}