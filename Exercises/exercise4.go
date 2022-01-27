package main

import "fmt"

type uttej int

var num1 uttej

func main() {

	num1 = 42
	num2 := int(num1)
	fmt.Println(num2)
	fmt.Printf("%T\n", num2)

}
