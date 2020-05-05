package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input_str string) (str string, err error) {
	var sb strings.Builder

	if input_str == "" {
		str = ""
	} else {
		var prev, curr rune
		for _, s := range input_str {
			prev = curr
			curr = s

			switch {
			case prev == 0 && unicode.IsDigit(curr):
				return "", ErrInvalidString
			case unicode.IsDigit(prev) && unicode.IsDigit(curr):
				return "", ErrInvalidString
			}

			if !unicode.IsDigit(curr) {
				sb.WriteString(string(curr))
			} else {
				i, _ := strconv.Atoi(string(curr))
				sb.WriteString(strings.Repeat(string(prev), i-1))
			}
		}
		str = sb.String()
	}
	return
}

func main() {
	s, e := Unpack("aaa10b")
	fmt.Printf("%s %s", s, e)
}
