# OS Environment

[![Go Reference](https://pkg.go.dev/badge/github.com/rusq/osenv.svg)][1]

This package aims to provide extension of the `os.Getenv` that
supports different types of variable types, and allows to specify a
default value, in case the environment variable is missing.

This version is implemented with [Go 1.18+ generics][2].  If you need
to use the package's functionailty in Go versions prior to 1.18, see
below.

There are two versions of the package:

- v2: uses generics, therefore can only be compiled with Go 1.18+
  (this version).
- v1: supports Go versions 1.12+.

Go 1.18+ (v2):
```go
import "github.com/rusq/osenv/v2"
```

Go 1.12+ (v1):
```go
import "github.com/rusq/osenv"
```

## Usage

```go
package main

import (
	"fmt"
	"math"
	"time"

	"github.com/rusq/osenv/v2"
)

func main() {
	fmt.Printf("OSENV_BOOL: %v\n"+
		"OSENV_DURATION: %s\n"+
		"OSENV_FLOAT: %.7f\n"+
		"OSENV_INT: %d\n"+
		"OSENV_INT64: %d\n"+
		"OSENV_STRING: %s\n"+
		"OSENV_TIME: %s\n",

		osenv.Value("OSENV_BOOL", true),
		osenv.Value("OSENV_DURATION", 60*time.Second),
		osenv.Value("OSENV_FLOAT", math.Pi),
		osenv.Value("OSENV_INT", math.MaxInt32),
		osenv.Value("OSENV_INT64", math.MaxInt64),
		osenv.Value("OSENV_STRING", "default string value"),
		osenv.Value("OSENV_TIME", time.Date(2020, 12, 31, 23, 59, 59, 0, time.UTC)),
	)
}
```

For more examples, refer to [package documentation][1]

The package defines a single function that handles all supported
types:

- Boolean
- time.Duration 
- Float64
- Int
- Int64
- String
- Secret*
- Time

For all of these types, call `osenv.Value(...)`.

\* Secret obtains a string from the environment (`osenv.Secret`).  The
only difference between `Value` and `Secret`, is that Secret unsets
the environment variable after getting the value.

## Upgrading from v1 to v2

If you were using osenv/v1 before, and upgrading to osenv/v2, follow
the instructions in this section.

Replace the call to the function in "From" with the call to the
function in "To":

| From           | To          |
|----------------|-------------|
| osenv.Bool     | osenv.Value |
| osenv.Duration | osenv.Value |
| osenv.Float    | osenv.Value |
| osenv.Int      | osenv.Value |
| osenv.Int64    | osenv.Value |
| osenv.Time     | osenv.Value |
| osenv.String   | osenv.Value |

[1]: https://pkg.go.dev/github.com/rusq/osenv/v2
[2]: https://go.dev/blog/intro-generics
