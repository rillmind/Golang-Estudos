package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var (
		wg    sync.WaitGroup
		links = []string{
			"https://google.com",
			"https://amazon.com",
			"https://golang.org",
			"https://my-siteh.com",
			"https://facebook.com",
			"https://stackoverflow.com",
		}
	)

	wg.Add(len(links))
	for _, link := range links {
		go checkLink(link, &wg)
	}
	wg.Wait()
}

func checkLink(link string, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "Might be down")
		return
	}

	fmt.Println(link, "Is OK")
}
