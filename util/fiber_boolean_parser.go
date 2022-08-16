package util

import "strconv"

func ParseBoolean(value string) bool {
	var err error
	var result bool
	if value == "" {
		result = false
	} else {
		result, err = strconv.ParseBool(value)
		if err != nil {
			result = false
		}
	}
	return result
}
