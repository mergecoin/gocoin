package utils

import (
	"strings"
)

func Includes(target string, args ...string) bool {
	found := false

	for _, x := range args {
		if x == target || strings.Contains(target, x) {
			found = true
		}
	}
	return found
}
