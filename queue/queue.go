package queue

type Queue struct {
	data []int
}

func (q *Queue) Enqueue(v int) {
	if len(q.data) == 0 {
		q.data = append(q.data)
	}
	q.data = append(q.data, v)
}

func (q *Queue) Dequeue() int {
	if len(q.data) == 0 && len(q.data) < 0 {
		return 0
	}

	firstelem := q.data[0]
	q.data = q.data[1:len(q.data)]
	return firstelem
}

func (q Queue) Values() []int {
	return q.data
}
