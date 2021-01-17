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

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
