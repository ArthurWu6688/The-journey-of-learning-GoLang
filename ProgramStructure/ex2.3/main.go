package main

import (
	"fmt"
	popcount "study_go/ProgramStructure/ex2.3/ex2_3_popcount"
	"time"
)

func main() {
	var a uint64 = 0x1234567890ABCDEF
	time1 := time.Now()
	ret1 := popcount.PopCount(a)
	end1 := time.Since(time1)

	time2 := time.Now()
	ret2 := popcount.PopCount1(a)
	end2 := time.Since(time2)

	fmt.Printf("ret1:%d t1=%f, ret2:%d t2=%f\n", ret1, end1.Seconds(), ret2, end2.Seconds())
}
