package queue

type Queue struct {
	values []int
}

func (q *Queue) Enqueue(v int) {
	q.values = append(q.values, v)
}
func (q *Queue) Dequeue() int {
	if len(q.values) == 0 {
		return 0
	}
	v := q.values[0]
	q.values = q.values[1:]
	return v
}
func (q Queue) Values() []int {
	return q.values
}
