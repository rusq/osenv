// Package osenv provides convenient functions to access
// environment variables.
//
// Copyright 2019 Rustam Gilyazov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package osenv

import (
	"os"
	"strconv"
	"time"
)

func Value[T int | int64 | string | float64 | time.Duration | bool](key string, defavlt T) T {
	v := envValue(key, defavlt)
	return v.(T)
}

// Secret returns the value of the environment variable with the name `key`.  If
// the environment variable with name KEY not found, it returns the `defavlt` value.
// The environment variable is unset after the value is retrieved.
func Secret[T int | int64 | string | float64 | time.Duration | bool](key string, defavlt T) T {
	v := Value(key, defavlt)
	os.Unsetenv(key)
	return v
}

func envValue(key string, defval interface{}) interface{} {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defval
	}

	switch defval.(type) {
	default:
		panic("unsupported")
	case string:
		return val
	case bool:
		r, err := strconv.ParseBool(val)
		if err == nil {
			return r
		}
	case int:
		r, err := strconv.Atoi(val)
		if err == nil {
			return r
		}
	case int64:
		r, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			return r
		}
	case float64:
		r, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return r
		}
	case time.Duration:
		r, err := time.ParseDuration(val)
		if err == nil {
			return r
		}
	}
	return defval
}
