package main

import (
	"fmt"
	"os"
	"strconv"
	commconv "study_go/ProgramStructure/ex2.2/ex2_2_conv"
)

func main() {
	args := os.Args[1:]
	if len(args) > 0 {
		if args[0] == "1" {
			for _, arg := range args[1:] {
				value, err := strconv.ParseFloat(arg, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "ParseFloat err")
					os.Exit(1)
				}

				m := commconv.Meter(value)
				ft := commconv.Feet(value)
				fmt.Printf("%s=%s, %s=%s\n", m, commconv.MtoF(m), ft, commconv.FtoM(ft))
			}
		} else {
			for _, arg := range args[1:] {
				value, err := strconv.ParseFloat(arg, 64)
				if err != nil {
					fmt.Fprintf(os.Stderr, "ParseFloat err")
					os.Exit(1)
				}

				k := commconv.Kilogram(value)
				p := commconv.Pound(value)
				fmt.Printf("%s=%s, %s=%s\n", k, commconv.KtoP(k), p, commconv.PtoK(p))
			}
		}
	} else {
		var value float64
		var flag int8
		fmt.Printf("input value and flag: ")
		fmt.Scan(&value, &flag)
		if flag == 1 {
			m := commconv.Meter(value)
			ft := commconv.Feet(value)
			fmt.Printf("%s=%s, %s=%s\n", m, commconv.MtoF(m), ft, commconv.FtoM(ft))
		} else {
			k := commconv.Kilogram(value)
			p := commconv.Pound(value)
			fmt.Printf("%s=%s, %s=%s\n", k, commconv.KtoP(k), p, commconv.PtoK(p))
		}
	}
}
