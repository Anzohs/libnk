package libnk

func NkStrlen(s string) int {
	c := []Char(s)
	return (len(c))
}
