package main

import (
	"fmt"
	"net/http"
	"time"
)



func main() {
	//A slice of sample websites
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com/",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}


	a_start_tume := time.Now()

	for _, url := range urls {
		go checkUrl(url)
	}

	a_end_time := time.Now()
	a_elapsed := a_end_time.Sub(a_start_tume)

	fmt.Println(a_elapsed, "after")


}

func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down !!!")
		return
	}

	fmt.Println(url, "is up and running.")
}
