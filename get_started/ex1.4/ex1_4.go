package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for str, n := range counts {
		fmt.Printf("%s\t%d\n", str, n)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

	for str, n := range counts {
		if n > 1 {
			StrNew := str + " 所属文件：" + f.Name()
			delete(counts, str)
			counts[StrNew] = n
		}
	}
}
