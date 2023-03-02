//go:build go1.20
// +build go1.20

package unsafeConvert

import (
	"strconv"
	"unsafe"
)

func Bytes(v string) []byte {
	return unsafe.Slice(unsafe.StringData(v), len(v))
}

func String(v []byte) string {
	return unsafe.String(&v[0], len(v))
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