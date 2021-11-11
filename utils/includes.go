package utils

import (
	"strings"
)

func Includes(array []string, target string) bool {
	found := false
	for _, x := range array {
		if x == target || strings.Contains(target, x) {
			found = true
		}
	}
	return found
}
