// Реализовать утилиту wget с возможностью скачивать сайты целиком
// Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.

// go vet - pass
// golint - pass
// tests - failed

package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

const (
	errArg = "[Error] check for argument correctnes"
)

func main() {

	// scan cli flags
	uri := flag.String("s", "", "Uniform Resource Identifier, specify base url for successful download of the site")
	flag.Parse()

	// checking argument correctnes
	if len(os.Args) < 2 {
		log.Println(errArg)
		return
	}

	// checking the format of arguments
	if ok, err := regexp.MatchString("^(http|https)://", *uri); !ok || err != nil {
		log.Println(errArg + ", invalid url format")
		return
	}

	if err := wget(*uri); err != nil {
		log.Fatalf("[Error] %v\n", err)
	}

}

// wget func - write resualts of a GET request to a file.
func wget(url string) error {

	// make request to a provided url
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// do response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// creating new file with the contents of the response result
	file, err := os.Create("index.html")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// write response to the file
	_, err = io.Copy(writer, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
