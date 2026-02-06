package utils

import (
	"os"
	"strconv"
)

func Getenv(key string, defaultVal ...string) string {
	get := os.Getenv(key)
	if get == "" && defaultVal != nil {
		return defaultVal[0]
	}
	return get
}

// ParseInt64 将字符串解析为 int64，如果解析失败则返回 0
func ParseInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
