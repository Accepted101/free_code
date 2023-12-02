package algorithm

func carPooling(trips [][]int, capacity int) bool {

	cnt := make([]int, 1002, 1002)

	for i := 0; i < len(trips); i++ {
		tmp := trips[i]
		cnt[tmp[1]] += tmp[0]
		cnt[tmp[2]+1] -= tmp[0]
	}

	for i := 1; i <= 1000; i++ {
		cnt[i] += cnt[i-1]
		if cnt[i] > capacity {
			return false
		}
	}

	return true

}
