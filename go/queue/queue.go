package queue

type Queue struct {
	elements []int
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Enqueue(data int) {
	q.elements = append(q.elements, data)
}

func (q *Queue) Dequeue() int {
	front := q.elements[0]
	q.elements = q.elements[1:]
	return front
}
