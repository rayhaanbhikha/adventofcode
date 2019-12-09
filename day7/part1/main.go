package main

import "fmt"

func main() {
	x := 5
	changeNum(&x)
	fmt.Println(x)
}

func changeNum(x *int) {
	*x = 10
	fmt.Println("in func", x)

}
