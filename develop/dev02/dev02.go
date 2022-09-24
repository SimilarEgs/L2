// Задача на распаковку
// Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
// 	- "a4bc2d5e" => "aaaabccddddde"
// 	- "abcd" => "abcd"
// 	- "45" => "" (некорректная строка)
// 	- "" => ""
// Дополнительное задание: поддержка escape - последовательностей
// 	- qwe\4\5 => qwe45 (*)
// 	- qwe\45 => qwe44444 (*)
// 	- qwe\\5 => qwe\\\\\ (*)
// В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.
// Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.

// go vet - pass
// golint - pass
// tests - pass
// additional tasks - failed

package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	defer func() {
		if re := recover(); re != nil {
			log.Println("[Error] incorrect string")
		}
	}()

	var str string

	fmt.Print("[Info] input string: ")
	fmt.Scanln(&str)

	res, err := StringBuilder(str)
	if err != nil {
		log.Printf("%s", err.Error())
		return
	}

	fmt.Printf("[Info] resualt: %s\n", res)

}

var re = regexp.MustCompile("[0-9]+")

// StringBuilder function - unpack input string and return result of unpacking
func StringBuilder(str string) (string, error) {

	// check for empty string
	if str == "" {
		return str, nil
	}

	strRune := []rune(str)

	// check for first-element error
	if _, err := strconv.Atoi(string(strRune[0])); err == nil {
		return "", errors.New("[Error] incorrect string")
	}

	var buffer bytes.Buffer

	for i := range strRune {

		x := re.FindAllString(string(strRune[i]), 1)

		// condition for checking digits
		// if len(x) equal 0, then it's a symbol, so write it to the buffer
		// else => continue
		if len(x) != 1 {
			buffer.WriteString(string(strRune[i]))
			continue
		} else {
			n, err := strconv.Atoi(x[0])
			if err != nil {
				log.Fatal(err.Error())
			}
			// take the iteration digit «n» and the previous char «strRune[i-1]»
			// write char to the buffer n times via str.Repeat
			buffer.WriteString(strings.Repeat(string(strRune[i-1]), n-1))
		}
	}

	return fmt.Sprint(buffer.String()), nil
}
