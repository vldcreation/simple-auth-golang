package utstring

import (
	"os"
	"strconv"
)

func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func Uint64ToString(v int64) string {
	return strconv.FormatInt(int64(v), 10)
}

// Env to get environment variable
func Env(key string, def ...string) string {
	return Chains(append([]string{os.Getenv(key)}, def...)...)
}

// Chains function
func Chains(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}
