package main

import "fmt"

// Определение: 1
// Это поведенческий паттерн, который определяет набор алгоритмов, для выполнения какого-то одного нацеленного действия определенный группы.
// Алгоритм выбирается в зависимости от требования клиента.
//
// Определение: 2
// Это поведенческий паттерн, который определяет семью алгоритмов, инкапсулирует каждый из них,
// и делает их взаимозаменяемыми.
// А также позволяет алгоритмам меняться независимо от клиентов которые его используют.
//
// Суть:
// - Смысл действия один, а реализации(алгоритмы) разные.
// - Все сводится к тому, чтобы изменять функционал действия для необходимой операции.
//
// Отличия от паттерна State:
// - Нет контекста, где хроним текущее состояние
// - Действие, как правило однаразовое (сортитовка/авторизация/платежка)
//
// Цель паттерна:
// - используем, если у нас есть объект поведение которого меняться в рантайме
//
// Достоинвства:
// - Гибкость в использовании.
//  Динамически изменяем поведение объекта в рантайме в зависимости от выбранного алгоритма.
//  Это полезно, если есть объект, который должен выполнить какое-то действие, но разными способами.
// - Разделяет алгортим от хоста, инкапсулировав его в собсвтенный тип
//
//
// Недостатки:
// -
//
// Пример реалезации:
// - Калькулятор.
//  Условно мы имеем калькулятор, который принимает два числа на вход, а также у которого есть кнопки, которые определяют
//  вычисления необходимые для выполнения с этими числами.
//  Добавляя функционал новых способов вычислений, мы не должны изменить логику работы калькулятора (какое-то действие над числами и вывод результата)
//  Решением будет - инкапсулировавать логику вычеслений в независимые конструкции, которым будем передвать единый интерфейс стратегии вычисления.
//  Конкретая стратегии - сложение, вычитание, деление и т.д
//  Контекст - калькулятор
//
                                                           
				      «interface»	     | strategy 1  |
  |  Context   |     strategy       |  Strategy   |          | algorithm() |
  |------------|  ----------------> |-----------  |  <-------|/////////////|
  | operation()|                    | algorithm() |	     |
                                                             | strategy 2  | 
							     | algorithm() |

type Authorization interface {
	Auth() error
}

type MobileAuth struct {
	phone string
}

type DBAuth struct {
	id        string
	loginHash string
}

type FileAuth struct {
	id    string
	login string
}

// конструкторы для наших стратегий
func newMobileAuth() Authorization {
	return &MobileAuth{}
}
func newDBAuth() Authorization {
	return &DBAuth{}
}
func newFileAuth() Authorization {
	return &FileAuth{}
}

func (a *MobileAuth) Auth() error {
	fmt.Println("[Info] authorization by phone number")
	{
		// здесь должна быть имплементация авторизации через мобильный телефон
	}
	return nil
}

func (a *DBAuth) Auth() error {
	fmt.Println("[Info] authorization via the DB...")
	{
		// здесь должна быть имплементация авторизации через базу данных
	}
	return nil
}

func (a *FileAuth) Auth() error {
	fmt.Println("[Info] authorization via the file system...")
	{
		// здесь должна быть имплементация авторизации через файловую систему
	}

	return nil
}

// бизнес-логика авторизации
func ProcessAuthorization(user string, do Authorization) {
	if err := do.Auth(); err != nil {
		return
	}
}

// В этом примере паттерн Стратегия изолирует логику авторизации юзера через третий сервис от конечной функции авторизации.
// То есть логика авторизации через некий сервис это лишь часть этапа обработки нашей авторизации, и наш паттерн в этом случае служит разделителем
func main() {

	user := "Ivanov"

	for i := 1; i <= 3; i++ {

		var auth Authorization

		switch i {
		case 1:
			auth = newDBAuth()
		case 2:
			auth = newFileAuth()
		case 3:
			auth = newMobileAuth()
		default:
		}

		ProcessAuthorization(user, auth)

	}

}
