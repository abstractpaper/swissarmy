package swissarmy

import (
	"os"
)

// GetEnv gets the environment variable `key` value if it
// exists, otherwise return default.
func GetEnv(key string, def string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return def
}