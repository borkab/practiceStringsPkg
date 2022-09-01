package my

import (
	"strings"
)

func MyTrim(s string) string {
	ss := strings.TrimSpace(s)
	return ss
}

func MyContains(str, str2 string) bool {
	bo := strings.Contains(str, str2)
	return bo
}

func MyHasPrefix(str1, str2 string) bool {
	b := strings.HasPrefix(str1, str2)
	return b
}

func MyHasSuffix(str1, str2 string) bool {
	bl := strings.HasSuffix(str1, str2)
	return bl
}

func MyNewReader(str1 string) *strings.Reader {
	new := strings.NewReader(str1)
	return new
}
