package swap

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_Nums(t *testing.T) {
	input := []int{11, 76, 21, 51, 70, 7, 99, 69, 17, 63, 55}
	expectedOutput := []int{76, 11, 51, 21, 7, 70, 69, 99, 63, 17, 55}

	assert.True(t, reflect.DeepEqual(Nums(input), expectedOutput), "`Nums()` output should be swapped.")
}