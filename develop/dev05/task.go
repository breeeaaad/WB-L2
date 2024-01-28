package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
type flags struct {
	A, B, C       int
	c, i, v, F, n bool
	file, sample  string
}

func grep(data []string, f *flags) (map[int]string, []int) {
	if f.C > 0 {
		f.A = f.C
		f.B = f.C
	}
	res := make(map[int]string)
	key := make([]int, 0)
	b := false
	for i, v := range data {
		w := v
		if f.i {
			w = strings.ToLower(w)
			f.sample = strings.ToLower(f.sample)
		}
		if f.F {
			if strings.Contains(w, f.sample) {
				res[i+1] = v
				key = append(key, i+1)
				b = true
			} else {
				check, err := regexp.MatchString(f.sample, w)
				if err != nil {
					log.Fatal(err)
				}
				if f.v {
					if !check {
						res[i+1] = v
						key = append(key, i+1)
						b = true
					}
				} else {
					if check {
						res[i+1] = v
						key = append(key, i+1)
						b = true
					}
				}
				if (f.B > 0 || f.A > 0) && b && !f.c {
					j := i + 1
					for n := f.A; n > 0; n-- {
						if j < len(data) {
							res[j+1] = data[j]
							key = append(key, j+1)
							j++
						}
					}
					j = i - 1
					for n := f.B; n > 0; n-- {
						res[j+1] = data[j]
						key = append(key, j+1)
						j--
					}
					b = false
				}
			}
		}
	}
	sort.Ints(key)
	if f.c {
		fmt.Print(len(res))
		return nil, nil
	}
	return res, key
}

func parse() (*flags, error) {
	var f flags
	flag.IntVar(&f.A, "A", -1, "печатает N строк после совпадения")
	flag.IntVar(&f.B, "B", -1, "печатает N строк до совпадения")
	flag.IntVar(&f.C, "C", -1, "печатает N строк вокруг совпадения")
	flag.BoolVar(&f.c, "c", false, "количество строк")
	flag.BoolVar(&f.i, "i", false, "игнорирование регистра")
	flag.BoolVar(&f.v, "v", false, "исключение")
	flag.BoolVar(&f.F, "F", false, "точное совпадение")
	flag.BoolVar(&f.n, "n", false, "номер строки")
	flag.Parse()
	if len(flag.Args()) != 2 {
		return nil, errors.New("Недостаточно аргументов")
	}
	f.sample = flag.Arg(0)
	f.file = flag.Arg(1)
	return &f, nil
}

func main() {
	f, err := parse()
	if err != nil {
		log.Fatal(err.Error())
	}
	file, err := os.Open(f.file)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(file)
	buf, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	res := strings.Split(string(buf), "\n")
	r, k := grep(res, f)
	for _, v := range k {
		if f.n {
			fmt.Println(v, ":", r[v])
		} else {
			fmt.Println(r[v])
		}
	}
}
