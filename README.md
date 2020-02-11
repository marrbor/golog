# golog

`import "github.com/marrbor/golog"`

Add level filtering to [standard log package](https://golang.org/pkg/log/).

This module holds a `log.Logger` instance inside.

## Usage

```go
package main
import (
    "os"
    "github.com/marrbor/golog"
)

func main() {
    // log level is read from environment variable "GOLOG_LEVEL" (if specified) when golog module initialized.

	// change log level for this application if needed.
	if err := golog.SetFilterLevel(golog.INFO); err != nil {
		panic(err)
	}

    golog.Info("Application start")

    golog.Info("Application finish.")
    os.Exit(0)
}
```
## Filter levels

Use following seven levels, shown from low to high. When filter level is "WARN", logger will outputs "WARN","ERROR","FATAL","PANIC" four levels that greater or equal to "WARN".

1. `TRACE`
1. `DEBUG`
1. `INFO`
1. `WARN`
1. `ERROR`
1. `FATAL` application exit with `1` after logging.
1. `PANIC` application panic after logging.

Default filter level is `WARN`. If you want to modify, use `SetFilterLevel` or `LoadFilterLevel`

## Environment Variables

### GOLOG_LEVEL
Set level names defined before e.g. `export GOLOG_LEVEL=DEBUG`. This is not reflect automatically.

## Preferences

golog use following preferences at start. If you want to modify them, get logger instance by `golog.GetLogger()` and call [standard log package](https://golang.org/pkg/log/) methods for the instance except log level.

Example for change log prefix:

```go
l := golog.GetLogger()
l.SetPrefix("My Application: ")
```

|Item|Default Value|To modify|
|:---|:---|:---|
|Destination|os.Stderr|use SetOutput()|
|Flags|Ldate+Ltime+Lshortfile|use SetFlags()|
|Prefix|""|use SetPrefix()|

## License
MIT
