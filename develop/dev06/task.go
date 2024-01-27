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

type fields []string

func (fields) String() string {
	return ""
}

func (fields) Set(value string) error {
	return nil
}

func cut(lines []string) ([]string, error) {
	var (
		res []string
		b   strings.Builder
		f   fields
	)
	delimiter := flag.String("d", "\t", "использовать другой разделитель")
	separated := flag.Bool("s", true, "только строки с разделителем")
	f = []string{"1", "3"}
	flag.Var(&f, "f", "выбрать поля (колонки)")
	for _, v := range lines {
		tmp := strings.Split(v, *delimiter)
		if (*separated && len(tmp) > 1) || (!*separated) {
			for _, column := range f {
				val, err := strconv.Atoi(column)
				if err != nil {
					return nil, err
				}
				if val <= len(tmp) {
					b.WriteString(tmp[val-1])
					b.WriteString(*delimiter)
				}
			}
			b.WriteString("\n")
			res = append(res, b.String())
			b.Reset()
		}
	}

	return res, nil
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}
	lines := make([]string, n)
	for i := 0; i < n; i++ {
		s, err := bufio.NewReader(os.Stdin).ReadString('\r')
		if err != nil {
			log.Fatal(err)
		}
		lines[i] = s
	}

	c, err := cut(lines)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	for _, v := range c {
		fmt.Print(v)
	}
}
