package bits

import (
	"fmt"
	"testing"
)

func powerOf2TestData() []struct {
	input uint8
	want  bool
} {
	return []struct {
		input uint8
		want  bool
	}{
		{0, false},
		{1, true},
		{2, true},
		{3, false},
		{4, true},
		{5, false},
		{6, false},
		{7, false},
		{8, true},
		{9, false},
		{10, false},
		{128, true},
		{255, false},
	}
}

func TestIsPowerOf2(t *testing.T) {

	functions := []func(uint8) bool{IsPowerOf2_1, IsPowerOf2_1}

	for i, f := range functions {
		// TODO : print function name instead of i
		fmt.Println("Testing function IsPowerOf2 ", i)

		for _, testcase := range powerOf2TestData() {
			result := f(testcase.input)
			if result == testcase.want {
				t.Log(`test success`)
			} else {
				t.Errorf(`IsPowerOf2(%v) generated: %v expected: %v`,
					testcase.input,
					result,
					testcase.want)
			}
		}

	}
}

// another probably better way of doing table tests
var flipBitsTests = []struct {
	input int
	want  int
}{
	{0, ^0},
	{1, ^1},
	{2, ^2},
	{3, ^3},
	{4, ^4},
	{5, ^5},
	{6, ^6},
	{7, ^7},
	{8, ^8},
	{9, ^9},
	{10, ^10},
	{128, ^128},
	{255, ^255},
}

func TestFlipBits(t *testing.T) {

	functions := []func(int) int{FlipBits1}

	for i, f := range functions {
		// TODO : print function name instead of i
		fmt.Println("Testing function IsPowerOf2 ", i)

		for _, testcase := range flipBitsTests {
			result := f(testcase.input)
			if result == testcase.want {
				t.Log(`test success`)
			} else {
				t.Errorf(`IsPowerOf2(%v) generated: %v expected: %v`,
					testcase.input,
					result,
					testcase.want)
			}
		}

	}
}
