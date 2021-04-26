package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}
	for _, link := range links{
		go checkLink(link)
		time.Sleep(2 * time.Second)
	}
}
func checkLink(link string){
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link," might be down!")
	}

	fmt.Println(link, "is up!")
}