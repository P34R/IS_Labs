package parsers

import (
	"errors"
	"strconv"
	"strings"
)

type Code int

const (
	ERROR   Code = 0
	MORE         = 1
	LESS         = 2
	BETWEEN      = 3
	EQUAL        = 4
)

func ParseIntegerSearch(str string) (code Code, ints []int, err error) {
	if str[0] == '>' {
		i, err := strconv.Atoi(str[1:])
		if err != nil {
			return ERROR, nil, err
		}
		ints = append(ints, i)
		return MORE, ints, nil
	}
	if str[0] == '<' {
		i, err := strconv.Atoi(str[1:])
		if err != nil {
			return ERROR, nil, err
		}
		ints = append(ints, i)
		return LESS, ints, nil
	}
	if a, err := strconv.Atoi(str); err == nil {
		ints = append(ints, a)
		return EQUAL, ints, nil
	}
	if ind := strings.Index(str, "-"); ind != -1 {
		a, err := strconv.Atoi(str[:ind])
		if err != nil {
			return ERROR, nil, err
		}
		b, err := strconv.Atoi(str[ind+1:])
		if err != nil {
			return ERROR, nil, err
		}
		ints = append(ints, a, b)
		return BETWEEN, ints, nil
	}
	return 0, nil, errors.New("error: invalid syntax")
}
