default: all

env_tests: test_env_trace test_env_debug test_env_info test_env_warn test_env_error test_env_fatal test_env_panic
all: $(env_tests) test_fatal test_panic

## environment variable read test
test_env_trace:
	GOLOG_LEVEL=TRACE go test -cover --run TestEnv

test_env_debug:
	GOLOG_LEVEL=DEBUG go test -cover --run TestEnv

test_env_info:
	GOLOG_LEVEL=INFO go test -cover --run TestEnv

test_env_warn:
	GOLOG_LEVEL=WARN go test -cover --run TestEnv

test_env_error:
	GOLOG_LEVEL=ERROR go test -cover --run TestEnv

test_env_fatal:
	GOLOG_LEVEL=FATAL go test -cover --run TestEnv

test_env_panic:
	GOLOG_LEVEL=PANIC go test -cover --run TestEnv

## Process stop test.
test_fatal:
	-go test -cover --run TestFatal
	-go test -cover --run TestFatal2
	-go test -cover --run TestFatal3
	-go test -cover --run TestFatal4

test_panic:
	-go test -cover --run TestPanic
	-go test -cover --run TestPanic2
