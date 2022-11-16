package queue

import (
	"reflect"
	"testing"
)

var tests = []struct {
	enqueue []int
	dequeue []int
	values  []int
}{
	{[]int{0}, []int{0}, []int{}},
	{[]int{1, 2, 3, 4}, []int{1, 2}, []int{3, 4}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9}},
}

func TestQueue(t *testing.T) {
	for _, test := range tests {
		actual := Queue{}
		for _, v := range test.enqueue {
			actual.Enqueue(v)
		}
		for _, v := range test.dequeue {
			tmp := actual.Dequeue()
			if v != tmp {
				t.Errorf("Queue(%v) expected %v, got %v", test.enqueue, v, tmp)
			}
		}
		if !reflect.DeepEqual(test.values, actual.Values()) {
			t.Errorf("Queue(%v) expected %v, got %v", test.enqueue, test.values, actual.Values())
		}
	}
}
