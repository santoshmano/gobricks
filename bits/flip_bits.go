package bits

import (
	"fmt"
	"unsafe"
)

func FlipBits1(val int) int {
	// Using XOR
	size := unsafe.Sizeof(val)
	size = size * 8
	mask := 0x1

	for size-1 > 0 {
		mask = mask<<1 | 1
		size--
	}
	// TODO - figure out how to print debug/log messages
	fmt.Println(val, mask, val^mask)
	return val ^ mask
}

func FlipBits2(val int) int {
	// Using Counting Technique
	size := unsafe.Sizeof(val)
	fmt.Println(size)
	return 1
}
