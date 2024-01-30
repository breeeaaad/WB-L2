package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/
//Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
//Применимость
//Когда вам нужно использовать разные вариации какого-то алгоритма внутри одного объекта.
//Когда у вас есть множество похожих классов, отличающихся только некоторым поведением.
//Когда вы не хотите обнажать детали реализации алгоритмов для других классов.
//Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора. Каждая ветка такого оператора представляет собой вариацию алгоритма.
//+:
/* Горячая замена алгоритмов на лету.
Изолирует код и данные алгоритмов от остальных классов.
Уход от наследования к делегированию.
Реализует принцип открытости/закрытости.*/
//-:
/*Усложняет программу за счёт дополнительных классов.
Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.*/
type Strategy interface {
	execute(int, int) int
}

type ConcreteStrategyAdd struct {
}

func (ConcreteStrategyAdd) execute(a, b int) int {
	return a + b
}

type ConcreteStrategySubstract struct {
}

func (ConcreteStrategySubstract) execute(a, b int) int {
	return a - b
}

type ConcreteStrategyMultiply struct {
}

func (ConcreteStrategyMultiply) execute(a, b int) int {
	return a * b
}

type Context struct {
	strategy Strategy
}

func (c *Context) executeStrategy(a, b int) int {
	return c.strategy.execute(a, b)
}

func main() {
	c := ConcreteStrategyAdd{}
	con := Context{
		strategy: c,
	}
	con.executeStrategy(4, 6)
}
