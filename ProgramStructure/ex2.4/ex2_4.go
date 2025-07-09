package main

import (
	"fmt"
	popcount "study_go/ProgramStructure/ex2.4/ex2_4_popcount"
	"time"
)

func main() {
	var a uint64 = 0x123456789ABCDEF
	start1 := time.Now()
	ret1 := popcount.PopCount(a)
	end1 := time.Since(start1)

	start2 := time.Now()
	ret2 := popcount.PopCount1(a)
	end2 := time.Since(start2)

	fmt.Printf("ret1:%d time1:%f\nret2:%d time2:%f\n",
		ret1, end1.Seconds(), ret2, end2.Seconds())
}
