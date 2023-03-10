package main

import "fmt"

type Stack struct {
	items []interface{}
}

func New() *Stack {
	return &Stack{}
}

func (this *Stack) Len() (size int) {
	return len(this.items)
}

func (this *Stack) Peek() (item interface{}) {
	if len(this.items) == 0 {
		return nil
	}

	return this.items[len(this.items)-1]
}

func (this *Stack) Pop() (item interface{}) {
	if len(this.items) == 0 {
		return nil
	}

	item = this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]
	return item
}

func (this *Stack) Push(value interface{}) {
	this.items = append(this.items, value)
}

func main() {
	s := New()
	s.Push("1")
	s.Push("2")
	s.Push(3)

	fmt.Println(s.Peek())
	fmt.Println(s.items...)

	fmt.Println(s.Pop())
	fmt.Println(s.items...)

	fmt.Println(s.Peek())
	fmt.Println(s.items...)
}
