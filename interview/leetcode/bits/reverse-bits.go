package bits

func reverseBits(num uint32) uint32 {
	var reverse uint32
	reverse = 0

	for i := 0; i < 32; i++ {
		reverse <<= 1
		if num&1 == 1 {
			reverse = reverse | 1
		}
		num >>= 1
	}

	return reverse
}

// looked up the below solution, super fun shifting going on
// a - 1010, c - 1100, 3 - 11, 5 - 101
func reverseBits_1(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}
