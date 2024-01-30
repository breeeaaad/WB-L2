package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/
//Порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово
//Плюсы
/*Позволяет создавать продукты пошагово.
Позволяет использовать один и тот же код для создания различных продуктов.
Изолирует сложный код сборки продукта от его основной бизнес-логики.*/
//Минусы
/*Усложняет код программы из-за введения дополнительных классов.
Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.*/
//Область применения
//Когда вы хотите избавиться от «телескопического конструктора».
//Когда ваш код должен создавать разные представления какого-то объекта.
type Car struct{}

type Manual struct{}

type Builder interface {
	reset()
	setSeates()
	setEngine()
	setGPS()
}

type CarBuilder struct{}

func (CarBuilder) reset() {
}
func (CarBuilder) setSeates() {
}
func (CarBuilder) setEngine() {
}
func (CarBuilder) setGPS() {
}

type CarManualBuilder struct{}

func (CarManualBuilder) reset() {
}
func (CarManualBuilder) setSeates() {
}
func (CarManualBuilder) setEngine() {
}
func (CarManualBuilder) setGPS() {
}

type Director struct {
	builder Builder
}

func (d Director) buildcar() {
	d.builder.reset()
	d.builder.setSeates()
	d.builder.setEngine()
	d.builder.setGPS()
}
