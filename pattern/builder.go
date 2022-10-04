package main

import "fmt"

// Builder:
//
// Реалезация паттерна заключатся в построении сложных объектов по частям.
// Что позволяет привнести в кодовую базу гибкость и легкую масштабируемость, вместе с удобочитаемым кодом.
//
// Когда использовать?
// Если стоит задача создать несколько экспемпляров одного и того же класса с разными характеристиками
//
// Достатки:
// - позваляет использовать один и тот же код для создания разных объектов 
// - позволяет создавать сложные объекты по шагам
//
// Недостатки:
// - Много кода (изобиле классов/структур/интерфейсы)
// - Для каждой сущности/продукта должен существовать свой конструктор
// - Привязка к объектом биледров, в случае если  не реалезуем метод интерфейса, его нужно будет добавить

// ПРИМЕР РЕАЛЕЗАЦИИ:
//
// Шаг 1 - объявление строителей и иницализация конструкторов
//
// структруа сложного объекта
type User struct {
	Name, Email, Address, Phone string

	Position, Salary string
}

// структура строителя UserBuilder
type UserBuilder struct {
	user *User
}

// структура строителя UserJobBuilder
type UserJobBuilder struct {
	UserBuilder
}

// конструктор строителя UserBuilder
func NewUserBuilder(name, email string) *UserBuilder {
	return &UserBuilder{
		user: &User{Name: name, Email: email},
	}
}

// констуктор строителя UserJobBuilder
func (ub *UserBuilder) Works() *UserJobBuilder {
	return &UserJobBuilder{*ub}
}

// Шаг 2 - Создаем методы для сборки объекта User по частям

// методы строителя UserBuilder:

// сетим адрес
func (ub *UserBuilder) StaysAt(address string) *UserBuilder {
	ub.user.Address = address
	return ub
}

// сетим номер
func (ub *UserBuilder) CallsAt(phone string) *UserBuilder {
	ub.user.Phone = phone
	return ub
}

// методы строителя UserJobBuilder

// сетим должность
func (jb *UserJobBuilder) As(position string) *UserJobBuilder {
	jb.user.Position = position
	return jb
}

// сетим зарплату
func (jb *UserJobBuilder) WithSalary(salary string) *UserJobBuilder {
	jb.user.Salary = salary
	return jb
}

// возвращаем объект User
func (ub *UserBuilder) Build() *User {
	return ub.user
}

// Шаг 3 - строим нашего юзера

func main() {

	userBuilder := NewUserBuilder("Alexey", "fishgame@gmail.com")

	userBuilder.StaysAt("Moscow").CallsAt("8(800)555-35-35").Works().As("trainee").WithSalary("0$")

	// создаем имутабельный объект типа юзуер
	user1 := userBuilder.Build()

	fmt.Printf("%+v\n", user1)

}
