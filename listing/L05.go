// Что выведет программа? Объяснить вывод программы.
//
// Ответ: error
//
// Почему:
// Потому как на мы присвоили переменной err результат выполнения функции test(), которая возвращает нам
// указатель на кастомную структуру, которая в свою очередь будет nil.
// Если мы запретим тип err, после присвоения от фунеции, то результат будет - *main.customErrorerror.
// Равенство с нил не происходит, т.к значение любого интерфейса является nil, в случае, когда значение и тип интерфейса это nil.
//
// Коротко: сравниваются два разных типа - <nill> != *main.customErrorerror

package main

import "fmt"

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}
func test() *customError {
	{
		// do something
	}
	return nil
}
func main() {
	var err error
	err = test()
	fmt.Printf("%T", err)
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
