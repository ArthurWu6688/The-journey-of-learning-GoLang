package main

import (
	"fmt"
	"os"
)

// 方式一
//func main() {
//	var args, sep string
//	for i := 0; i < len(os.Args); i++ {
//		args += sep + os.Args[i]
//		sep = " "
//	}
//	fmt.Println(args)
//}
///////////////////////

// 方式二
//func main() {
//	args, sep := "", ""
//	for _, arg := range os.Args[:] {
//		args += sep + arg
//		sep = " "
//	}
//	fmt.Println(args)
//}
///////////////////////

// 方式三
//func main() {
//	fmt.Println(strings.Join(os.Args[:], " "))
//}

///////////////////////

// 方式四
func main() {
	fmt.Println(os.Args)
}

///////////////////////
