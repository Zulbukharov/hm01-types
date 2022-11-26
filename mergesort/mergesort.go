package mergesort

func MergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := MergeSort(items[:len(items)/2])
	second := MergeSort(items[len(items)/2:])

	return merge(first, second)
}

func merge(first []int, second []int) []int {
	res := []int{}
	i := 0
	j := 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			res = append(res, first[i])
			i++
		} else {
			res = append(res, second[j])
			j++
		}
	}
	for ; i < len(first); i++ {
		res = append(res, first[i])
	}
	for ; j < len(second); j++ {
		res = append(res, second[j])
	}
	return res
}

