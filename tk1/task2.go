package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += strconv.Itoa(i) + sep + os.Args[i] + "// "
		sep = " "
	}

	fmt.Println(s)
}
