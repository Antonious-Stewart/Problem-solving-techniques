package slidingwindow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSumOfSubArraySizeK(t *testing.T) {
	assertion := assert.New(t)
	arr := []int{2, 1, 5, 1, 3, 2}
	actual := MaxSumOfSubArraySizeK(arr, 3)

	assertion.Exactly(9, actual, "Value should be 9")
}
