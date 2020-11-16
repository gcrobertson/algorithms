package main

func main() {

}

type node struct {
	value int
	next  *node
}

type MinStack struct {
	stackTop *node // stack of all values
	minTop   *node // stack of `min` values
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	t := &node{
		value: x,
	}
	if this.stackTop == nil {
		this.stackTop = t
		this.minTop = t
	} else {
		t.next = this.stackTop
		this.stackTop = t
		if x <= this.minTop.value {
			n := &node{
				value: x,
				next:  this.minTop,
			}
			this.minTop = n
		}
	}
}

func (this *MinStack) Pop() {
	t := this.stackTop
	this.stackTop = this.stackTop.next
	if t.value == this.minTop.value {
		this.minTop = this.minTop.next
	}
	// return t
}

func (this *MinStack) Top() int {
	if this.stackTop == nil {
		return 0
	}
	return this.stackTop.value
}

func (this *MinStack) GetMin() int {
	if this.minTop == nil {
		return 0
	}
	return this.minTop.value
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
