package utils

import (
	"errors"
	"fmt"
)

func SplitSlice(data []int64, sizeOfBatch int64) ([][]int64, error) {
	if sizeOfBatch <= 0 || data == nil {
		return nil, errors.New("input argument invalid")
	}

	var batchCount int
	batchSize := int(sizeOfBatch)

	if len(data)%batchSize == 0 {
		batchCount = len(data) / batchSize
	} else {
		batchCount = (len(data) / batchSize) + 1
	}

	res := make([][]int64, 0, batchCount)

	for i := 0; i < len(data); {
		end := i + batchSize

		if end > len(data) {
			end = len(data)
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
