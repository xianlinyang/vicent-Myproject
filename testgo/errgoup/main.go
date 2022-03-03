package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func ExampleGroup_justErrors() {
	var g errgroup.Group
	var urls = []string{
		"`444`",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			// Fetch the URL.
			fmt.Println(url)
			resp, err := http.Get(url)
			if err == nil {
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	//不管是否有协程执行失败, wait()都要等待所有协程执行完成,就算中途有协程出错也要执行完其他协程
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println(err)
	}
}

func main() {
	ExampleGroup_justErrors()
}
