package golog

import (
	"fmt"
	"log"
	"os"
)

// Logging level
const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	PANIC

	MIN = TRACE
	MAX = PANIC

	LevelEnv = "GOLOG_LEVEL"

	// DefaultcalldepthCount is the number of skipped callstack(s). The default value is 3.
	// (2 (function nesting in the "log" package) + 1 (function nesting in this package))
	DefaultCallDepth = 2 + 1
)

// filter level
var (
	filter                = INFO
	OutOfLevelRangeError  = fmt.Errorf("spacified level is not within log level range")
	InvalidLevelNameError = fmt.Errorf("spacified level string is not correct level name")
	levelStr              = []string{
		"TRACE",
		"DEBUG",
		"INFO",
		"WARN",
		"ERROR",
		"FATAL",
		"PANIC",
	}

	logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
)

///// Local functions

// init loads filtering levels when specified.
func init() {
	le := os.Getenv(LevelEnv)
	if 0 < len(le) {
		for lv, str := range levelStr {
			if str == le {
				if err := SetFilterLevel(lv); err != nil {
					Output(ERROR, err.Error())
					return
				}
			}
		}
	}
}

// canOut check whether filter or not
func canOut(level int) bool {
	return filter <= level
}

// checkLevel checks whether given number is within log level
func checkLevel(level int) bool {
	return MIN <= level && level <= MAX
}

// form formed level and strings.
func form(level int, msg interface{}) string {
	return fmt.Sprintf("[%-5s] %+v", levelStr[level], msg)
}

///// Global functions (configuration)

// GetLogger returns logger instance to fix configuration for logger.
func GetLogger() *log.Logger {
	return logger
}

// GetFilterLevel returns current filtering level
func GetFilterLevel() int {
	return filter
}

// GetFilterLevelStr returns current filtering level
func GetFilterLevelStr(level int) (string, error) {
	if level < MIN || MAX < level {
		return "", OutOfLevelRangeError
	}
	return levelStr[level], nil
}

// SetFilterLevel sets filtering level
func SetFilterLevel(level int) error {
	if !checkLevel(level) {
		return OutOfLevelRangeError
	}
	filter = level
	return nil
}

// LoadFilterLevel loads filtering level from environment variable.
func LoadFilterLevel() error {
	le := os.Getenv(LevelEnv)
	if 0 < len(le) {
		for lv, str := range levelStr {
			if str == le {
				return SetFilterLevel(lv)
			}
		}
		return InvalidLevelNameError
	}
	return nil
}

///// Global Functions (logging)

// Output out log when the specified level is greater or equal than filtering level and return whether out log or not.
func Output(level int, msg interface{}, calldepth ...int) bool {
	if !canOut(level) {
		return false
	}

	depth := DefaultCallDepth
	if 0 < len(calldepth) {
		depth += calldepth[0]
	}

	if err := logger.Output(depth, form(level, msg)); err != nil {
		return false
	}
	return true
}

// Trace outs log when filter level is TRACE.
func Trace(msg interface{}, calldepth ...int) bool {
	if len(calldepth) <= 0 {
		return Output(TRACE, msg)
	}
	return Output(TRACE, msg, calldepth[0])
}

// Debug outs log when filter level is TRACE or DEBUG.
func Debug(msg interface{}, calldepth ...int) bool {
	if len(calldepth) <= 0 {
		return Output(DEBUG, msg)
	}
	return Output(DEBUG, msg, calldepth[0])
}

// Info outs log when filter level less or equal to INFO.
func Info(msg interface{}, calldepth ...int) bool {
	if len(calldepth) <= 0 {
		return Output(INFO, msg)
	}
	return Output(INFO, msg, calldepth[0])
}

// Warn outs log when filter level less or equal to WARN.
func Warn(msg interface{}, calldepth ...int) bool {
	if len(calldepth) <= 0 {
		return Output(WARN, msg)
	}
	return Output(WARN, msg, calldepth[0])
}

// Error outs log when filter level less or equal to ERROR.
func Error(msg interface{}, calldepth ...int) bool {
	if len(calldepth) <= 0 {
		return Output(ERROR, msg)
	}
	return Output(ERROR, msg, calldepth[0])
}

// Fatal exit application with code 1 after logging.
func Fatal(msg interface{}, calldepth ...int) {
	if !canOut(FATAL) {
		os.Exit(1)
	}

	if len(calldepth) <= 0 {
		_ = Output(FATAL, msg)
		os.Exit(1)
	}
	_ = Output(FATAL, msg, calldepth[0])
	os.Exit(1)
}

// Panic panics application after logging.
func Panic(msg interface{}, calldepth ...int) {
	if len(calldepth) <= 0 {
		_ = Output(PANIC, msg)
		panic(msg)
	}
	_ = Output(PANIC, msg, calldepth[0])
	panic(msg)
}
