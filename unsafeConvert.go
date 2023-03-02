//go:build go1.20
// +build go1.20
package unsafeConvert

import "unsafe"

func Bytes(v string) []byte {
	return unsafe.Slice(unsafe.StringData(v), len(v))
}

func String(v []byte) string {
	return unsafe.String(&v[0], len(v))
}