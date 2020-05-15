package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"http://www.amazon.com",
		"http://www.golang.org",
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.stackoverflow.com",
		"http://www.youtube.com",
		"http://www.apple.com",
		"http://www.airbnb.com",
		"http://www.netflix.com",
		"http://www.linkedin.com",
	}

	c := make(chan string)

	for _, s := range sites {
		go checkLink(s, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(time.Second * 2)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	c <- link

	fmt.Println(link, "is up!")
}
