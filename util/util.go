package util

import (
	"strings"
)

func ParseMultiImage(s string) []string {
	return strings.Split(s, ";")
}

func MergeMultiImage(slice []string) string {
	return strings.Join(slice, ";")
}

func GetFirstImage(s string) string {
	idx := strings.Index(s, ";")
	if idx == -1 {
		return s
	}

	return s[:idx]
}
