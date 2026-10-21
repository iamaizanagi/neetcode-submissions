type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	var s MinStack

	for j := 0; j < len(s.stack); j++{
		if s.stack[j] != int(s.stack[j]){
			s.Pop()
		}else{
			s.Push(int(s.stack[j]))
		}
	}
	return s 
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	//return &this
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	//return &this
}

func (this *MinStack) Top() int {
	top := this.stack[len(this.stack)-1]
	return top
}

func (this *MinStack) GetMin() int {
	minimum := this.stack[0]
	for i := 0; i < len(this.stack); i++{
		if this.stack[i] < minimum{
			minimum = this.stack[i]
		}
	}
	return minimum
}
