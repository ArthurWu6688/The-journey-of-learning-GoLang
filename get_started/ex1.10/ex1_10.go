package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	urls := os.Args[1:]

	//	第一次请求
	for _, url := range urls {
		go fetah(url, ch, 1)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	//	第二次请求
	for _, url := range urls {
		go fetah(url, ch, 2)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetah(url string, ch chan string, times int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s, %v\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs, %7d, %s, 第%d次请求", secs, nbytes, url, times)
}
