package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	urls := os.Args[1:]
	ch := make(chan string)
	start := time.Now()
	i := 1

	for _, u := range urls {
		go fetch(u, ch, i)
		i++
	}

	for range urls {
		fmt.Println(<-ch)
	}

	for _, u := range urls {
		go fetch(u, ch, i)
		i++
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("总耗时%.2fs\n", time.Since(start).Seconds())
}

func fetch(u string, ch chan string, times int) {
	start := time.Now()
	resp, err := http.Get(u)
	if err != nil {
		ch <- fmt.Sprintf("http.Get err, %v\n", err)
		os.Exit(1)
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("io.Copy err, %v\n", err)
		os.Exit(1)
	}
	//ch <- fmt.Sprintf("对%s第%d次请求读取了%d字节，耗时%.2fs\n", u, times, nbytes, time.Since(start).Seconds())
	ch <- fmt.Sprintf("对%s请求读取了%d字节，耗时%.2fs\n", u, nbytes, time.Since(start).Seconds())
}

//func getDomain(u string) string{
//
//}
