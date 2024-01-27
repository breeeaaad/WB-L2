package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	fields       string
	delimiter    string
	is_separated bool
}

func parseFlags() *flags {
	f := flags{}

	flag.StringVar(&f.fields, "f", "0", "fields")
	flag.StringVar(&f.delimiter, "d", "\t", "delimiter")
	flag.BoolVar(&f.is_separated, "s", false, "separated")
	flag.Parse()

	return &f
}

func cut(input string, f *flags) string {
	if f.is_separated && !strings.Contains(input, f.delimiter) {
		return ""
	}

	sb := strings.Builder{}
	sp := strings.Split(input, f.delimiter)
	columns := strings.Split(f.fields, ",")
	for i := 0; i < len(columns); i++ {
		column, err := strconv.Atoi(columns[i])
		if err != nil {
			log.Fatalln("Некорректный ввод номера столбцов: ", err.Error())
		}

		sb.WriteString(sp[column])
	}
	return sb.String()
}

func main() {
	f := parseFlags()
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		text := sc.Text()
		fmt.Println(">> ", cut(text, f))
	}
}
