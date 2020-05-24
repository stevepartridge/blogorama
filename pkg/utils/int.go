package utils

import (
	"strconv"
	"strings"
)

func IntSliceToSeparatedString(list []int, sep string) string {
	result := make([]string, len(list))
	for i := range list {
		result[i] = strconv.Itoa(list[i])
	}

	return strings.Join(result, sep)
}
