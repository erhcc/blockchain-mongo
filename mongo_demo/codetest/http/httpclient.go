package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	c1 := httpClient()
	c2 := httpClient()
	c3 := httpClient()

	fmt.Println(c1,c2,c3)

	fmt.Println(c1==c2)

}

func httpClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 10 * time.Second,
	}

	return client
}