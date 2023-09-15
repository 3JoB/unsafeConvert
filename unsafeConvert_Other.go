//go:build go1.20
// +build go1.20

package unsafeConvert

import (
	"strconv"
)

func Float32(f float32) string {
	return strconv.FormatFloat(float64(f), 'g', -1, 32)
}

func Float64(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}

func IntToString(v int) string {
	return strconv.Itoa(v)
}

func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func StringToInt64(v string) int64 {
	i, _ := strconv.ParseInt(v, 10, 64)
	return i
}

func StringToInt(v string) int {
	i, _ := strconv.Atoi(v)
	return i
}
