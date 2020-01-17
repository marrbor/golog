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

	LEVEL_ENV = "GOLOG_LEVEL"
)

// filter level
var (
	filter                = WARN
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

// canOut check whether filter or not
func canOut(level int) bool {
	return filter <= level
}

// checkLevel checks whether given number is within log level
func checkLevel(level int) bool {
	return MIN <= level && level <= MAX
}

// prefix set prefix strings [level]
func prefix(level int) {
	logger.SetPrefix(fmt.Sprintf("[%-5s] ", levelStr[level]))
}

// GetFilterLevel returns current filtering level
func GetFilterLevel() int {
	return filter
}

// setFilterLevel sets filtering level
func SetFilterLevel(level int) error {
	if !checkLevel(level) {
		return OutOfLevelRangeError
	}
	filter = level
	return nil
}

// LoadFilterLevel loads filtering level from environment variable.
func LoadFilterLevel() error {
	le := os.Getenv(LEVEL_ENV)
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

// Output out log when the specified level is greater or equal than filtering level and return whether out log or not.
func Output(level, skip int, msg string) bool {
	if !canOut(level) {
		return false
	}

	prefix(level)
	if err := logger.Output(skip, msg); err != nil {
		return false
	}
	return true
}

func Trace(skip int, msg string) bool {
	return Output(TRACE, skip, msg)
}

func Debug(skip int, msg string) bool {
	return Output(DEBUG, skip, msg)
}

func Info(skip int, msg string) bool {
	return Output(INFO, skip, msg)
}

func Warn(skip int, msg string) bool {
	return Output(WARN, skip, msg)
}

func Error(skip int, msg string) bool {
	return Output(ERROR, skip, msg)
}

func Fatal(msg string) bool {
	if !canOut(FATAL) {
		return false
	}
	prefix(FATAL)
	logger.Fatal(msg)
	return true
}

func Panic(msg string) bool {
	prefix(PANIC)
	logger.Panic(msg)
	return true
}
