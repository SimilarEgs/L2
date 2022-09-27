// Or channel

// Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
// Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
// однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
// В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.
// Определение функции:
// var or func(channels ...<- chan interface{}) <- chan interface{}
// Пример использования функции:
// sig := func(after time.Duration) <- chan interface{} {
// 	c := make(chan interface{})
// 	go func() {
// 		defer close(c)
// 		time.Sleep(after)
// }()
// return c
// }
// start := time.Now()
// <-or (
// 	sig(2*time.Hour),
// 	sig(5*time.Minute),
// 	sig(1*time.Second),
// 	sig(1*time.Hour),
// 	sig(1*time.Minute),
// )
// fmt.Printf(“fone after %v”, time.Since(start))

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-merge(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
	)
	fmt.Printf("fone after %v", time.Since(start))
}

// merge func - recives multiple channels and merge them into single channel
func merge(cs ...<-chan interface{}) <-chan interface{} {

	// a single out chan that recives all values from the input channles
	out := make(chan interface{})

	wg := sync.WaitGroup{}

	wg.Add(len(cs))

	// traversing through input channels and writing their values to the output channel
	for _, c := range cs {
		go func(c <-chan interface{}) { // avoiding race condition by passing c to the func argument
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(c)
	}

	// background gorutina with block, to close out channel, after reciving all values
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
