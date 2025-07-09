package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	start := time.Now()
	urls := os.Args[1:]
	ch := make(chan string)
	i := 1

	for _, u := range urls {
		parsedURL, err := url.Parse(u)
		if err != nil {
			fmt.Fprintf(os.Stderr, "url.Parse err, %v\n", err)
			os.Exit(1)
		}
		file := parsedURL.Host + ".txt"
		f, err := os.Create(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "文件创建失败, %v\n", err)
			os.Exit(1)
		}
		go fetch(f, u, ch, i)
		i++
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("总耗时：%.2fs\n", time.Since(start).Seconds())
}

func fetch(f *os.File, url string, ch chan string, times int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get失败, %v\n", err)
		os.Exit(1)
	}
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("io.Copy err, %v\n", err)
		os.Exit(1)
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs, %7d, %s, 第%d次请求\n", secs, nbytes, url, times)
}
