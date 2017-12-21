package util

import (
	"strconv"
)

func Atoi(s string) int {
	var err error
	var i int
	i, err = strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
