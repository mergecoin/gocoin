package utils

import (
	"strings"
	"fmt"
)

func Includes(target string, args ...string) bool {
	found := false

	for _, x := range args {
		if x == target || strings.Contains(target, x) {
			fmt.Println(x, target)
			found = true
		}
	}
	return found
}
