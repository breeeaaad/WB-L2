package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/
//Состояние — это поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего состояния.
//Применимость
//Когда у вас есть объект, поведение которого кардинально меняется в зависимости от внутреннего состояния, причём типов состояний много, и их код часто меняется.
//Когда код класса содержит множество больших, похожих друг на друга, условных операторов, которые выбирают поведения в зависимости от текущих значений полей класса.
//Когда вы сознательно используете табличную машину состояний, построенную на условных операторах, но вынуждены мириться с дублированием кода для похожих состояний и переходов.
//+:
/* Избавляет от множества больших условных операторов машины состояний.
Концентрирует в одном месте код, связанный с определённым состоянием.
Упрощает код контекста.*/
//-:
/*Может неоправданно усложнить код, если состояний мало и они редко меняются.*/
type Player struct {
	currentState State
	idleState    State
	walkingState State
	jumpingState State
	x            int
	y            int
}

func (p *Player) setState(state State) {
	p.currentState = state
}

func (p *Player) Up() {
	p.currentState.Up()
}

func (p *Player) Down() {
	p.currentState.Down()
}

func (p *Player) Left() {
	p.currentState.Left()
}

func (p *Player) Right() {
	p.currentState.Right()
}

func (p *Player) jump() {
	p.currentState.jump()
}

type State interface {
	Up()
	Down()
	Left()
	Right()
	jump()
}

type IdleState struct {
	player *Player
}

func (s *IdleState) Up() {}

func (s *IdleState) Down() {}

func (s *IdleState) Left() {}

func (s *IdleState) Right() {}

func (s *IdleState) jump() {
	fmt.Println("Jumped!")
	s.player.setState(s.player.jumpingState)
}

type WalkingState struct {
	player *Player
}

func (s *WalkingState) Up() {
	s.player.y -= 1
}

func (s *WalkingState) Down() {
	s.player.y += 1
}

func (s *WalkingState) Left() {
	s.player.x -= 1
}

func (s *WalkingState) Right() {
	s.player.x += 1
}

func (s *WalkingState) jump() {
	fmt.Println("Jumped!")
	s.player.setState(s.player.jumpingState)
}

type JumpingState struct {
	player *Player
}

func (s *JumpingState) Up() {
	s.player.y -= 2
}

func (s *JumpingState) Down() {
	s.player.y += 2
}

func (s *JumpingState) Left() {
	s.player.x -= 2
}

func (s *JumpingState) Right() {
	s.player.x += 2
}

func (s *JumpingState) jump() {}

func main() {
	idleState := &IdleState{}
	walkingState := &WalkingState{}
	jumpingState := &JumpingState{}

	player := &Player{
		currentState: idleState,
		idleState:    idleState,
		walkingState: walkingState,
		jumpingState: jumpingState,
		x:            0,
		y:            0,
	}

	player.Right()
	player.Up()
	player.jump()
	player.Right()
}
