package util

import (
	"os"
	"strconv"
)

func GetString(key string) string {
	return os.Getenv(key)
}

func GetInt(key string) (int64, error) {
	v := GetString(key)
	i, err := strconv.ParseInt(v, 10, 0)
	if err != nil {
		return 0, err
	}

	return i, nil
}
