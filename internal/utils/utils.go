package utils

func SplitSlice(slice []int64, n int64) [][]int64 {
	var i int64
	result := make([][]int64, 0)
	lastIndex := (int64(len(slice)) / n) * n

	for i < lastIndex {
		result = append(result, slice[i:i+n])
		i += n
	}
	if lastIndex < int64(len(slice)) {
		result = append(result, slice[lastIndex:])
	}
	return result
}

func ReverseKey(m map[int64]int64) map[int64]int64 {
	result := make(map[int64]int64)

	for k, v := range m {
		result[v] = k
	}
	return result
}

func FilterSlice(slice []int64) []int64 {
	omitValues := [5]int64{1, 3, 5, 7, 9}
	result := make([]int64, 0)

	needOmit := func(v int64, list *[5]int64) bool {
		for _, value := range list {
			if value == v {
				return true
			}
		}
		return false
	}

	for _, sliceValue := range slice {
		if needOmit(sliceValue, &omitValues) {
			continue
		}
		result = append(result, sliceValue)
	}
	return result
}
