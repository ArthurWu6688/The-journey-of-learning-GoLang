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
			fmt.Fprintf(os.Stderr, "http.Get err, %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%s\n", resp.Status)
		_, errCp := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if errCp != nil {
			fmt.Fprintf(os.Stderr, "io.Copy err, %v\n", err)
			os.Exit(1)
		}
	}
}
