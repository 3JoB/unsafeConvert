//go:build go1.20
// +build go1.20

package unsafeConvert

import (
	"bytes"
	"errors"
	"strings"
	"unsafe"
)

func ByteSlice(v string) []byte {
	return unsafe.Slice(unsafe.StringData(v), len(v))
}

func BytePointer(v string) []byte {
	return *(*[]byte)(unsafe.Pointer(&v))
}

func ByteBytes(v string) []byte {
	var i bytes.Buffer
	i.WriteString(v)
	defer i.Reset()
	return i.Bytes()
}

func ByteCopy(s string) []byte {
	buf := make([]byte, len(s))
	copy(buf, s)
	return buf
}

func StringSlice(v []byte) string {
	return unsafe.String(unsafe.SliceData(v), len(v))
}

func StringPointer(v []byte) string {
	return *(*string)(unsafe.Pointer(&v))
}

func StringStrings(v []byte) string {
	var i strings.Builder
	i.Write(v)
	defer i.Reset()
	return i.String()
}

// Note that this method is extremely dangerous,
// please do not use this method unless absolutely necessary!
//
// Before generics were used, this method could cause memory
// usage exceptions, leading to system unresponsiveness
// (no matter how much memory is installed).
// This problem has been fixed by adding generics,
// but there is no guarantee that memory exceptions will not
// continue to occur in the future.
//
// ----------------------------------------------------------------
//
// This method is used to convert any type to a string type.
// However, please note that not all types can be converted.
//
// Abusing this method can lead to program crashes or even
// system crashes. So before converting, please test in a virtual
// environment whether the type you want to convert can be converted,
// in order to avoid loss.
func AnyString[T ~string | ~[]byte](v T) string {
	return *(*string)(unsafe.Pointer(&v))
}

func Itoa(i int) string {
	var b [20]byte
	bp := len(b) - 1

	var sign bool
	if i < 0 {
		sign = true
		i = -i
	}

	for i >= 10 {
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}

	b[bp] = byte('0' + i)
	if sign {
		bp--
		b[bp] = '-'
	}

	return string(b[bp:])
}

func Atoi(s string) (int, error) {
	if len(s) == 0 {
		return 0, errors.New("empty string")
	}

	var neg bool
	if s[0] == '-' {
		neg = true
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	const maxUint = 1<<32 - 1

	var n uint64
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			continue
		}
		digit := uint64(ch - '0')
		if n > maxUint/10 || (n == maxUint/10 && digit > maxUint%10) {
			// overflow
			return 0, errors.New("overflow")
		}
		n = n*10 + digit
	}

	if neg {
		if n > maxUint+1 {
			return 0, errors.New("underflow")
		}
		return -int(n), nil
	} else {
		if n > maxUint {
			return 0, errors.New("overflow")
		}
		return int(n), nil
	}
}
