package stacks

type Stack struct {
	elems []interface{}
}

func NewStack(len int) *Stack {
	stk := Stack{}
	stk.elems = make([]interface{}, len)
	return &stk
}

func (stk *Stack) Push(elem interface{}) {
	stk.elems = append(stk.elems, elem)
}

func (stk *Stack) Pop() interface{} {
	if len(stk.elems) == 0 {
		return nil
	}
	elem := stk.elems[len(stk.elems)-1]
	stk.elems = stk.elems[:len(stk.elems)-1]
	return elem
}

func (stk *Stack) Peek() interface{} {
	if len(stk.elems) == 0 {
		return nil
	}
	return stk.elems[len(stk.elems)-1]
}
