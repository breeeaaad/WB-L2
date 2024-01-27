Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
от 0 до 9 и deadlock. Range ждет еще значения из канала. Чтоб выйти из цикла, нужно закрыть канал.

```
