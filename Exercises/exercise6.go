package main

import "fmt"

func main() {

	a := (42 == 42)
	b := (32 <= 16)
	c := (9 >= 22)
	d := (12 != 10)
	e := (10 > 10)
	f := (9 < 8)

	fmt.Printf("a = %v, b = %v, c = %v, d = %v, e = %v, f = %v\n", a, b, c, d, e, f)
}
