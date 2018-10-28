package mergesort

func MergeSort(arr []int) {
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, l, r int) {
	if l < r {
		m := (l + r) / 2
		mergeSort(arr, l, m)
		mergeSort(arr, m+1, r)
		merge(arr, l, m, r)
	}
}

func merge(arr []int, l, m, r int) {
	n1 := m-l+1
	n2 := r-m
	left := make([]int, n1)
	right := make([]int, n2)
	copy(left, arr[l:m+1])
	copy(right, arr[m+1:r+1])
	i := 0
	j := 0
	k := l
	for i < n1 && j < n2 {
		if left[i] < right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}
	for i < n1 {
		arr[k] = left[i]
		i++
		k++
	}
	for j < n2 {
		arr[k] = right[j]
		j++
		k++
	}
}
