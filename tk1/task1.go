/**
Изменить программу echo так, чтобы она выводила также os.Args[0], имя выполняемой команды
*/

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
	cout >> /var/folders/4t/p0dklsgn2hs5mwvfm1q39nbr0000gp/T/go-build014016523/b001/exe/task1 Hello world
*/
func main() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
