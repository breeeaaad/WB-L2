package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
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

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}
	if _, err := strconv.Atoi(string(s[0])); err == nil {
		return "", errors.New("некорректная строка")
	}
	s1 := []rune(s)
	var b strings.Builder
	for i := 0; i < len(s1); i++ {
		if s1[i] == 92 {
			i++
		}
		if i+1 != len(s1) {
			if unicode.IsDigit(s1[i+1]) {
				n, _ := strconv.Atoi(string(s1[i+1]))
				for j := 0; j < n; j++ {
					b.WriteRune(s1[i])
				}
				i++
			} else {
				b.WriteRune(s1[i])
			}
		} else {
			b.WriteRune(s1[i])
		}
	}
	return b.String(), nil
}
