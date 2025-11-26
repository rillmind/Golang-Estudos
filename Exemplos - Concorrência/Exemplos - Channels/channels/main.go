package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://amazon.com",
		"https://golang.org",
		"https://my-siteh.com",
		"https://facebook.com",
		"https://stackoverflow.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	for link := range channel {
		f := func(l string) {
			time.Sleep(2 * time.Second)
			checkLink(l, channel)
		}
		go f(link)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is OK")
	c <- link
}
