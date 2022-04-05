package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSplitSlice(t *testing.T) {

	type Dataset struct {
		ChunkSize   int64
		InputSlice  []int64
		OutputSlice [][]int64
	}

	datasets := []Dataset{
		{
			ChunkSize:   2,
			InputSlice:  []int64{1, 2, 3, 4},
			OutputSlice: [][]int64{{1, 2}, {3, 4}},
		},
		{
			ChunkSize:   3,
			InputSlice:  []int64{1, 2, 3, 4, 5},
			OutputSlice: [][]int64{{1, 2, 3}, {4, 5}},
		},
		{
			ChunkSize:   2,
			InputSlice:  []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			OutputSlice: [][]int64{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9}},
		},
		{
			ChunkSize:   10,
			InputSlice:  []int64{1, 2, 3, 4, 5},
			OutputSlice: [][]int64{{1, 2, 3, 4, 5}},
		},
	}

	for _, data := range datasets {
		result, _ := SplitSlice(data.InputSlice, data.ChunkSize)
		require.Equalf(
			t, data.OutputSlice, result, "Test failed. Expected: %v, Given %v", data.OutputSlice, result,
		)
	}
}

func TestSplitSliceError(t *testing.T) {
	_, err := SplitSlice(nil, 2)
	require.Error(t, err, "Error does not raise")
}

func TestReverseKey(t *testing.T) {

	type Dataset struct {
		InputMap  map[int64]string
		OutputMap map[string]int64
	}

	datasets := []Dataset{
		{
			InputMap:  map[int64]string{1: "one", 2: "two", 3: "three"},
			OutputMap: map[string]int64{"one": 1, "two": 2, "three": 3},
		},
		{
			InputMap:  map[int64]string{},
			OutputMap: map[string]int64{},
		},
	}

	for _, data := range datasets {
		result, _ := ReverseKey(data.InputMap)
		require.Equalf(t, data.OutputMap, result, "Test failed. Expected: %v, Given %v", data.OutputMap, result)
	}
}

func TestErrorReverseKey(t *testing.T) {
	wrongMap := map[int64]string{1: "one", 2: "one"}
	_, err := ReverseKey(wrongMap)
	require.Error(t, err, "Error does not raise")
}

func TestFilterSlice(t *testing.T) {
	type Dataset struct {
		InputSlice  []int64
		FilterSlice []int64
		OutputSlice []int64
	}

	datasets := []Dataset{
		{
			InputSlice:  []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			FilterSlice: []int64{1, 3, 5, 7, 9},
			OutputSlice: []int64{2, 4, 6, 8},
		},
		{
			InputSlice:  []int64{1, 1, 3, 3, 4, 4},
			FilterSlice: []int64{1, 3, 5, 7, 9},
			OutputSlice: []int64{4, 4},
		},
		{
			InputSlice:  []int64{1, 2, 3},
			FilterSlice: []int64{},
			OutputSlice: []int64{1, 2, 3},
		},
		{
			InputSlice:  nil,
			FilterSlice: nil,
			OutputSlice: nil,
		},
	}

	for _, data := range datasets {
		result, _ := FilterSlice(data.InputSlice, data.FilterSlice)
		require.Equalf(
			t, data.OutputSlice, result, "Test failed. Expected: %v, Given %v", data.OutputSlice, result,
		)
	}
}

func TestErrorFilterSlice(t *testing.T) {
	_, err := FilterSlice(nil, []int64{1, 2, 3})
	require.Error(t, err, "Error does not raise")

	_, err = FilterSlice([]int64{1, 2, 3}, nil)
	require.Error(t, err, "Error does not raise")
}
