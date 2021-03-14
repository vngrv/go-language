package main

import "fmt"

func printType(value interface{}) {
	switch value.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("something else")
	}
}

func main() {
	printType(3)
	printType("string")
	printType([]string{"slices", "slices"})
}
