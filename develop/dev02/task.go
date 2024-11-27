package main

import (
	"errors"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unpack(unpacked string) (string, error) {
	str := []rune(unpacked)
	var stringBuilder strings.Builder
	var lastRune rune
	for i := 0; i < len(str); i++ {
		switch func(j int) int {
			if str[j] >= '0' && str[j] <= '9' {
				return 1
			}
			if str[j] == '\\' {
				return 2
			}
			return 0
		}(i) {
		case 0:
			lastRune = str[i]
			stringBuilder.WriteRune(lastRune)

		case 1:
			if lastRune == 0 {
				return "", errors.New("An incorrect string was passed to the function - unable to unpack")
			}
			for j, _ := strconv.Atoi(string(str[i])); j > 1; j-- {
				stringBuilder.WriteRune(lastRune)
			}
			lastRune = 0
		case 2:
			i++
			lastRune = str[i]
			stringBuilder.WriteRune(str[i])
		}

	}
	return stringBuilder.String(), nil
}
