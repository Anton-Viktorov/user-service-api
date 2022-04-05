package utils

import (
	"errors"
)

func SplitSlice(data []int64, batchSize int64) ([][]int64, error) {
	if batchSize <= 0 || data == nil {
		return nil, errors.New("input argument invalid")
	}

	var res [][]int64
	batchCount := int64(len(data)) / batchSize

	for i := int64(1); i <= batchCount; i++ {
		res = append(res, data[(i-1)*batchSize:i*batchSize])
	}

	if batchSize*batchCount < int64(len(data)) {
		res = append(res, data[batchSize*batchCount:])
	}

	return res, nil
}

func ReverseKey(data map[int64]string) (map[string]int64, error) {
	if data == nil {
		return nil, errors.New("input argument invalid")
	}

	res := make(map[string]int64)

	for k, v := range data {
		if _, found := res[v]; !found {
			res[v] = k
		} else {
			return nil, errors.New("duplicate key: " + v)
		}
	}

	return res, nil
}

func FilterSlice(data []int64, filter []int64) ([]int64, error) {
	if data == nil || filter == nil {
		return nil, errors.New("input argument invalid")
	}

	filterMap := make(map[int64]struct{})
	for _, val := range filter {
		filterMap[val] = struct{}{}
	}

	var res []int64
	for _, v := range data {
		if _, found := filterMap[v]; !found {
			res = append(res, v)
		}
	}

	return res, nil
}
