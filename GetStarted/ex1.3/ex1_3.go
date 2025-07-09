package main

import (
	"fmt"
	"strings"
	"time"
)

const N = 10000
const fragment = "hello"

func concatPlus() string {
	s := ""
	sep := ""
	for i := 0; i < N; i++ {
		s += sep + fragment
		sep = " "
	}
	return s
}

func concatJoin() string {
	parts := make([]string, N)
	for i := 0; i < N; i++ {
		parts[i] = fragment
	}
	return strings.Join(parts, " ")
}

func measure(name string, f func() string) {
	start := time.Now()
	result := f()
	duration := time.Since(start)
	fmt.Println(name, "，字符串长度：", len(result), "，耗时：", duration)
}

func main() {
	measure("plus(+)", concatPlus)
	measure("strings.Join()", concatJoin)
}
