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

// String returns the environment value as a string.
// If environment variable with such a `key` is not
// present, `defval` is returned instead.
func String(key, defavlt string) string {
	return envValue(key, defavlt).(string)
}

// Bool returns the environment value as bool
func Bool(key string, defavlt bool) bool {
	return envValue(key, defavlt).(bool)
}

// Int returns an integer value from the environment
func Int(key string, defavlt int) int {
	return envValue(key, defavlt).(int)
}

// Int64 returns an int64 value from the environment
func Int64(key string, defavlt int64) int64 {
	return envValue(key, defavlt).(int64)
}

// Float returns float64 value from the environment
func Float(key string, defavlt float64) float64 {
	return envValue(key, defavlt).(float64)
}

// Duration returns time.Duration value from the environment
func Duration(key string, defavlt time.Duration) time.Duration {
	return envValue(key, defavlt).(time.Duration)
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
