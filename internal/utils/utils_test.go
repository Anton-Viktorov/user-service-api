package utils

import (
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSplitSlice(t *testing.T) {
	t.Run("accurate split", func(t *testing.T) {
		input := []int64{1, 2, 3, 4}
		expected := [][]int64{{1, 2}, {3, 4}}
		res, err := SplitSlice(input, 2)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("last slice is less than chunk size", func(t *testing.T) {
		input := []int64{1, 2, 3, 4, 5}
		expected := [][]int64{{1, 2, 3}, {4, 5}}
		res, err := SplitSlice(input, 3)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("chunk size is equal slice len", func(t *testing.T) {
		input := []int64{1, 2, 3, 4, 5}
		expected := [][]int64{{1, 2, 3, 4, 5}}
		res, err := SplitSlice(input, 5)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("chunk size is larger than slice len", func(t *testing.T) {
		input := []int64{1, 2, 3, 4, 5}
		expected := [][]int64{{1, 2, 3, 4, 5}}
		res, err := SplitSlice(input, 10)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("nil data raises error", func(t *testing.T) {
		_, err := SplitSlice(nil, 2)
		require.Error(t, err)
		require.Equal(t, errors.New("input argument invalid"), err)
	})

}

func TestReverseKey(t *testing.T) {
	t.Run("base success case", func(t *testing.T) {
		input := map[int64]string{1: "one", 2: "two", 3: "three"}
		expected := map[string]int64{"one": 1, "two": 2, "three": 3}
		res, err := ReverseKey(input)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("empty map", func(t *testing.T) {
		input := map[int64]string{}
		expected := map[string]int64{}
		res, err := ReverseKey(input)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("duplicate key raises error", func(t *testing.T) {
		wrongData := map[int64]string{1: "one", 2: "one"}
		_, err := ReverseKey(wrongData)
		require.Error(t, err)
		require.Equal(t, errors.New("duplicate key: one"), err)
	})

	t.Run("nil data raises error", func(t *testing.T) {
		_, err := ReverseKey(nil)
		require.Error(t, err)
		require.Equal(t, errors.New("input argument invalid"), err)
	})

}

func TestFilterSlice(t *testing.T) {
	t.Run("base success case", func(t *testing.T) {
		input := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
		filter := []int64{1, 3, 5, 7, 9}
		expected := []int64{2, 4, 6, 8}
		res, err := FilterSlice(input, filter)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("filter all data items", func(t *testing.T) {
		input := []int64{1, 2, 3}
		filter := []int64{1, 2, 3}
		var expected []int64
		res, err := FilterSlice(input, filter)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("empty filter", func(t *testing.T) {
		input := []int64{1, 2, 3}
		filter := make([]int64, 0) // because we need empty slice [], not a nil-slice
		expected := []int64{1, 2, 3}
		res, err := FilterSlice(input, filter)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("empty data", func(t *testing.T) {
		input := make([]int64, 0)
		filter := []int64{1, 2, 3}
		var expected []int64 // because inside a func res slice is declared like a nil-slice var res []int64
		res, err := FilterSlice(input, filter)
		require.NoError(t, err)
		require.Equal(t, expected, res)
	})

	t.Run("nil data raises error", func(t *testing.T) {
		_, err := FilterSlice(nil, []int64{1, 2, 3})
		require.Error(t, err)
		require.Equal(t, errors.New("input argument invalid"), err)
	})

	t.Run("nil filter raises error", func(t *testing.T) {
		_, err := FilterSlice([]int64{1, 2, 3}, nil)
		require.Error(t, err)
		require.Equal(t, errors.New("input argument invalid"), err)
	})

}
