package stacks

import (
	"strconv"
)

func evalRPN(tokens []string) int {

	stk := stacks2.NewArrayStack(len(tokens))

	for _, token := range tokens {

		switch token {
		case "*":
			second, _ := stk.Pop().(int)
			first, _ := stk.Pop().(int)
			stk.Push(first * second)

		case "/":
			second, _ := stk.Pop().(int)
			first, _ := stk.Pop().(int)
			stk.Push(first / second)

		case "+":
			second, _ := stk.Pop().(int)
			first, _ := stk.Pop().(int)
			stk.Push(first + second)

		case "-":
			second, _ := stk.Pop().(int)
			first, _ := stk.Pop().(int)
			stk.Push(first - second)

		default:
			t, _ := strconv.Atoi(token)
			stk.Push(t)

		}

	}

	return stk.Pop().(int)
}
