package main

import (
	"fmt"
	"net/http"
	"sync"
)

var w sync.WaitGroup
var mu sync.Mutex
var su = []string{
	"test",
}

func main() {

	apis := []string{
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.amazon.ca/amazonprime",
		"https://www.linkedin.com",
	}

	for _, val := range apis {
		go greeting(val)
		w.Add(1)
	}
	w.Wait()

}

func greeting(s string) {
	mu.Lock()
	su = append(su, s)
	mu.Unlock()
	response, err := http.Get(s)
	if err != nil {
		fmt.Println("Errro caling " + s)
	} else {
		fmt.Printf("Response code is %d\n", response.StatusCode)
	}
	defer w.Done()
}
