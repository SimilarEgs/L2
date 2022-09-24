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
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type Args struct {
	A, B, C       int
	c, i, v, F, n bool

	search string
	files  []string
}

func getArgs() (*Args, error) {

	A := flag.Int("A", 0, "Print NUM lines of trailing context after matching lines")
	B := flag.Int("B", 0, "Print NUM lines of leading context before matching lines")
	C := flag.Int("C", 0, "Print NUM lines of output context")
	c := flag.Bool("c", false, "Suppress normal output; instead print a count of matching lines for each input file")
	i := flag.Bool("i", false, "Ignore case distinctions in both the PATTERN and the input files")
	v := flag.Bool("v", false, "Invert the sense of matching, to select non-matching lines")
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
		return nil, errors.New("[Error] check for argument correctnes")
	}

	search := flag.Args()[0]

	// if -i flag was provided, handle that case on stage of receiving arguments
	if args.i {
		args.search = strings.ToLower(search)
	} else {
		args.search = search
	}

	args.files = append(args.files, flag.Args()[1:]...)

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
	args, err := getArgs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(args.files)
}
