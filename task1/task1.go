package task1

import (
	"strconv"
	"unicode/utf8"
)

func contains(val string) bool {
	for _, a := range []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"} {
		if val == a {
			return true
		}
	}
	return false
}

func getIntStr(s string) string {
	if utf8.RuneCountInString(s) == 0 {
		return "0"
	}
	res := ""
	for _, r := range s {
		if contains(string(r)) {
			res = res + string(r)
			continue
		}
		if !contains(string(r)) && utf8.RuneCountInString(res) == 0 {
			continue
		}
		if !contains(string(r)) && utf8.RuneCountInString(res) > 0 {
			return res
		}
	}
	return res
}

func StrToInt(s string) (int, error) {
	line := getIntStr(s)
	v, err := strconv.Atoi(line)
	if err != nil {
		return 0, err
	}
	return v, nil
}