package main

import "fmt"

const (
	year1 = 2018 + iota
	year2 = year1 + iota
	year3 = year1 + iota
	year4 = year1 + iota
)

func main() {

	fmt.Println(year1, year2, year3, year4)

}
