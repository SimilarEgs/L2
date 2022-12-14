// Что выведет программа? Объяснить вывод программы.
// Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.
//
// Ответ: [3 2 3]
//
// Почему:
// Внутри «modifySlice» мы работаем с копией слайса main, то бишь, изменяя слайс внутри фрейма функции modifySlice, мы
// работает с указателем на базовый слайс s, это будет так, до применения функци append,
// затем, исходя из работы этой функции, мы будем ссылаться на новый массив, что в свою очередь, не изменит наш базовый массив в мейне.
// Поэтому все изменения после 26 строчки затронут слайс существующий только внутри этой функции.
//
// Решение:
// Вернуть i из modifySlice и присвоить результат работы этой функции переменной s
//
// Устройсво слайсов:
//
// type Slice struct {
//    arr *int  // указатель на массив
//    cap int  // емкость слайса
//    len int //  длина слайса
// }
//
// Слайс - это структура данных с полями: arr, cap, len.
// Эти свойства описывают строение базового массива, на который ссылается слайс.
// По своей сути, слайс не хранит внутри никаких данных, а лишь ссылается на нижележащий массив с этими данными.

package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
