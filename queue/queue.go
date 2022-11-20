package queue

type Queue struct {
	QueueInt []int
}

func (q *Queue) Enqueue(v int) {
	q.QueueInt = append(q.QueueInt, v)
}
func (q *Queue) Dequeue() int {
	if len(q.QueueInt) == 1 {
		res := q.QueueInt[0]
		q.QueueInt = []int{}
		return res
	}
	for len(q.QueueInt) > 0 {
		res := q.QueueInt[0]
		q.QueueInt = q.QueueInt[1:]
		return res
	}
	return 0
}
func (q Queue) Values() []int {
	return q.QueueInt
}
