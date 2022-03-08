package utils

import (
	"strconv"
	"strings"
)

func HexToInt64(value string) (int64, error) {
	v, err := strconv.ParseInt(strings.Replace(value, "0x", "", -1), 16, 64)
	if err != nil {
		return 0, err
	}

	return v, nil
}
