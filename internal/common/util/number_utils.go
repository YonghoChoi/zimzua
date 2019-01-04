package util

import (
	"bytes"
	"strconv"
)

func ScanNumeric(str string) (int, error) {
	buf := new(bytes.Buffer)

	for i := 0; i < len(str); i++ {
		c := str[i]
		if c < '0' || c > '9' {
			break
		}
		buf.WriteByte(c)
	}
	return strconv.Atoi(buf.String())
}

func MaxUInt() uint {
	return ^uint(0)
}

func MinUInt() uint {
	return 0
}

func MaxInt() int {
	const maxUint = ^uint(0)
	return int(maxUint >> 1)
}

func MinInt() int {
	const maxUint = ^uint(0)
	const maxInt = int(maxUint >> 1)
	return -maxInt - 1
}
