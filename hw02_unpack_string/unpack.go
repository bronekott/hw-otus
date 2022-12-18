package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result strings.Builder
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			return "", ErrInvalidString
		}

		if i+1 < len(runes) && runes[i] == '\\' {
			i++
			if !unicode.IsDigit(runes[i]) && runes[i] != '\\' {
				return "", ErrInvalidString
			}
		}

		if i+1 < len(runes) && unicode.IsDigit(runes[i+1]) {
			repeatCount, _ := strconv.Atoi(string(runes[i+1]))
			result.WriteString(strings.Repeat(string(runes[i]), repeatCount))
			i++
		} else {
			result.WriteString(string(runes[i]))
		}
	}
	return result.String(), nil
}
