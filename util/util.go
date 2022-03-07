package util

import (
	"fmt"
	"strings"
)

func ParseMultiImage(s string) []string {
	if s == "" {
		return make([]string, 0)
	}
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

func GenOrderCode(id uint32) (s string) {
	timestamp := GetCurrentTime().Unix()
	s = fmt.Sprintf("%d0%d", timestamp, id)
	return
}
