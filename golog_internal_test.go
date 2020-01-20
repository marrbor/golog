package golog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFilterLevelStr(t *testing.T) {
	_, err := GetFilterLevelStr(-1)
	assert.EqualError(t, err, OutOfLevelRangeError.Error())

	for i := 0; i <= MAX; i++ {
		lStr, err := GetFilterLevelStr(i)
		assert.NoError(t, err)
		assert.EqualValues(t, levelStr[i], lStr)
	}

	_, err = GetFilterLevelStr(MAX + 1)
	assert.EqualError(t, err, OutOfLevelRangeError.Error())
}
