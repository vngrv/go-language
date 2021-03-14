package main

import "fmt"

type age int

func (a age) isAdult() bool {
	return a >= 18
}

func main() {
	myAge := age(25)

	fmt.Println(myAge.isAdult())
}
