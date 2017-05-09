package main

import (
	"container/list"
	"fmt"

	"github.com/joncalhoun/queue"
)

func main() {
	iq := queue.NewInt()
	iq.Enqueue(123)
	iq.Enqueue(456)
	fmt.Println(iq.Dequeue())

	isq := queue.NewIntSlice()
	isq.Enqueue([]int{1, 2, 3})
	isq.Enqueue([]int{4, 5, 6})
	fmt.Println(isq.Dequeue())

	sq := queue.NewString()
	sq.Enqueue("jon")
	sq.Enqueue("calhoun")
	fmt.Println(sq.Dequeue())

	// And we can even get a little meta
	lq := queue.NewList()
	list1 := list.New()
	list1.PushBack("jon")
	list2 := list.New()
	list2.PushBack(456)
	lq.Enqueue(list1)
	lq.Enqueue(list2)
	fmt.Println(lq.Dequeue().Front().Value)
	fmt.Println(lq.Dequeue().Front().Value)
}
