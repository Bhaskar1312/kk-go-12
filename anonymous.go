package main
import "fmt"

func main() {

	var (my_func = func(s string) {
		fmt.Println("Hey there,", s)
	})
	fmt.Println(my_func("heelo"))
}

// var (
//         cube = func(i int) string {
//                 c := i * i * i
//                 return strconv.Itoa(c)
//         }
// )

// func main() {
//         x := cube(8)
//         fmt.Printf("%T %v", x, x)
// }