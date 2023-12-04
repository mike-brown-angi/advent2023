package advcommon

import (
	"github.com/bmizerany/assert"
	"math"
	"math/rand"
	"testing"
)

func TestSimpleGeneric(t *testing.T) {
	s := SimpleGeneric([]int{1}, []int{2})
	assert.Equal(t, len(s), 0)
	assert.Equal(t, s, []int{})

	s = SimpleGeneric([]int{1, 2}, []int{2})
	assert.Equal(t, s, []int{2})
}

func TestSortedGeneric(t *testing.T) {
	s := SortedGeneric([]int{1}, []int{2})
	assert.Equal(t, len(s), 0)
	assert.Equal(t, s, []int{})

	s = SortedGeneric([]int{1, 2}, []int{2})
	assert.Equal(t, s, []int{2})
}

func TestHashGeneric(t *testing.T) {
	s := HashGeneric([]int{1}, []int{2})
	assert.Equal(t, len(s), 0)
	assert.Equal(t, s, []int{})

	s = HashGeneric([]int{1, 2}, []int{2})
	assert.Equal(t, s, []int{2})
}

func createRandomSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(int(math.Pow(float64(size), 2)))
	}
	return slice
}
