package golog_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/marrbor/golog"
	"github.com/stretchr/testify/assert"
)

func TestGetLogger(t *testing.T) {
	os.Clearenv()
	l := golog.GetLogger()
	oldflg := l.Flags()
	l.SetFlags(0)
	flg := l.Flags()
	assert.EqualValues(t, 0, flg)
	l.SetFlags(oldflg)
	assert.EqualValues(t, oldflg, l.Flags())
}

func TestGetFilterLevel(t *testing.T) {
	os.Clearenv()
	lvl := golog.GetFilterLevel()
	assert.EqualValues(t, golog.INFO, lvl)
}

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
	_ = os.Setenv(golog.LevelEnv, "abc")
	err = golog.LoadFilterLevel()
	assert.EqualError(t, err, golog.InvalidLevelNameError.Error())

	_ = os.Setenv(golog.LevelEnv, fmt.Sprintf("%d", golog.MAX+1))
	err = golog.LoadFilterLevel()
	assert.EqualError(t, err, golog.InvalidLevelNameError.Error())

	_ = os.Setenv(golog.LevelEnv, fmt.Sprintf("%d", golog.MIN-1))
	err = golog.LoadFilterLevel()
	assert.EqualError(t, err, golog.InvalidLevelNameError.Error())

	_ = os.Setenv(golog.LevelEnv, fmt.Sprintf("%d", golog.MIN))
	err = golog.LoadFilterLevel()
	assert.EqualError(t, err, golog.InvalidLevelNameError.Error())

	_ = os.Setenv(golog.LevelEnv, fmt.Sprintf("%d", golog.MAX))
	err = golog.LoadFilterLevel()
	assert.EqualError(t, err, golog.InvalidLevelNameError.Error())

	// valid
	_ = os.Setenv(golog.LevelEnv, "TRACE")
	err = golog.LoadFilterLevel()
	assert.NoError(t, err)
	assert.EqualValues(t, golog.TRACE, golog.GetFilterLevel())

	_ = os.Setenv(golog.LevelEnv, "DEBUG")
	err = golog.LoadFilterLevel()
	assert.NoError(t, err)
	assert.EqualValues(t, golog.DEBUG, golog.GetFilterLevel())

	_ = os.Setenv(golog.LevelEnv, "INFO")
	err = golog.LoadFilterLevel()
	assert.NoError(t, err)
	assert.EqualValues(t, golog.INFO, golog.GetFilterLevel())

	_ = os.Setenv(golog.LevelEnv, "WARN")
	err = golog.LoadFilterLevel()
	assert.NoError(t, err)
	assert.EqualValues(t, golog.WARN, golog.GetFilterLevel())

	_ = os.Setenv(golog.LevelEnv, "ERROR")
	err = golog.LoadFilterLevel()
	assert.NoError(t, err)
	assert.EqualValues(t, golog.ERROR, golog.GetFilterLevel())

	_ = os.Setenv(golog.LevelEnv, "FATAL")
	err = golog.LoadFilterLevel()
	assert.NoError(t, err)
	assert.EqualValues(t, golog.FATAL, golog.GetFilterLevel())

	_ = os.Setenv(golog.LevelEnv, "PANIC")
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
	assert.False(t, golog.Output(golog.TRACE, "golog trace"))
	assert.True(t, golog.Output(golog.DEBUG, "golog debug", 1))
	assert.True(t, golog.Output(golog.INFO, "golog info", 1))
}

func TestTrace(t *testing.T) {
	_ = golog.SetFilterLevel(golog.TRACE)
	assert.True(t, golog.Trace("golog trace"))
	assert.True(t, golog.Trace("golog trace with depth", 1))
	_ = golog.SetFilterLevel(golog.DEBUG)
	assert.False(t, golog.Trace("golog trace"))
	assert.False(t, golog.Trace("golog trace with depth", 1))
}

func TestDebug(t *testing.T) {
	_ = golog.SetFilterLevel(golog.DEBUG)
	assert.True(t, golog.Debug("golog debug"))
	assert.True(t, golog.Debug("golog debug with depth", 1))
	_ = golog.SetFilterLevel(golog.INFO)
	assert.False(t, golog.Debug("golog debug"))
	assert.False(t, golog.Debug("golog debug with depth", 1))
}

func TestInfo(t *testing.T) {
	_ = golog.SetFilterLevel(golog.INFO)
	assert.True(t, golog.Info("golog info"))
	assert.True(t, golog.Info("golog info with depth", 1))
	_ = golog.SetFilterLevel(golog.WARN)
	assert.False(t, golog.Info("golog info"))
	assert.False(t, golog.Info("golog info with depth", 1))
}

func TestWarn(t *testing.T) {
	_ = golog.SetFilterLevel(golog.WARN)
	assert.True(t, golog.Warn("golog warn"))
	assert.True(t, golog.Warn("golog warn with depth", 1))
	_ = golog.SetFilterLevel(golog.ERROR)
	assert.False(t, golog.Warn("golog warn"))
	assert.False(t, golog.Warn("golog warn with depth", 1))
}

func TestError(t *testing.T) {
	_ = golog.SetFilterLevel(golog.ERROR)
	assert.True(t, golog.Error("golog error"))
	assert.True(t, golog.Error("golog error with depth", 1))
	_ = golog.SetFilterLevel(golog.FATAL)
	assert.False(t, golog.Error("golog error"))
	assert.False(t, golog.Error("golog error with depth"))
}
