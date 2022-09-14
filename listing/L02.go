// Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

// Ответ:
// 1
// 2
//
// Почему:
// В документации языка сказано:
// «A defer statement pushes a function call onto a list. The list of saved calls is executed after the surrounding function returns».
// «Deferred functions may read and assign to the returning function's named return values».
//
// Функция: test()
// В сигнатуре функции test() мы возвращаем именованное значение.
// Далее создаем анонимную функцию, которая заинкременит возвращаемое значение на 1, то есть в сумме мы вернем x = 2.
// Мы так же используем «голый» return, который возвращает наше именованное значение, т.к мы явно указали это в сигнатуре.
// «Naked returns don’t offer any technically unique feature, they are syntactic saccharin that harms readability. It’s ok to be a prude and say no to naked returns».
//
// Функция: anotherTest()
// Здесь, в свою очередь, мы таким же образом инкрементим x, но т.к именованное значение отсутствует
// на конечный результат возвращаемого x это не повлияет

package main

import (
	"fmt"
)

func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}
func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
