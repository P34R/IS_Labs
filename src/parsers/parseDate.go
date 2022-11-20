package parsers

import (
	"strconv"
	"strings"
)

func ParsePeriod(str string) (integers []int) {
	strs := strings.Split(str, "-")
	if len(strs) == 1 {
		nums := strings.Split(str, ".")
		for _, b := range nums {
			num, err := strconv.Atoi(b)
			if err == nil {
				integers = append(integers, num)
			}
		}
	} else if len(strs) == 2 {
		for i := range strs {
			nums := strings.Split(strs[i], ".")
			for _, b := range nums {
				num, err := strconv.Atoi(b)
				if err == nil {
					integers = append(integers, num)
				}
			}
		}
	} else {
		return nil
	}
	return
}
