package main

import (
	"fmt"
	"os"
)

/**

Usage:
	go run echo1.go Hello world
Result:
	cout >> Hello world
*/
func main() {
	var s, sep string

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
