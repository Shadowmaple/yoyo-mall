package util

import (
	"time"
)

// 获取当前时间，东八区
func GetCurrentTime() *time.Time {
	// loc, _ := time.LoadLocation("Asia/Shanghai")
	// loc := time.FixedZone("CST", 8*3600)
	t := time.Now().UTC().Add(8 * time.Hour)
	return &t
}

func ParseTime(t *time.Time) (string, string) {
	return t.Format("2006-01-02"), t.Format("15:04:05")
}

func GetStandardTime(t *time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
