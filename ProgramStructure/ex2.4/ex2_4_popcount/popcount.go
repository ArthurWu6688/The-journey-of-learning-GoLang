// Package popcount 是一个计算一个数的二进制中有多少个1的包
package popcount

var pc [256]byte = func() (pc [256]byte) {
	for i, _ := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()

func PopCount(x uint64) int {
	return int(pc[byte(x>>(8*0))] +
		pc[byte(x>>(8*1))] +
		pc[byte(x>>(8*2))] +
		pc[byte(x>>(8*3))] +
		pc[byte(x>>(8*4))] +
		pc[byte(x>>(8*5))] +
		pc[byte(x>>(8*6))] +
		pc[byte(x>>(8*7))])
}

func PopCount1(x uint64) int {
	count := 0
	for i := 0; i < 64; i++ {
		if (x & 1) == 1 {
			count++
		}
		x >>= 1
	}
	return count
}
