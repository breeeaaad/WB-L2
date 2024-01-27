Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Случайный порядок значений переданные в функции asChan, после этих значений, из-за бесконечного цикла в функции merge, будут выводится нули до остановки программы извне.
Дело в том, что 2 функции asChan породят 2 горутины пишущие в канал c задержкой в виде time.Sleep при записи, а функция merge породит еще одну горутину,которая по очереди будет вытаскивать данные из двух каналов(кто первее доставит). После закрытия каналов будут передаваться нули, ведь нет условий проверок на закрытие канала.

```
