package main

import "fmt"

type Queue struct {
	items []interface{}
	s     int
	e     int
}

func New() *Queue {
	return &Queue{}
}

func (this *Queue) Enqueue(item interface{}) {
	this.items = append(this.items, item)
	this.e = len(this.items) - 1
}

func (this *Queue) Dequeue() interface{} {
	if len(this.items) == 0 {
		return nil
	}
	if this.s > this.e {
		return nil
	}

	item := this.items[this.s]
	this.s++
	return item
}

func main() {
	q := New()

	fmt.Println(q.Dequeue())
	q.Enqueue("Item1")
	q.Enqueue("Item2")
	q.Enqueue("Item3")
	q.Enqueue("Item4")
	q.Enqueue("Item5")
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
}
