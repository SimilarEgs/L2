package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Утилита sort
// Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.
//
// Реализовать поддержку утилитой следующих ключей:
//
// -k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
// -n — сортировать по числовому значению
// -r — сортировать в обратном порядке
// -u — не выводить повторяющиеся строки
//
// Дополнительно
//
// Реализовать поддержку утилитой следующих ключей:
//
// -M — сортировать по названию месяца
// -b — игнорировать хвостовые пробелы
// -c — проверять отсортированы ли данные
// -h — сортировать по числовому значению с учетом суффиксов

const (
	errArgs = "[Error] check for argument correctnes"
)

type Args struct {
	k          int
	n, r, u, c bool

	inputFilename  string
	outputFilename string
}

func main() {

	args, err := getFlags()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(args)

}

//  SetFlags function - scans cli flags and sets return «Args» struct with corresponding provided flags for further use of sort
func getFlags() (*Args, error) {

	n := flag.Bool("n", false, "sorting on digits")
	c := flag.Bool("c", false, "check if the data is sorted")
	k := flag.Int("k", 0, "define on witch column apply sort")
	u := flag.Bool("u", false, "check if outputs only unique strings")
	r := flag.Bool("r", false, "reverse sorting")

	flag.Parse()

	args := &Args{
		k: *k,
		n: *n,
		r: *r,
		u: *u,
		c: *c,
	}

	if args.k < 0 {
		return nil, errors.New(errArgs)
	}

	args.inputFilename = flag.Arg(0)
	args.outputFilename = flag.Arg(1)

	// thorw an error => if there is no data about the file names
	if args.inputFilename == "" || args.outputFilename == "" {
		return nil, errors.New(errArgs)
	}

	return args, nil

}

// ReadLines function - take filename as args, read whole file and returns two-dimensional slice (columns and lines)
func readFile(fileName string) (res [][]string, err error) {

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	scaner := bufio.NewScanner(f)

	for scaner.Scan() {

		var words []string

		words = strings.Split(scaner.Text(), " ") // split words by columns
		res = append(res, words)                  // append splited slice

	}
	return res, err

}

// writeFile function - take input-data and file-name as args, and write input-data to the output destination
func writeFile(input [][]string, fileName string) error {

	// create new-file and name it with provided fileName
	file, err := os.Create(fileName + ".txt")
	if err != nil {
		return err
	}

	defer file.Close()

	lines := make([]string, len(input))

	// converting input data to the plain text
	for i, data := range input {
		str := strings.Join(data, " ")
		lines[i] = str
	}

	// write output data to the file
	_, err = file.WriteString(strings.Join(lines, "\n"))

	return err
}
