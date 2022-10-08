package main

import "fmt"

// Это порождающий паттерн, который описывает общую реализацию создания объектов с помощью фабричных методов
//
// суть:
// - описание создания конечного объекта через фабрику
// дабы абстрагироваться от деталей создания объекта, без указания точной сущности объекта который будет создан.
//
// Достоинвства:
// - решение проблемы повторяющегося кода
// - инкапусуляция
// - откртый/закртый принцип
// - принцип единой отвественности
//
// Недостатки:
// - привязка к созданию объектов через единый конструктор + единый божественный интерфейс
//
// “A factory defined by a function that helps us to create instances of a certain structure
// with determinate values or values that could be provided in the function arguments.”

func main() {
	iphone13pro := NewPhone("Iphone", "pro")
	iphone13pro.getType()
	iphone13pro.printDetails()
}

type Phone interface {
	getType() string
	printDetails()
}

// Factory - общая реализация создания телефонов
func NewPhone(t, model string) Phone {
	switch t {
	case "Iphone":
		return NewIphone(model)
	case "Android":
		return NewAndroid(model)
	default:
		fmt.Println("Uknown type")
		return nil
	}
}

type Iphone struct {
	Type   string
	OS     string
	Camera string
	faceID bool
	Core   string
}

func (i *Iphone) getType() string {
	return i.Type
}

func (i *Iphone) printDetails() {
	fmt.Printf("%+v\n", i)
}

func NewIphone(model string) Phone {
	switch model {
	case "pro":
		return &Iphone{
			Type:   "Iphone 13 pro",
			OS:     "IOS",
			Camera: "12mpx",
			faceID: true,
			Core:   "5nm",
		}
	case "standart":
		return &Iphone{
			Type:   "Iphone SE",
			OS:     "IOS",
			Camera: "12mpx",
			faceID: false,
			Core:   "7nm",
		}
	default:
		return nil
	}
}

type Android struct {
	Type   string
	Camera string
	OS     string
	Core   string
}

func NewAndroid(model string) Phone {
	switch model {
	case "ultra":
		return &Android{
			Type:   "Sasung ultra S20",
			OS:     "Android",
			Camera: "108mpx",
			Core:   "7nm",
		}
	case "standart":
		return &Android{
			Type:   "Sasung s20",
			OS:     "Android",
			Camera: "64mpx",
			Core:   "10nm",
		}
	default:
		return nil
	}
}

func (a *Android) getType() string {
	return a.Type
}

func (a *Android) printDetails() {
	fmt.Printf("%+v\n", a)
}
