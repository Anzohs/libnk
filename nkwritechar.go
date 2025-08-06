package libnk

import (
	"errors"
	"os"
	"unicode/utf8"
)

type Char rune

func NkWriteChar(r Char) (int, error) {
	c := utf8.RuneLen(rune(r))
	if c < 0 {
		return -1, errors.New("The rune is not a valid value to enconde UTF-8")
	}
	p := make([]byte, c)
	utf8.EncodeRune(p, rune(r))
	i, err := os.Stdout.Write(p)
	return i, err
}
