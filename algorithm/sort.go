package algorithm

func quickSort(i, j int, arr []int) {
	if i >= j {
		return
	}
	l, r := i, j
	midValue := arr[j]
	for l < r {
		for l < r && arr[l] <= midValue {
			l++
		}
		for l < r && arr[r] > midValue {
			r--
		}
		if l < r {
			arr[l], arr[r] = arr[r], arr[l] //swap l,r
		}
	}
	quickSort(i, l-1, arr)
	quickSort(l+1, j, arr)
}

func bubbleSort(arr []int) []int {
	n := len(arr)
	quickSort(0, n-1, arr)
	return arr
}
