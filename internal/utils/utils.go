package utils

import (
	"errors"
	"fmt"
)

func SplitSlice(data []int64, batchSize int64) ([][]int64, error) {
	if batchSize <= 0 || data == nil {
		return nil, errors.New("input argument invalid")
	}

	var batchCount int64

	if int64(len(data))%batchSize == 0 {
		batchCount = int64(len(data)) / batchSize
	} else {
		batchCount = (int64(len(data)) / batchSize) + 1
	}

	res := make([][]int64, 0, batchCount)

	for i := int64(0); i < int64(len(data)); {
		end := i + batchSize

		if end > int64(len(data)) {
			end = int64(len(data))
		}

		res = append(res, data[i:end])
		i = end
	}

	return res, nil
}

func ReverseKey(data map[int64]string) (map[string]int64, error) {
	if data == nil {
		return nil, errors.New("input argument invalid")
	}

	res := make(map[string]int64, len(data))

	for k, v := range data {
		if _, found := res[v]; found {
			return nil, fmt.Errorf("duplicate key: %s", v)
		}
		res[v] = k
	}

	return res, nil
}

func FilterSlice(data []int64, filter []int64) ([]int64, error) {
	if data == nil || filter == nil {
		return nil, errors.New("input argument invalid")
	}

	filterMap := make(map[int64]struct{}, len(filter))
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
