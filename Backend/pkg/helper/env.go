package helper

import "os"

func EnvOrDefault(key, default_ string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return default_
}
