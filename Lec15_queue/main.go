package main

import "fmt"

type Queue struct {
	data  []int
	front int
	rear  int
}

func (q *Queue) Enqueue(x int) {
	q.rear++
	q.data = append(q.data, x)
}

func (q *Queue) Dequeue() int {
	if q.front == q.rear {
		return -1
	}

	x := q.data[q.front]
	q.front++
	return x
}

func (q *Queue) Front() int {
	if q.front == q.rear {
		return -1
	}
	return q.data[q.front]
}

func (q *Queue) IsEmpty() bool {
	return q.front == q.rear
}

func main() {
	q := &Queue{front: 0, rear: 0}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println("Front element is: ", q.Front())
	fmt.Println("Dequeue: ", q.Dequeue())
	fmt.Println("Front element is: ", q.Front())
	fmt.Println("Queue is empty?", q.IsEmpty())

}
