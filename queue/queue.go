package queue

type Queue struct {
	Obj []int
}

func (q *Queue) Enqueue(v int) {
	q.Obj = append(q.Obj, v)
}

func (q *Queue) Dequeue() int {
	if len(q.Obj) == 0 {
		return 0
	}
	elem := q.Obj[0]
	q.Obj = q.Obj[1:]
	return elem
}

func (q Queue) Values() []int {
	return q.Obj
}
