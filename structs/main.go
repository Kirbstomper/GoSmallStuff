package main

//Status Checker
import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://amazon.com",
		"http://twitter.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	for link := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(link)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
