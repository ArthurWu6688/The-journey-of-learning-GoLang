package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入一些内容：")
	for input.Scan() {
		fmt.Println(input.Text())
	}
}
