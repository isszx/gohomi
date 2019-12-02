package task2

import "fmt"

func StrToInt(s string) (int, error) {
	var res int
	_, err := fmt.Sscanf(s, "%d", &res)
	return res, err
}