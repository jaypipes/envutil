package envutil

import (
	"os"
	"strconv"
	"time"
)

// Returns the string value of the supplied environ variable or, if not
// present, the supplied default value
func WithDefault(key string, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}

// Returns the int value of the supplied environ variable or, if not present,
// the supplied default value. If the int conversion fails, returns the default
func WithDefaultInt(key string, def int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return def
	}
	return i
}

// Returns the bool value of the supplied environ variable or, if not present,
// the supplied default value. If the conversion fails, returns the default
func WithDefaultBool(key string, def bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		return def
	}
	return b
}

// Returns the duration value of the supplied environ variable or, if not present,
// the supplied default value. If the conversion fails, returns the default
func WithDefaultDuration(key string, def time.Duration) time.Duration {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	d, err := time.ParseDuration(val)
	if err != nil {
		return def
	}
	return d
}
