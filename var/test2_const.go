package main

import "fmt"

//iota关键字逐行累加
const (
	SHANGHAI = 10 * iota
	BEIJING
	SHENZHEN
)

func main() {
	const length int = 10
	fmt.Println("length = ", length)
	fmt.Println("SHANGHAI=", SHANGHAI)
	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("SHENZHEN=", SHENZHEN)
}
