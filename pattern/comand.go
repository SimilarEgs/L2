package main

import (
	"log"
)

// Command:
//
// Это поведенческий паттерн основаный на инкапсуляции и абстракции.
// Каждая команда это отдельный объект.
//
// - Абстракция в этом патарне выражаеится в том смысле, что нам не нужно знать детали о работае конкретной команды,
//  суть этой команды укладываеться в действие которое она выполняет
// - Инкапусуляция позваляет не зависить командам друг от друга, т.к команда имеет всю необходимую информацию о запросе
//   и сможет выполнить его самостоятельно
//
// Устройсвто паттерна:
// · Caller/Invoker класс - инвокер ничего не знает об имплементации конкртетной команды, а лишь хронит в себе командный иетерфейс и вызывает метод вызова наших команд.
// · Command/Action Object - конкретная команда
// · Receiver - это класс, который содержит в себе бизнес-логику. Он хронит методы, которые вызываються с помощью метода вызова команды execute()
//
// Достоинвства:
// - Изолируем компоненты от прямой зависимости
// - Возможность собирать сложные команды из простых, используя один и тот же код, пример: однавременное закрытие и сохранение программы
// - Предстовление наших запросов (команд) в виде объектов
//
// Недостатки:
// - Загруженность кода из-за наличия множества методов/структур классов

// Receiver interface
type device interface {
	on()
	off()
	turnUp()
	turnDown()
}

// Commande interface
type command interface {
	execute()
}

// Concrete Command
type onCommand struct {
	device device
}

// Concrete Command method execution
func (on *onCommand) execute() {
	on.device.on()
}

// Concrete receiver that's implements device interface
type TV struct {
	isOn   bool
	volume int
}

// Implementing actual logic
func (tv *TV) on() {
	tv.isOn = true
	log.Println("[Info] tv is on")
}

func (tv *TV) off() {
	tv.isOn = false
	log.Println("[Info] tv is off")
}

func (tv *TV) turnDown() {

	if tv.isOn {

		if tv.volume >= 0 && tv.volume < 100 {
			tv.volume++
			log.Println("[Info] volume increased: ", tv.volume)
			return
		} else {
			log.Println("[Info] max volume")
		}

	} else {
		log.Println("[Error] tv is off")
	}

}
func (tv *TV) turnUp() {

	if tv.isOn {

		if tv.volume >= 0 && tv.volume <= 100 {
			tv.volume--
			log.Println("[Info] volume decreased: ", tv.volume)
			return
		} else {
			log.Println("[Info] min volume")
		}

	} else {
		log.Println("[Error] tv is off")
	}

}

// Invoker struct
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

type offCommand struct {
	device device
}

func (off *offCommand) execute() {
	off.device.off()
}

type turnUp struct {
	device device
}

func (up *turnUp) execute() {
	up.device.turnUp()

}

type turnDown struct {
	device device
}

func (down *turnDown) execute() {
	down.device.turnDown()
}

func main() {

	tv := &TV{
		isOn:   false,
		volume: 50,
	}

	// Comands instance
	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	turnUp := &turnUp{
		device: tv,
	}

	turnDown := &turnDown{
		device: tv,
	}

	// Invoker instance
	buttonOn := &button{
		command: onCommand,
	}
	buttonOff := &button{
		command: offCommand,
	}
	buttonUp := &button{
		command: turnUp,
	}
	buttonDown := &button{
		command: turnDown,
	}

	// Execution
	buttonOn.press()
	buttonDown.press()
	buttonUp.press()
	buttonOff.press()

}
