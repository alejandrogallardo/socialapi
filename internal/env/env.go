package env

import (
	"os"
	"strconv"
)

func GetString(key string, fallback string) string {
	// val, ok := os.LookupEnv(key)
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	//fmt.Printf("%s=%s\n", key, val)
	return val
}

func GetInt(key string, fallback int) int {
	// val, ok := os.LookupEnv(key)
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return valAsInt
}
