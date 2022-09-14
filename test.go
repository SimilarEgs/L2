package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {

	res := findDigits("test.txt")

	d := binary.BigEndian.Uint16(res)
	fmt.Println(d)
}

var regExp = regexp.MustCompile("[0-9]+")

func findDigits(filename string) []byte {

	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	b = regExp.Find(b)

	res := make([]byte, len(b))

	copy(res, b)

	return res

}
