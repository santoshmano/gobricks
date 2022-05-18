package strings

import "fmt"

// probably how not to approach such a problem :(

func myAtoi(s string) int {

	isNum := func(i int) bool {
		if s[i] >= '0' && s[i] <= '9' {
			return true
		}
		return false
	}

	i := 0
	spaceCount := 0
	signCount := 0
	// ignore all characters until first digit
	for i < len(s) && !isNum(i) {
		if s[i] == ' ' {
			spaceCount++
			if signCount > 0 {
				return 0
			}
		} else if s[i] == '-' || s[i] == '+' {
			signCount++
		} else {
			return 0
		}
		i++
	}

	// check if there is a sign
	sign := 1
	if i < len(s) && i != 0 {
		if s[i-1] == '-' {
			sign = -1
		}

		if s[i-1] == '-' || s[i-1] == '+' {
			if spaceCount != i-1 {
				return 0
			}
		}
	}

	// read all digits & form the number
	var num uint64 = 0
	for i < len(s) && isNum(i) {
		num = num*10 + uint64(s[i]-'0')
		fmt.Println(num, s[i])
		i++
	}

	fmt.Println(num, sign)

	if sign == 1 && num > (1<<31-1) {
		return 1<<31 - 1
	} else if sign == -1 && num > (1<<31) {
		return -1 << 31
	}

	return int(num) * sign
}
