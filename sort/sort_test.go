package sorter

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	// Arrange
	ints := []int{9, 8, 7, 6, 5}

	// Act
	BubbleSort(ints)

	// Assert
	assert.NotNil(t, ints)
	assert.Equal(t, 5, len(ints))
	assert.Equal(t, 5, ints[0])
	assert.Equal(t, 6, ints[1])
	assert.Equal(t, 7, ints[2])
	assert.Equal(t, 8, ints[3])
	assert.Equal(t, 9, ints[4])
}

func TestBubbleSort_nilInput(t *testing.T) {
	var ints []int
	BubbleSort(ints)
	assert.Nil(t, ints)
}

func createSliceInts(n int) []int {
	result := make([]int, n)
	for i := 0; i < n; i++ {
		result[i] = n - (i + 1)
	}
	return result
}

func BenchmarkBubbleSort_10(b *testing.B) {
	ints := createSliceInts(10)
	for i := 0; i < b.N; i++ {
		BubbleSort(ints)
	}
}
func BenchmarkSort_10(b *testing.B) {
	ints := createSliceInts(10)
	for i := 0; i < b.N; i++ {
		sort.Ints(ints)
	}
}
func BenchmarkBubbleSort_1000(b *testing.B) {
	ints := createSliceInts(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(ints)
	}
}
func BenchmarkSort_1000(b *testing.B) {
	ints := createSliceInts(1000)
	for i := 0; i < b.N; i++ {
		sort.Ints(ints)
	}
}

func BenchmarkBubbleSort_50000(b *testing.B) {
	ints := createSliceInts(50000)
	for i := 0; i < b.N; i++ {
		BubbleSort(ints)
	}
}
func BenchmarkSort_50000(b *testing.B) {
	ints := createSliceInts(50000)
	for i := 0; i < b.N; i++ {
		sort.Ints(ints)
	}
}

func BenchmarkBubbleSort_100000(b *testing.B) {
	ints := createSliceInts(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(ints)
	}
}
func BenchmarkSort_100000(b *testing.B) {
	ints := createSliceInts(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(ints)
	}
}
