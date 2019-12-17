package task2

import (
	"fmt"
	"strconv"
)

func StrToInt0(s string) (int, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func StrToInt(s string) (int, error) {
	var res int
	_, err := fmt.Sscanf(s, "%d", &res)
	return res, err
}
