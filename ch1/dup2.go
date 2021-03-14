package main

import (
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Printf("failed")
	}
}

func countLines() {

}
