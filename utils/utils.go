package utils

import (
	"bytes"
)
func MultiJoinString(str ...string) string {
	var b bytes.Buffer
	length := len(str)
	for i := 0; i < length; i++ {
		b.WriteString(str[i])
	}
	return b.String()
}