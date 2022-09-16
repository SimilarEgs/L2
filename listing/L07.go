// Что выведет программа? Объяснить вывод программы.
//
// Ответ:
// Последовательность чисел с новой строки от 1 до 8, а затем «бесконечное» кол-во нулей
//
// Почему:
// В цикле for-range мы читаем канал с, который в то же время, получает некие значения из каналов (a, b)
// Далее, в канал c перестают поступать данны и мы «бесконечно» выводим значение по умолчанию канала c, то бишь 0
//
// Решение:
// В функции merge задать циклу булевое условие для его завершения.
// Когда каналы a, b закроются мы выйдем из цикла, и закроем канал c.

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

	aDone, bDone := false, false

	go func() {
		defer close(c)

		for !aDone || !bDone {

			select {
			case v, ok := <-a:
				if !ok {
					aDone = true
					continue
				}
				c <- v
			case v, ok := <-b:
				if !ok {
					bDone = true
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
