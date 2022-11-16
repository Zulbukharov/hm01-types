package queue

type Queue struct{}

func (q *Queue) Enqueue(v int) {}
func (q *Queue) Dequeue() int  { return 0 }
func (q Queue) Values() []int  { return nil }
