# OS Environment

[![Go Reference](https://pkg.go.dev/badge/github.com/rusq/osenv.svg)][1]

This package aims to provide extension of the `os.Getenv` that supports
different types of variable types, and allows to specify a default value, in
case the environment variable is missing.

There are two versions of the package:

- v1: supports Go versions 1.12+ (this version)*
- v2: uses generics, therefore can only be compiled with Go 1.18+

\* v1 most likely will run on any version of go, but I only used it with 1.12+.

Go 1.12+ (v1):
```go
import "github.com/rusq/osenv"
```

Go 1.18+ (v1):
```go
import "github.com/rusq/osenv/v2"
```

## Usage

```go
import (
    "fmt"

    "github.com/rusq/osenv"
)

func main() {
	fmt.Printf("OSENV_BOOL: %v\n"+
		"OSENV_DURATION: %s\n"+
		"OSENV_FLOAT: %.7f\n"+
		"OSENV_INT: %d\n"+
		"OSENV_INT64: %d\n"+
		"OSENV_STRING: %s\n"+
		"OSENV_TIME: %s\n",

		osenv.Bool("OSENV_BOOL", true),
		osenv.Duration("OSENV_DURATION", 60*time.Second),
		osenv.Float("OSENV_FLOAT", 3.1415927),
		osenv.Int("OSENV_INT", 42),
		osenv.Int64("OSENV_INT64", 64),
		osenv.String("OSENV_STRING", "default string value"),
		osenv.Time("OSENV_TIME", time.Date(2020, 12, 31, 23, 59, 59, 0, time.UTC)),
	)
}
```

For more examples, refer to [package documentation][1]

The package defines a set of functions, each for the respective supported type:

- Boolean
- time.Duration 
- Float64
- Int
- Int64
- String
- Secret*
- Time

\* Secret obtains a string from the environment (`osenv.Secret`).  The only difference
between `String` and `Secret`, is that Secret unsets the environment variable after
getting the value.

[1]: https://pkg.go.dev/github.com/rusq/osenv
