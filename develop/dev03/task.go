package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

# Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительное

# Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
type flags struct {
	file    string
	col     int
	num     bool
	reverse bool
	unique  bool
}

func parse() *flags {
	var in flags
	flag.IntVar(&in.col, "k", -1, "колонки для сортировки")
	flag.BoolVar(&in.num, "n", false, "сортировка по числовому значению")
	flag.BoolVar(&in.reverse, "r", false, "сортировка в обратном порядке")
	flag.BoolVar(&in.unique, "u", false, "уникальные значения")
	flag.Parse()
	in.file = flag.Arg(0)
	return &in
}

func sorting(l []string, in *flags) []string {
	if in.col >= 0 {
		colsort(l, in.col)
	} else {
		sort.Strings(l)
	}
	if in.reverse {
		reverse(l)
	}
	if in.unique {
		l = unique(l)
	}
	return l
}

func colsort(l []string, col int) []string {
	m := make(map[string]string)
	for i := 0; i < len(l); i++ {
		sp := strings.Split(l[i], " ")
		m[sp[col]] = l[i]
	}
	k := make([]string, 0, len(m))
	for k1 := range m {
		k = append(k, k1)
	}
	sort.Strings(k)
	sorted := make([]string, len(l))
	for i, k1 := range k {
		sorted[i] = m[k1]
	}
	return sorted
}

func reverse(d []string) {
	for i, j := 0, len(d)-1; i < j; i, j = i+1, j-1 {
		d[i], d[j] = d[j], d[i]
	}
}

func unique(s []string) []string {
	m := make(map[string]struct{})
	var res []string
	for _, s1 := range s {
		if _, ok := m[s1]; !ok {
			m[s1] = struct{}{}
			res = append(res, s1)
		}
	}
	return res
}

func main() {
	flags := parse()
	var data []string
	in, err := os.Open(flags.file)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(in)
	if err != nil {
		log.Fatal(err.Error())
	}
	sc := bufio.NewScanner(in)
	for sc.Scan() {
		data = append(data, sc.Text())
	}
	sorted := sorting(data, flags)
	out, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(out)
	for i := 0; i < len(sorted); i++ {
		_, err := fmt.Fprintf(out, "%s\n", data[i])
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
