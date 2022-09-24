// Напишите функцию поиска всех множеств анаграмм по словарю.
// Например:
// 'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
// 'листок', 'слиток' и 'столик' - другому.
// Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
// Выходные данные: Ссылка на мапу множеств анаграмм.
// Ключ - первое встретившееся в словаре слово из множества
// Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
// Множества из одного элемента не должны попасть в результат.
// Все слова должны быть приведены к нижнему регистру.
// В результате каждое слово должно встречаться только один раз.
// Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

// go vet - pass
// golint - pass
// tests - pass

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	dictionary := []string{
		"Кабан",
		"банка",
		"мышка",
		"Камыш",
		"Соринка",
		"росинка",
		"пятка",
		"ТЯпка",
		"Листок",
		"коршун",
		"шнурок",
		"листок",
	}
	fmt.Println(dictionary)
}

// Anagrams func - return sets of anagrams
func anagrams(words []string) map[string][]string {

	tmp := make(map[string][]string)

	// filling temp set
	for _, word := range words {

		lowerW := strings.ToLower(word)

		// split + sort input data to make values of the same kind, for further separations
		v := strings.Split(lowerW, "")

		sort.Strings(v)
		letters := strings.Join(v, "")

		tmp[letters] = append(tmp[letters], lowerW)

	}

	cleaner(tmp)

	anagrams := make(map[string][]string, len(tmp))

	for _, v := range tmp {
		sort.Strings(v)
		anagrams[v[0]] = v
	}

	return anagrams
}

// Cleaner func - handle input set, by removing unnecessary values
func cleaner(tmp map[string][]string) {

	unique := make(map[string]bool)

	for k, v := range tmp {

		// removing short sets
		if len(v) < 2 {
			delete(tmp, k)
		}

		// removing dupes
		for i := range v {
			if !unique[v[i]] {
				unique[v[i]] = true
			} else {
				v[i] = v[len(v)-1]
				tmp[k] = v[:len(v)-1]
			}
		}

	}

}
