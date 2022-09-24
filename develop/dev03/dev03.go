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
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	errArgs = "[Error] check for argument correctnes"
)

// representation of CLI args
type args struct {
	k          int
	n, r, u, c bool

	inputFilename  string
	outputFilename string
}

func main() {

	data, err := Sort()
	if err != nil {
		log.Fatalln(err)
	}
	if len(data) > 1 {
		for _, elem := range data {
			fmt.Println(elem)
		}
	}
}

// Sort function - reads name of the input and output file from the cli, scans provided flags, sorts data with corresponding method based on the flags
// and write the result to the otput file
func Sort() ([][]string, error) {

	args, err := getFlags()
	if err != nil {
		return nil, err
	}

	data, err := readFile(args.inputFilename)
	if err != nil {
		return nil, err
	}

	// representation of sort function for sort.Slice method
	var sortFunc func(i, j int) bool

	switch true {
	// digits sort
	case args.n:
		sortFunc = func(i, j int) bool {
			a, _ := strconv.ParseFloat(getElem(data, i, args.k), 64)
			b, _ := strconv.ParseFloat(getElem(data, j, args.k), 64)
			if args.r { // reverse sorting
				return a > b
			}
			return a < b
		}

	// returns only unique output
	case args.u:
		res := make([][]string, 0, 10)

		tmp := make([]string, 0, 10)
		set := make(map[string]bool)

		for _, v := range data {
			tmp = append(tmp, v...)
		}

		for _, elem := range tmp {

			var words []string

			if _, ok := set[elem]; !ok {
				words = strings.Split(elem, " ")
				set[elem] = true
				res = append(res, words)
			}
		}

		writeFile(res, args.outputFilename)
		return res, nil

	// basic sort
	default:
		sortFunc = func(i, j int) bool {
			if args.r {
				return getElem(data, i, args.k) > getElem(data, j, args.k)
			}
			return getElem(data, i, args.k) < getElem(data, j, args.k)
		}
	}

	// check if input data is sorted
	if args.c {
		sorted := sort.SliceIsSorted(data, sortFunc)
		log.Printf("[Info] sorting status: %t\n", sorted)
		return nil, nil
	}

	sort.Slice(data, sortFunc)

	writeFile(data, args.outputFilename)

	return data, nil

}

// getElem func - returns an element of 2d slice by the given argument
func getElem(data [][]string, i, k int) string {
	if k < len(data[i]) {

		return data[i][k]
	}
	return ""
}

// SetFlags function - scans cli flags and sets return «Args» struct with corresponding provided flags for further use of sort
func getFlags() (*args, error) {

	n := flag.Bool("n", false, "sorting on digits")
	c := flag.Bool("c", false, "check if the data is sorted")
	k := flag.Int("k", 0, "define on witch column apply sort")
	u := flag.Bool("u", false, "check if outputs only unique strings")
	r := flag.Bool("r", false, "reverse sorting")

	flag.Parse()

	args := &args{
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

// ReadFile function - take filename as args, read whole file and returns two-dimensional slice (columns and lines)
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
		res = append(res, words)

	}
	return res, err

}

// writeFile function - take input-data and file-name as args, and write input-data to the output destination
func writeFile(input [][]string, fileName string) error {

	// create new-file and name it with provided fileName
	file, err := os.Create(fileName)
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
