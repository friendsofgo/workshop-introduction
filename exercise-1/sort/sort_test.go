package sort

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_Nums(t *testing.T) {
	input := []int{11, 76, 21, 51, 70, 7, 99, 69, 17, 63}
	expectedOutput := []int{7, 11, 17, 21, 51, 63, 69, 70, 76, 99}

	assert.True(t, reflect.DeepEqual(Nums(input), expectedOutput), "The sort output should be sorted.")
}