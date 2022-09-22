package main

import (
	"testing"
)

func TestStringBuilder(t *testing.T) {

	input := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
	}

	expected := []string{
		"aaaabccddddde",
		"abcd",
		"",
		"",
	}

	for i := range input {
		result, _ := StringBuilder(input[i])
		if result != expected[i] {
			t.Errorf("%s != %s\n", expected[i], result)
		}

	}

}
