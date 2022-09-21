package main

import (
	"fmt"
	"log"
	"os"

	"github.com/beevik/ntp"
)

// Базовая задача
// Создать программу печатающую точное время с использованием NTP -библиотеки.
// Инициализировать как go module.
// Использовать библиотеку github.com/beevik/ntp.
// Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
// Требования:
// Программа должна быть оформлена как go module
// Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Printf("[Error] %s\n", err.Error())
		os.Exit(0)
	}
	fmt.Printf("current time: %s\n", time)
}
