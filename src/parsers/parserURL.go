package parsers

import (
	"strconv"
	"strings"
)

func GetIntsFromURL(str string) (integers []int) {
	strs := strings.Split(str, "/")
	for _, b := range strs {
		num, err := strconv.Atoi(b)
		if err == nil {
			integers = append(integers, num)
		}
	}
	return integers
}
