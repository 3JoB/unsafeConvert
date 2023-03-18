//go:build go1.20
// +build go1.20

package unsafeConvert

import (
	"strconv"
	"unsafe"

	//"github.com/goccy/go-reflect"
)

func Bytes(v string) []byte {
	return unsafe.Slice(unsafe.StringData(v), len(v))
}

func BytesReflect(v string) []byte {
	return *(*[]byte)(unsafe.Pointer(&v))
}

func String(v []byte) string {
	return unsafe.String(unsafe.SliceData(v), len(v))
}

func StringReflect(v []byte) string {
	return *(*string)(unsafe.Pointer(&v))
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
