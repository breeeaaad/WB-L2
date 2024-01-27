package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/
//Паттерн Фасад позволяет упростить работу с подсистемами и одновременно не ограничивает их.

//Плюсы: простая работа с подсистемами без знания внутренней структуры
//Минусы: Интерфейс фасада может быть нагромажден

//Примером могут служить любая группа подсистем, связь с которыми можно осуществить с помощью "Фасада"
//К примеру Администратор крупного приложения может регулировать процессы в микросервисах этого приложения

//Мой пример заказ продукта. Три структуры покупки, оформления заказа и доставки. Фасадом выступает либо консультант, либо интерфейс сайта.
type Product struct {
}

func (p *Product) buy() {
	fmt.Println("Покупка продукта")
}

type Order struct{}

func (o *Order) registration() {
	fmt.Println("Оформление заказа")
}

type Package struct{}

func (p *Package) send() {
	fmt.Println("Отправка посылки")
}

type Сonsultant struct {
	product *Product
	order   *Order
	pack    *Package
}

func New() *Сonsultant {
	return &Consultant{
		product: &Product{},
		order:   &Order{},
		pack:    &Package{},
	}
}

func (c Consultant) Request() {
	c.product.buy()
	c.order.registration()
	c.pack.send()
}
