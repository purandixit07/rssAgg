package main

import (
	"strconv"
)

func unitTestFunc(number int) string {
	isTrue := number%3 == 0

	if isTrue {
		return "correct"
	}

	return strconv.Itoa(number)

}
