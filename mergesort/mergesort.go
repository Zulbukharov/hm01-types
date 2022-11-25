package mergesort

func MergeSort(input []int) []int {
	if len(input) == 0 || len(input) == 1 {
		return input
	}

	left := MergeSort(input[:len(input)/2])
	right := MergeSort(input[len(input)/2:])

	// fmt.Println(left, right)
	return Merge(left, right)
}

func Merge(left []int, right []int) []int {
	res := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {

			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}

	for i < len(left) {
		res = append(res, left[i])
		i++
	}

	for j < len(right) {
		res = append(res, right[j])
		j++
	}

	return res
}

// HAVE TRIED THIS OPTION BUT COULDN'T PROPERLY IMPLEMENT RECURSIVE CALL FOR THIS FUNCTION

// var (
// 	k    int
// 	arr  = []int{}
// 	arr2 = []int{}
// )

// func MergeSort(input []int) []int {
// 	if len(input) == 0 || len(input) == 1 {
// 		return input
// 	}
// 	count := -201
// 	for i := 0; i < len(input); i++ {
// 		if input[i] > count {
// 			count = input[i]
// 			k = i

// 		}
// 	}
// 	arr = append(arr, input[k])

// 	count2 := 0
// 	for j := 0; j < len(input); j++ {
// 		if input[k] != input[j] {
// 			input[count2] = input[j]
// 			count2++
// 		}
// 	}

// 	arr2 = input[:count2]

// 	if len(arr2) >= 1 {
// 		Repeat(arr2)
// 	}

// 	unreversed := arr
// 	rev := []int{}
// 	for n := range unreversed {
// 		rev = append(rev, unreversed[len(unreversed)-1-n])
// 	}

// 	return rev
// }

