package golog_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/marrbor/golog"
	"github.com/stretchr/testify/assert"
)

func TestSetFilterLevel(t *testing.T) {
	err := golog.SetFilterLevel(-1)
	assert.EqualError(t, err, golog.OutOfLevelRangeError.Error())

	err = golog.SetFilterLevel(golog.MAX + 1)
	assert.EqualError(t, err, golog.OutOfLevelRangeError.Error())

	err = golog.SetFilterLevel(golog.MIN)
	assert.Nil(t, err)
	assert.EqualValues(t, golog.MIN, golog.GetFilterLevel())

	err = golog.SetFilterLevel(golog.MAX)
	assert.Nil(t, err)
	assert.EqualValues(t, golog.MAX, golog.GetFilterLevel())
}

func TestLoadFilterLevel(t *testing.T) {
	err := golog.SetFilterLevel(golog.MAX)
	assert.Nil(t, err)
	assert.EqualValues(t, golog.MAX, golog.GetFilterLevel())

	// no effect since environment variables not used.
	os.Clearenv()
	err = golog.LoadFilterLevel()
	assert.Nil(t, err)
	assert.EqualValues(t, golog.MAX, golog.GetFilterLevel())

	// error case
	_ = os.Setenv(golog.LEVEL_ENV, "abc")
	err = golog.LoadFilterLevel()
	assert.EqualError(t, err, golog.InvalidLevelNameError.Error())

	_ = os.Setenv(golog.LEVEL_ENV, fmt.Sprintf("%d", golog.MAX+1))
	err = golog.LoadFilterLevel()
	assert.EqualError(t, err, golog.InvalidLevelNameError.Error())

	_ = os.Setenv(golog.LEVEL_ENV, fmt.Sprintf("%d", golog.MIN-1))
	err = golog.LoadFilterLevel()
	assert.EqualError(t, err, golog.InvalidLevelNameError.Error())

	_ = os.Setenv(golog.LEVEL_ENV, fmt.Sprintf("%d", golog.MIN))
	err = golog.LoadFilterLevel()
	assert.EqualError(t, err, golog.InvalidLevelNameError.Error())

	_ = os.Setenv(golog.LEVEL_ENV, fmt.Sprintf("%d", golog.MAX))
	err = golog.LoadFilterLevel()
	assert.EqualError(t, err, golog.InvalidLevelNameError.Error())

	// valid
	_ = os.Setenv(golog.LEVEL_ENV, "TRACE")
	err = golog.LoadFilterLevel()
	assert.NoError(t, err)
	assert.EqualValues(t, golog.TRACE, golog.GetFilterLevel())

	_ = os.Setenv(golog.LEVEL_ENV, "DEBUG")
	err = golog.LoadFilterLevel()
	assert.NoError(t, err)
	assert.EqualValues(t, golog.DEBUG, golog.GetFilterLevel())

	_ = os.Setenv(golog.LEVEL_ENV, "INFO")
	err = golog.LoadFilterLevel()
	assert.NoError(t, err)
	assert.EqualValues(t, golog.INFO, golog.GetFilterLevel())

	_ = os.Setenv(golog.LEVEL_ENV, "WARN")
	err = golog.LoadFilterLevel()
	assert.NoError(t, err)
	assert.EqualValues(t, golog.WARN, golog.GetFilterLevel())

	_ = os.Setenv(golog.LEVEL_ENV, "ERROR")
	err = golog.LoadFilterLevel()
	assert.NoError(t, err)
	assert.EqualValues(t, golog.ERROR, golog.GetFilterLevel())

	_ = os.Setenv(golog.LEVEL_ENV, "FATAL")
	err = golog.LoadFilterLevel()
	assert.NoError(t, err)
	assert.EqualValues(t, golog.FATAL, golog.GetFilterLevel())

	_ = os.Setenv(golog.LEVEL_ENV, "PANIC")
	err = golog.LoadFilterLevel()
	assert.NoError(t, err)
	assert.EqualValues(t, golog.PANIC, golog.GetFilterLevel())

	// again, test when variable is blank.
	err = golog.SetFilterLevel(golog.MAX)
	os.Clearenv()
	assert.Nil(t, err)
	assert.EqualValues(t, golog.MAX, golog.GetFilterLevel())
}

func TestOutput(t *testing.T) {
	_ = golog.SetFilterLevel(golog.DEBUG)
	assert.False(t, golog.Output(golog.TRACE, 0, "golog trace"))
	assert.True(t, golog.Output(golog.DEBUG, 1, "golog debug"))
	assert.True(t, golog.Output(golog.INFO, 2, "golog info"))
}

func TestTrace(t *testing.T) {
	_ = golog.SetFilterLevel(golog.TRACE)
	assert.True(t, golog.Trace(0, "golog trace"))
	_ = golog.SetFilterLevel(golog.DEBUG)
	assert.False(t, golog.Trace(0, "golog trace"))
}

func TestDebug(t *testing.T) {
	_ = golog.SetFilterLevel(golog.DEBUG)
	assert.True(t, golog.Debug(0, "golog debug"))
	_ = golog.SetFilterLevel(golog.INFO)
	assert.False(t, golog.Debug(0, "golog debug"))
}

func TestInfo(t *testing.T) {
	_ = golog.SetFilterLevel(golog.INFO)
	assert.True(t, golog.Info(0, "golog info"))
	_ = golog.SetFilterLevel(golog.WARN)
	assert.False(t, golog.Info(0, "golog info"))
}

func TestWarn(t *testing.T) {
	_ = golog.SetFilterLevel(golog.WARN)
	assert.True(t, golog.Warn(0, "golog warn"))
	_ = golog.SetFilterLevel(golog.ERROR)
	assert.False(t, golog.Warn(0, "golog warn"))
}

func TestError(t *testing.T) {
	_ = golog.SetFilterLevel(golog.ERROR)
	assert.True(t, golog.Error(0, "golog error"))
	_ = golog.SetFilterLevel(golog.FATAL)
	assert.False(t, golog.Error(0, "golog error"))
}

func TestFatal(t *testing.T) {

}

func TestPanic(t *testing.T) {

}
