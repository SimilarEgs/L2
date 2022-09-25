// Реализовать утилиту фильтрации (man grep)
// Поддержать флаги:
// -A - "after" печатать +N строк после совпадения
// -B - "before" печатать +N строк до совпадения
// -C - "context" (A+B) печатать ±N строк вокруг совпадения
// -c - "count" (количество строк)
// -i - "ignore-case" (игнорировать регистр)
// -v - "invert" (вместо совпадения, исключать)
// -F - "fixed", точное совпадение со строкой, не паттерн
// -n - "line num", печатать номер строки
// Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	errArg = "[Error] check for argument correctnes"
)

// Args struct - representation of grep arguments
type Args struct {
	A, B, C       int
	c, i, v, F, n bool

	search string
	file   string
}

// grep func - main logic of the proggram. Inside, calls all the subfunctions (getArgs, readFile)
// and perform required computation to find coresponding patterns inside provided file
func grep() error {

	args, err := getArgs()
	if err != nil {
		return err
	}

	rowsIndex := make([]int, 0, 10)

	rows, err := readFile(args.file)
	if err != nil {
		return err
	}

	for i, v := range rows {

		// flag -i "ignore-case"
		if args.i {
			v = strings.ToLower(v)
		}

		// flag -F "fixed"
		if args.F {
			if strings.Contains(v, args.search) {
				rowsIndex = append(rowsIndex, i)
			}
			// if flag not specified, find all lines that match the pattern(«args.search»)
		} else {
			match, err := regexp.MatchString(args.search, v)
			if err != nil {
				continue
			}

			// if provided pattern matches the value of current loop iteration,
			// append rowsIndedx with corresponding line index
			if match {
				rowsIndex = append(rowsIndex, i)
			}

		}
	}

	// flag -n "line num"
	if args.n {

		res := make([]string, len(rows))

		for i := range rows {
			nLine := bytes.Buffer{}

			number := fmt.Sprintf("%d", i+1)
			nLine.Grow(len(number) + 1 + len(rows[i]))

			// adding line number to the buffer
			nLine.WriteString(number)
			nLine.WriteString(":")
			nLine.WriteString(rows[i])

			// write resualt
			res[i] = nLine.String()

		}
		rows = res
	}

	// flag -c "count"
	if args.c {
		fmt.Printf("[Info] count of matches: %d\n", len(rowsIndex))
		return nil
	}

	// flag -v "invert"
	if args.v {
		for i := 0; i < len(rows); i++ {

			// func for checking line index => if true, return «false» and do not output matched line index
			match := func(rowsIndex []int, index int) bool {
				for i := 0; i < len(rowsIndex); i++ {
					if index == rowsIndex[i] {
						return false
					}
				}
				return true
			}(rowsIndex, i)

			if match {
				fmt.Println(rows[i])
			}
		}
	}

	if args.C > 0 {

	}

	indexes := getIndexes(rowsIndex)

	// if flags was not specified print all lines that matches pattern
	for i := range rows {
		if indexes[i] {
			fmt.Println(rows[i])
		}
	}

	return nil

}

// getIndex func - returns map filled with indexes of all matching lines
func getIndexes(rowsIndex []int) map[int]bool {
	indexes := make(map[int]bool)
	for _, v := range rowsIndex {
		indexes[v] = true
	}
	return indexes
}

// getArgs func - scans grep flags and returns corresponding struct with entered args
func getArgs() (*Args, error) {

	A := flag.Int("A", 0, "Print NUM rows of trailing context after matching rows")
	B := flag.Int("B", 0, "Print NUM rows of leading context before matching rows")
	C := flag.Int("C", 0, "Print NUM rows of output context")
	c := flag.Bool("c", false, "Suppress normal output; instead print a count of matching rows for each input file")
	i := flag.Bool("i", false, "Ignore case distinctions in both the PATTERN and the input files")
	v := flag.Bool("v", false, "Invert the sense of matching, to select non-matching rows")
	F := flag.Bool("F", false, "Interpret PATTERN as a list of fixed strings")
	n := flag.Bool("n", false, "Prefix each line of output with the line number within its input file")

	flag.Parse()

	args := &Args{
		A: *A,
		B: *B,
		C: *C,
		c: *c,
		i: *i,
		v: *v,
		F: *F,
		n: *n,
	}

	// throw an error if cli args are empty
	if len(flag.Args()) < 1 {
		return nil, errors.New(errArg)
	}

	search := flag.Args()[0]

	// if -i flag was provided, handle that case on stage of receiving arguments
	if args.i {
		args.search = strings.ToLower(search)
	} else {
		args.search = search
	}

	args.file = flag.Args()[1]

	return args, nil

}

// ReadFile func - reads the entire file and returns its contents
func readFile(fileName string) ([]string, error) {

	result := make([]string, 0, 10)

	f, err := os.Open(fileName)
	if err != nil {
		errF := fmt.Sprint(err.Error())
		return nil, errors.New("[Error] occured while oppening provided file: " + errF)
	}

	defer f.Close()

	scaner := bufio.NewScanner(f)

	for scaner.Scan() {
		result = append(result, scaner.Text())
	}

	if err := scaner.Err(); err != nil {
		errF := fmt.Sprint(err.Error())
		return nil, errors.New("[Error] occured while reading provided file: " + errF)
	}

	return result, nil
}

func main() {

	err := grep()
	if err != nil {
		log.Fatal(err)
	}
}
