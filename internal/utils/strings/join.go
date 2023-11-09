package strings

import "strings"

func StringJoin(strs ...string) string {
	n := 0
	for i := 0; i < len(strs); i++ {
		n += len(strs[i])
	}

	var b strings.Builder
	b.Grow(n)
	for _, s := range strs {
		b.WriteString(s)
	}
	return b.String()
}
