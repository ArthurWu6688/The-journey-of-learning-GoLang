package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") || !strings.HasPrefix(url, "http://") {
			url = "https://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "http.Get() error, %v\n", err)
			os.Exit(1)
		}
		io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
	}
}
