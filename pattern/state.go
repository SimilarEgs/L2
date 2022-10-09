package main

import "fmt"

// Это поведенческий паттерн, который позволяет изменять поведение у объектов в зависимости от их внутреннего состояния
//
// Суть:
// У нас есть интерфейс состояния (State) и имплементация каждого состояния которое хотим достичь, а так же контекст, который содержит промежуточную информацию наших состояний.
// 1. State Interface  (интерфейс State)
// 2. State Interface implementation (метод Handle)
// 3. State Context (структура MP3Player)
//
// Цель паттерна:
// - Иметь тип, который меняет свое поведение в зависимости от своего состояния
//
// Достоинства:
// - принцип единой отвественности. Организуем код, в отдельныее структуры для конкретных состояний состояний
// - отктрытый/закрытый принцип. Добовляем новые состояния, при этом не изменяя существующие
//
// Недостатки:
// - оверкил для программ с небольшим количеством состояний

// Интерфейс который отвечает за состояния плеера
type PlayerState interface {
	Handle(player *MP3Player)
}

// Контекст состояний плеера
type MP3Player struct {
	state PlayerState
}

func NewMP3Player(state PlayerState) *MP3Player {
	// создаем новый плеер и устанавливаем его начальное состояние
	player := new(MP3Player)
	player.state = state
	return player

}

func (p *MP3Player) SetState(state PlayerState) {
	p.state = state
}


func (pl *MP3Player) PressButton() {
	// переключаем состояние  
	pl.state.Handle(pl)
}

// Pause - конкретное состояние
type Pause struct{}

// Play - конкретное состояние
type Play struct{}

func (s *Pause) Handle(player *MP3Player) {
	fmt.Println("[Info] player stopped playing")
	player.SetState(new(Play)) // меняем состояние
}

func (p *Play) Handle(player *MP3Player) {
	// проигрывем музыку и меняем состояние на паузу
	fmt.Println("[Info] music is playing: David Bowie - Changes")
	player.SetState(new(Pause))
}

func main() {

	// инициализируем новый плеер и сетим его начальное состояние
	player := NewMP3Player(new(Play))

	// [Info] music is playing: David Bowie - Changes
	player.PressButton()

	// [Info] player stopped playing
	player.PressButton()
}
