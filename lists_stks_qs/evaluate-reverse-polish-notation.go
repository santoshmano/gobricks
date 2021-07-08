package lists_stks_qs

import (
	"github.com/santoshmano/gobricks/lists_stks_qs/stacks"
	"strconv"
)

func evalRPN(tokens []string) int {

	stk := stacks.NewArrayStack(len(tokens))

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
