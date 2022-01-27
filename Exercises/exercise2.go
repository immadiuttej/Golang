package main

import "fmt"

func main() {
	var x int = 42
	var y string = "Janes Bond"
	var z bool = true

	s := fmt.Sprintf("x = %v, y = %v, z = %v", x, y, z)

	fmt.Println(s)
}
