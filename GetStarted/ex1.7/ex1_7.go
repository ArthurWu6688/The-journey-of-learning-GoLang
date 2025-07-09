package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "http.Get(url) failed, %v", err)
			os.Exit(1)
		}

		_, cpErr := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if cpErr != nil {
			fmt.Fprintf(os.Stderr, "io.Copy err, %v", cpErr)
			os.Exit(1)
		}
	}
}
