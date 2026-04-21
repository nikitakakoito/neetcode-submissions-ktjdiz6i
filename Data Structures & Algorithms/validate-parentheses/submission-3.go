func isValid(s string) bool {
    st := NewStack()
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, v := range s {
		opener, okk := pairs[v]

		if okk {
			vv, ok := st.Top()
			if !ok {return false}
			if vv == opener {
				st.Pop()
			} else {
				return false
			}
		} else {
			st.Push(v)
		}
	}
_, ok := st.Top()

return !ok
}


type Stack struct {
	st []rune
}

func NewStack() Stack {
	return Stack{[]rune{}}
}

func (this *Stack) Push(val rune) {
	this.st = append(this.st, val)
}

func (this *Stack) Pop() {
	if len(this.st) == 0{
		return
	}
	this.st = this.st[:len(this.st)-1]
}

func (this *Stack) Top() (rune, bool) {
	if len(this.st) == 0 {
		return 0, false
	}
	return this.st[len(this.st)-1], true
}