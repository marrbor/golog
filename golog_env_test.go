/*
 Test that have to be called independent. They called via `go test --run xxxx`.
*/
package golog

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	lvl := GetFilterLevel()
	lvlStr := os.Getenv(LevelEnv)
	assert.EqualValues(t, lvlStr, levelStr[lvl])
}

func TestFatal(t *testing.T) {
	_ = SetFilterLevel(PANIC)
	Fatal("golog fatal")
}

func TestFatal2(t *testing.T) {
	_ = SetFilterLevel(PANIC)
	Fatal("golog fatal with depth", 1)
}

func TestFatal3(t *testing.T) {
	_ = SetFilterLevel(FATAL)
	Fatal("golog fatal")
}

func TestFatal4(t *testing.T) {
	_ = SetFilterLevel(FATAL)
	Fatal("golog fatal with depth", 1)
}

func TestPanic(t *testing.T) {
	Panic("panic")
}

func TestPanic2(t *testing.T) {
	Panic("panic", 1)
}
