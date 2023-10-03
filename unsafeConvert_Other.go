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

func IntTo64(v int) int64 {
	i, _ := strconv.ParseInt(Itoa(v), 10, 64)
	return i
}

func It64(i int) int64 {
	var result int64
	if i == 0 {
		return 0
	}
	neg := i < 0
	if neg {
		i = -i
	}
	for i > 0 {
		result = result*10 + int64(i%10)
		i /= 10
	}
	if neg {
		result = -result
	}
	return result
}


func Int64ToString(v int64) string {
	return strconv.FormatInt(v, 10)
}

func StringToInt64(v string) int64 {
	i, _ := strconv.ParseInt(v, 10, 64)
	return i
}

func Uint32ToString(n uint32) string {
    var buf [10]byte
    i := len(buf)

    for n >= 10 {
        i--
        buf[i] = byte(n%10 + '0') 
        n /= 10
    }

    i--
    buf[i] = byte(n + '0')

    return string(buf[i:])
}

func StringToInt(v string) int {
	i, _ := strconv.Atoi(v)
	return i
}
