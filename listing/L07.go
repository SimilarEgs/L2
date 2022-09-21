// Что выведет программа? Объяснить вывод программы.
//
// Ответ:
// Последовательность чисел с новой строки от 1 до 8, а затем «бесконечное» кол-во нулей
//
// Почему:
// В цикле for-range мы читаем канал с, который в то же время, получает некие значения из каналов (a, b)
// Далее, в канал c перестают поступать данны и мы «бесконечно» выводим значение по умолчанию для канала c, то бишь 0
//
// Решение:
// Внутри кейсов селекта задать проверку на закрытый канал a,b.
// В случае, если канал закроют, мы сделаем этот канал nil, что на выходе даст нам блокировку кейса, («Sending to a nil channel blocks forever»)
// дабы не выполнять лишних итераций.
// И наконец, закрыв два канала, мы выйдем из цикла по условию.
// Далее закроем канал c через defer и завершим работу программы корректно

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)

	for v := range c {
		fmt.Println(v)
	}

}

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)))
		}
		close(c)
	}()

	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)

	go func() {
		defer close(c)

		for a != nil || b != nil {

			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}
				c <- v
			}
		}
	}()

	return c
}

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func asChan(vs ...int) <-chan int {
// 	c := make(chan int)

// 	go func() {
// 		for _, v := range vs {
// 			c <- v
// 			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
// 		}

// 		close(c)
// 	}()
// 	return c
// }

// func merge(a, b <-chan int) <-chan int {
// 	c := make(chan int)
// 	go func() {
// 		for {
// 			select {
// 			case v := <-a:
// 				c <- v
// 			case v := <-b:
// 				c <- v
// 			}
// 		}
// 	}()
// 	return c
// }

// func main() {

// 	a := asChan(1, 3, 5, 7)
// 	b := asChan(2, 4, 6, 8)
// 	c := merge(a, b)
// 	for v := range c {
// 		fmt.Println(v)
// 	}
// }
