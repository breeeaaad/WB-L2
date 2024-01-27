package main

import (
	"fmt"
	"hash/fnv"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sorts(s string) uint32 {
	b := []byte(s)
	sort.Slice(
		b,
		func(i, j int) bool {
			return (b)[i] < (b)[j]
		})
	hash := fnv.New32a()
	hash.Write(b)
	return hash.Sum32()
}

func find(s []string) map[string][]string {
	tmp := make(map[uint32][]string, len(s))
	var hash uint32
	for _, v := range s {
		hash = sorts(v)
		if _, found := tmp[hash]; !found {
			tmp[hash] = make([]string, 0)
			tmp[hash] = append(tmp[hash], v)
			continue
		}
		tmp[hash] = append(tmp[hash], v)
	}
	res := make(map[string][]string, len(tmp))
	var t []string
	for _, v := range tmp {
		if len(v) <= 1 {
			continue
		}
		t = v[1:]
		for i, j := range t {
			t[i] = strings.ToLower(j)
		}
		sort.Strings(t)
		res[v[0]] = t
	}
	return res
}

func main() {
	words := []string{"пятак", "листок", "слиток", "пятка", "тяпка", "столик"}
	fmt.Print(find(words))
}
