// Что выведет программа? Объяснить вывод программы.
//
// Ответ: последовательность числе от 0 до 9 + deadlock
//
// Почему:
// Внутри анонимной горуитины мы записываем данные в канал, а после, читаем данные из канала в цикле range
// Т.к range получает данные из канала, до тех пор, пока он не закрыт, range будет весеть в ожидании новых данных, которые никогда в него не придут
// что и породит ошибку deadlock
//
// Решение:
// Закрыть канал после цикла в анонимной горутине

package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}
