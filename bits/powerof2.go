package bits

func IsPowerOf2_1(val uint8) bool {
	return (val > 0) && ((val & (val - 1)) == 0)
}

func IsPowerOf2_2(val uint8) bool {
	return (val > 0) && ((val & ^(val - 1)) == val)
}

//val > 1 ? return !(i & 0x1) ? return (i & 0x1)
