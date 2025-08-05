package libnk

import (
	"errors"
	"os"
	"unicode/utf8"
)

func WriteRune(r rune) (int, error) {
	c := utf8.RuneLen(r)
	if c < 0 {
		return -1, errors.New("The rune is not a valid value to enconde UTF-8")
	}
	p := make([]byte, c)
	utf8.EncodeRune(p, r)
	i, err := os.Stdout.Write(p)
	return i, err
}
