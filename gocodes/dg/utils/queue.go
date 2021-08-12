package utils

import (
	"fmt"
)

// type Queue []interface{}

type Queue struct {
	queue []interface{}
	ptr   int
}

func (q Queue) Len() int {
	return len(q.queue)
}

func (q Queue) IsEmpty() bool {
	return q.Len() == 0
}

// 入队
func (q *Queue) EnQueue(v interface{}) int {
	index := q.Len()
	q.queue = append(q.queue, v)

	return index
}

// 出队
func (q *Queue) DeQueue() interface{} {
	v := q.queue[0]
	q.queue = q.queue[1:]

	return v
}

// 指定某一个出队
func (q *Queue) DeQueueWithIndex(index int) error {
	length := q.Len()
	if index >= length {
		fmt.Println("index:", index, " queue len:", length)
		return fmt.Errorf("index out of range")
	}

	q.queue = append((q.queue)[:index], (q.queue)[index+1:]...)

	return nil
}

func (q *Queue) EnQueueWithIndex(v interface{}, index int) {
	q.queue[index] = v
}

func (q *Queue) Next() interface{} {
	length := q.Len()

	if q.ptr >= length {
		q.ptr = 0
		return nil
	}

	v := q.queue[q.ptr]
	q.ptr++

	return v
}
