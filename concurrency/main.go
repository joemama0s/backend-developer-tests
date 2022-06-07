package main

import (
	"fmt"
	"github.com/joemama0s/backend-developer-tests/concurrency/concur"
	"time"
)

func main() {
	mockUrls := []string{
		"google.com",
		"yahoo.com",
		"bing.com",
		"askjeeves.com",
		"github.com",
	}
	pool := concur.NewSimplePool(3)
	for _, url := range mockUrls {
		pool.Submit(func(){
			fmt.Printf("Processing URL %s\n", url)
			time.Sleep(1 * time.Second)
		})
	}

	time.Sleep(15 * time.Second)
}
