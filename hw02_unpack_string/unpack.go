package hw02unpackstring

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {

	var sb strings.Builder
	regexpString := `^[0-9]|[0-9][0-9]+`
	matched, _ := regexp.MatchString(regexpString, inputString)
	stringSlice := strings.Split(inputString, "")

	// Проверка с помощью регулярки на корректность строки, если строка некорректна, то возвращает ошибку
	if matched {
		return "", ErrInvalidString
	}
	
	for i, v := range stringSlice {
		tmpRune := []rune(v)

		if unicode.IsDigit(tmpRune[0]) {
			intSubString, _ := strconv.Atoi(v)

			if intSubString == 0 {
				tmpString := sb.String()
				sb.Reset()
				sb.WriteString(tmpString[:len(tmpString)-1])

			} else {
				newSubString := strings.Repeat((string(stringSlice[i-1])), intSubString-1)
				sb.WriteString(newSubString)
			}
		} else {
			sb.WriteString(v)
		}
	}
	fmt.Println("Распакованная строка: ", sb.String())
	return sb.String(), nil
}