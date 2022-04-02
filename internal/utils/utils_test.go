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
		{
			ChunkSize:   2,
			InputSlice:  nil,
			OutputSlice: nil,
		},
	}

	for _, data := range datasets {
		result := SplitSlice(data.InputSlice, data.ChunkSize)
		require.Equalf(
			t, data.OutputSlice, result, "Test failed. Expected: %v, Given %v", data.OutputSlice, result,
		)
	}
}

func TestReverseKey(t *testing.T) {

	type Dataset struct {
		InputMap  map[int64]int64
		OutputMap map[int64]int64
	}

	datasets := []Dataset{
		{
			InputMap:  map[int64]int64{1: 12, 2: 22, 3: 32},
			OutputMap: map[int64]int64{12: 1, 22: 2, 32: 3},
		},
		{
			InputMap:  map[int64]int64{},
			OutputMap: map[int64]int64{},
		},
		{
			InputMap:  nil,
			OutputMap: nil,
		},
	}

	for _, data := range datasets {
		result := ReverseKey(data.InputMap)
		require.Equalf(t, data.OutputMap, result, "Test failed. Expected: %v, Given %v", data.OutputMap, result)
	}
}

func TestPanicReverseKey(t *testing.T) {
	wrongMap := map[int64]int64{1: 10, 2: 10}
	require.Panicsf(t, func() { ReverseKey(wrongMap) }, "Panic does not raise")
}

func TestFilterSlice(t *testing.T) {
	type Dataset struct {
		InputSlice  []int64
		OutputSlice []int64
	}

	datasets := []Dataset{
		{
			InputSlice:  []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			OutputSlice: []int64{2, 4, 6, 8},
		},
		{
			InputSlice:  []int64{1, 1, 3, 3, 4, 4},
			OutputSlice: []int64{4, 4},
		},
		{
			InputSlice:  []int64{},
			OutputSlice: []int64{},
		},
		{
			InputSlice:  nil,
			OutputSlice: nil,
		},
	}

	for _, data := range datasets {
		result := FilterSlice(data.InputSlice)
		require.Equalf(
			t, data.OutputSlice, result, "Test failed. Expected: %v, Given %v", data.OutputSlice, result,
		)
	}
}
