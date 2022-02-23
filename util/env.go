package util

import (
	"os"
	"strconv"
)

func GetEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	return value
}

func GetEnvInt64OrDefault(key string, defaultValue int64) (int64, error) {
	v := os.Getenv(key)
	if v == "" {
		return defaultValue, nil
	}
	i, err := strconv.ParseInt(v, 10, 0)
	if err != nil {
		return 0, err
	}

	return i, nil
}
