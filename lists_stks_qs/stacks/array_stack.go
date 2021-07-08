package stacks

type ArrayStack struct {
	elems []interface{}
}

func NewArrayStack(len int) *ArrayStack {
	stk := ArrayStack{}
	stk.elems = make([]interface{}, len)
	return &stk
}

func (stk *ArrayStack) Push(elem interface{}) {
	stk.elems = append(stk.elems, elem)
}

func (stk *ArrayStack) Pop() interface{} {
	if len(stk.elems) == 0 {
		return nil
	}
	elem := stk.elems[len(stk.elems)-1]
	stk.elems = stk.elems[:len(stk.elems)-1]
	return elem
}

func (stk *ArrayStack) Peek() interface{} {
	if len(stk.elems) == 0 {
		return nil
	}
	return stk.elems[len(stk.elems)-1]
}
