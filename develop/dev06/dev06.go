// Утилита cut

// Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

// Реализовать поддержку утилитой следующих ключей:
// -f - "fields" - выбрать поля (колонки)
// -d - "delimiter" - использовать другой разделитель
// -s - "separated" - только строки с разделителем

package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Args struct - representation of the flags of the cut command
type Args struct {
	f    []int
	d    string
	s    bool
	file string
}

const (
	errArg = "[Error] check for argument correctnes"
)

// getArgs func - scans cut flags and returns corresponding struct with entered args
func getArgs() (*Args, error) {

	f := flag.String("f", "", "fields: select only these fields;  also print any line that contains no delimiter character, unless the -s option is specified")
	d := flag.String("d", "\t", "delimiter: use DELIM instead of TAB for field delimiter")
	s := flag.Bool("s", false, "separator: do not print lines not containing delimiters")

	flag.Parse()

	// handle zero flags input
	if len(os.Args) < 3 {
		return nil, errors.New(errArg)
	}

	if len(flag.Args()) < 1 {
		return nil, errors.New(errArg)
	}

	commaFields := strings.Split(*f, ",")

	fileds := make([]int, len(commaFields))

	// parsing f flag
	for i := range commaFields {
		num, err := strconv.Atoi(commaFields[i])
		if err != nil {
			errF := fmt.Sprint(err.Error())
			return nil, errors.New("[Error] occurred while parsing field flag: " + errF)
		}
		fileds[i] = num
	}

	args := &Args{
		f: fileds,
		d: *d,
		s: *s,
	}

	// scan input file
	args.file = flag.Args()[0]

	return args, nil
}

// ReadFile func - reads the entire file and returns its contents
func readFile(fileName string) ([]string, error) {

	lines := make([]string, 0, 10)

	f, err := os.Open(fileName)
	if err != nil {
		errF := fmt.Sprint(err.Error())
		return nil, errors.New("[Error] occurred while oppening provided file: " + errF)
	}

	defer f.Close()

	scaner := bufio.NewScanner(f)

	for scaner.Scan() {
		lines = append(lines, scaner.Text())

	}

	if err := scaner.Err(); err != nil {
		errF := fmt.Sprint(err.Error())
		return nil, errors.New("[Error] occured while reading provided file: " + errF)
	}

	return lines, nil
}

// cut func - main logic of the proggram. Inside, calls all the subfunctions (getArgs, readFile)
// returns to handled lines to STDOUT
func cut() ([]string, error) {

	args, err := getArgs()
	if err != nil {
		return nil, err
	}

	lines, err := readFile(args.file)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0)

	delimiter := "\t"

	if args.d != delimiter {
		delimiter = args.d
	}

	for _, line := range lines {

		if delimiter != "" && strings.Contains(line, delimiter) {
			// splitting by delimiter (TAB)
			words := strings.Split(line, delimiter)

			// creating string buffer for handy concatenation
			cut := bytes.Buffer{}

			for _, val := range args.f {
				if len(words) >= val {
					cut.WriteString(words[val-1]) // write provided field to the buffer
					cut.WriteString(delimiter)
				}
				// append res, and TrimSuffix for an unnecessary delimiter
				res = append(res, strings.TrimSuffix(cut.String(), delimiter))
			}

		} else if !args.s {
			fmt.Println(line)
			res = append(res, line)
		}
	}

	return res, nil
}

func main() {
	lines, err := cut()
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

}
