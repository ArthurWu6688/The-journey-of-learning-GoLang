package main

import (
	"fmt"
	"os"
	"strconv"
	tempconv2 "study_go/ProgramStructure/ex2.1/ex2.1_tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ParseFloat error\n")
			os.Exit(1)
		}
		//k := tempconv.Kelwin(t)
		f := tempconv2.Fahrenheit(t)
		c := tempconv2.Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv2.FtoC(f), c, tempconv2.CtoF(c))

	}
}
