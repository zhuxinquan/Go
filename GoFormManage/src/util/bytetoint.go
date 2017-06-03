package util

import (
	"strconv"
)

func ByteToInt(b []byte) int {

	x, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}
	return x
}
