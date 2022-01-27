package main

import "fmt"

func main() {
	num := 100
	fmt.Printf("num in decimal = %d, num in hex = %#x, num in binary = %b\n", num, num, num) // for hex, either %x or %#x can be used

	num2 := num << 1 //shifting 1 bit towards left

	fmt.Printf("num2 in decimal = %d, num2 in hex = %#x, num2 in binary = %b", num2, num2, num2) // for hex, either %x or %#x can be used

}
