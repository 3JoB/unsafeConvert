//go:build go1.20
// +build go1.20

package unsafeConvert

import (
	"strconv"
	"unsafe"

	"github.com/goccy/go-reflect"
)

func Bytes(v string) []byte {
	return unsafe.Slice(unsafe.StringData(v), len(v))
}

func BytesSliceReflect(v string) []byte {
	return *(*[]byte)(unsafe.Pointer(&v))
}

func String(v []byte) string {
	return unsafe.String(unsafe.SliceData(v), len(v))
}

func StringReflect(v []byte) string {
	return *(*string)(unsafe.Pointer(&v))
}

// String converts a slice of bytes into a string without performing a copy.
// NOTE: This is an unsafe operation and may lead to problems if the bytes
// passed as argument are changed while the string is used.  No checking whether
// bytes are valid UTF-8 data is performed.
func StringBasic(bytes []byte) string {
	hdr := *(*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	return *(*string)(unsafe.Pointer(&reflect.StringHeader{
		Data: hdr.Data,
		Len:  hdr.Len,
	}))
}

// Bytes converts a string into a slice of bytes without performing a copy.
// NOTE: This is an unsafe operation and may lead to problems if the bytes are
// changed.
func BytesBasic(str string) []byte {
	hdr := *(*reflect.StringHeader)(unsafe.Pointer(&str))
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Data: hdr.Data,
		Len:  hdr.Len,
		Cap:  hdr.Len,
	}))
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
