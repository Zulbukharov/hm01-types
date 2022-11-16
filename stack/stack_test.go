package stack

import (
	"reflect"
	"testing"
)

var tests = []struct {
	push   []int
	pop    []int
	values []int
}{
	{[]int{0}, []int{0}, []int{}},
	{[]int{1, 2, 3, 4}, []int{4, 3}, []int{1, 2}},
	{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{9, 8, 7, 6}, []int{1, 2, 3, 4, 5}},
}

func TestStack(t *testing.T) {
	for _, test := range tests {
		actual := Stack{}
		for _, v := range test.push {
			actual.Push(v)
		}
		for _, v := range test.pop {
			tmp := actual.Pop()
			if v != tmp {
				t.Errorf("Stack(%v) expected %v, got %v", test.push, v, tmp)
			}
		}
		if !reflect.DeepEqual(test.values, actual.Values()) {
			t.Errorf("Stack(%v) expected %v, got %v", test.push, test.values, actual.Values())
		}
	}
}
