package main

import (
	"strings"
	"testing"
)

func TestGet(t *testing.T) {

	expected := map[string][]string{
		"банка":   {"банка", "кабан"},
		"камыш":   {"камыш", "мышка"},
		"коршун":  {"коршун", "шнурок"},
		"пятка":   {"пятка", "тяпка"},
		"росинка": {"росинка", "соринка"},
		"листок":  {"листок", "столик"},
	}

	input := []string{
		"Кабан",
		"банка",
		"мышка",
		"Камыш",
		"Соринка",
		"росинка",
		"пятка",
		"ТЯпка",
		"лтсток",
		"коршун",
		"шнурок",
		"Стилок",
	}

	result := anagrams(input)

	for k, v := range result {
		var (
			anag []string
			ok   bool
		)
		if anag, ok = expected[k]; !ok {
			t.Errorf("extra key: %s", k)
		}

		joinedSetAnagrams := strings.Join(anag, " ")
		joinedResAnagrams := strings.Join(v, " ")

		if joinedResAnagrams != joinedSetAnagrams {
			t.Errorf("diverges result:\nexpected: %s\ngot: %s\n", joinedSetAnagrams, joinedResAnagrams)
		}

	}

}
