default: all

env_tests: test_env_trace test_env_debug test_env_info test_env_warn test_env_error test_env_fatal test_env_panic
all: $(env_tests) test_fatal test_panic

## environment variable read test
test_env_trace:
	GOLOG_LEVEL=TRACE go test  --run TestEnv

test_env_debug:
	GOLOG_LEVEL=DEBUG go test  --run TestEnv

test_env_info:
	GOLOG_LEVEL=INFO go test  --run TestEnv

test_env_warn:
	GOLOG_LEVEL=WARN go test  --run TestEnv

test_env_error:
	GOLOG_LEVEL=ERROR go test  --run TestEnv

test_env_fatal:
	GOLOG_LEVEL=FATAL go test  --run TestEnv

test_env_panic:
	GOLOG_LEVEL=PANIC go test  --run TestEnv

## Process stop test.
test_fatal:
	-go test --run TestFatal
	-go test --run TestFatal2
	-go test --run TestFatal3
	-go test --run TestFatal4

test_panic:
	-go test --run TestPanic
	-go test --run TestPanic2
