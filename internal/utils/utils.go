package utils

import (
	"errors"
)

func SplitSlice(data []int64, batchSize int64) ([][]int64, error) {
	if batchSize <= 0 || data == nil {
		return nil, nil
	}

	res := make([][]int64, 0)
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
		return nil, nil
	}

	res := make(map[string]int64)

	for k, v := range data {
		if _, ok := res[v]; !ok {
			res[v] = k
		} else {
			err := errors.New("duplicate key: " + v)
			return nil, err
		}
	}
	return res, nil

}

func FilterSlice(data, omitValues []int64) ([]int64, error) {
	if data == nil || omitValues == nil {
		return data, nil
	}

	omitValuesMap := make(map[int64]struct{})
	for _, val := range omitValues {
		omitValuesMap[val] = struct{}{}
	}

	res := make([]int64, 0)
	for _, sliceValue := range data {
		if _, ok := omitValuesMap[sliceValue]; ok {
			continue
		}
		res = append(res, sliceValue)
	}
	return res, nil
}
