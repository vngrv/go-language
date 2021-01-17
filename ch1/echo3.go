package main

import (
	"fmt"
	"os"
	"strings"
)

/**

Usage:
	go run echo3.go Hello world
Result:
	cout >> Hello world
*/
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
