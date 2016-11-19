package stack

import "bytes"

func CompareStrings(this, other string) bool {
	if this == "" || other == "" {
		return false
	}
	return bytes.Equal([]byte(this), []byte(other))
}
