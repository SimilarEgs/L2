package main

import "fmt"

// Это поведенческий паттерн, цель которого внедрить дополнительное поведение для определенных объектов,
// не изменяя при этом конечные объекты.
//
// Цель паттерна:
// - Расширить существующий функционал объектов(структур), не изменяя при этом сам объект.
// Условно есть интерфейс и некие объекты, которые имлементируют его.
// Стоит задача добавить новое поведение для определенных объектов этого интерфейса,
// но мы не хотим внедрять новые методы в наш интерфейс изменяя устоявшиеся объекты.
//
// - (Другими словами)
//   Отделить логику, необходимую для работы с конкретным объектом от самого объекта
//
// Достоинвства:
// - Реалезуем открытый/закрытый принцип
//   изменяем/добовляем поведение наших классам не затрягивая основную структуру
//
// - Реализуем принцип единой ответственности
//   объявляя логику поведения объектов через отдельные структуры и интерфейсы
//
// Недостатки:
// - У посетителей может не быть доступа к приватным полям и методам элементов которые они поддерживают
// - Нужно обновлять всех посетитлий, всякий раз, когда мы добовляем/удаляем новые структры
//
// Структура паттерна:
// · Visiotr Interface - интерфейс визиотра, который привносит в наши классы новое поведение.
// · Visiotr Methods - методы, которые имплементируют интерфейс визитора, определяем их для структуры поведение которой хотим расширить
// · Visiotr Structs - структуры, которые имплементируют интерфейс визитора

type Shape interface {
	getType() string
	accept(v Visiotr) // добавив этот метод, мы изменили существующие структуры единожды, затем все дальнейшие изменения будут работать через этот метод
}

type Circle struct{}

type Square struct{}

type Triangle struct{}

func (c *Circle) getType() string {
	return "circle"
}
func (s *Square) getType() string {
	return "square"
}
func (t *Triangle) getType() string {
	return "triangle"
}

func (s *Square) accept(v Visiotr) {
	v.squareVisitor(s)
}
func (t *Triangle) accept(v Visiotr) {
	v.triangleVisitor(t)
}

func (c *Circle) accept(v Visiotr) {
	v.circleVisitor(c)
}

// Visiotr - расширяет поведение объектов интерфейса Shape
type Visiotr interface {
	squareVisitor(s *Square)
	circleVisitor(c *Circle)
	triangleVisitor(t *Triangle)
}

type AreaCalculator struct{}

type RadiusCalculator struct{}

func (a *AreaCalculator) squareVisitor(s *Square) {

	{
		// производим вычисление площади квадрата
	}

	fmt.Println("[Info] calculating area for:", s.getType())
}

func (a *AreaCalculator) circleVisitor(c *Circle) {

	{
		// производим вычисление площади круга
	}

	fmt.Println("[Info] calculating area for:", c.getType())
}

func (a *AreaCalculator) triangleVisitor(t *Triangle) {

	{
		// производим вычисление площади треугольника
	}

	fmt.Println("[Info] calculating area for:", t.getType())
}

func (r *RadiusCalculator) circleVisitor(c *Circle) {

	{
		// находим радиус круга
	}

}
func (r *RadiusCalculator) triangleVisitor(t *Triangle) {

	{
		// находим радиус треугольника
	}

}

func main() {

	squareArea := AreaCalculator{}
	triangleArea := AreaCalculator{}
	circleArea := AreaCalculator{}

	square := Square{}
	triangle := Triangle{}
	circle := Circle{}

	squareArea.squareVisitor(&square)
	triangleArea.triangleVisitor(&triangle)
	circleArea.circleVisitor(&circle)

}
