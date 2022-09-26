package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

type argsT struct {
	file    string
	flags   []string
	pattern string
}

func TestGrep(t *testing.T) {

	cases := []argsT{
		// {
		// 	file:    "bash_help.txt",
		// 	pattern: "arg",
		// 	flags:   []string{"-F", "-n"},
		// },
		// {
		// 	file:    "bash_help.txt",
		// 	pattern: "arg",
		// 	flags:   []string{"-v", "-c"},
		// },
		// {
		// 	file:    "go_help.txt",
		// 	pattern: "get",
		// 	flags:   []string{"-v", "-n", "-A", "2"},
		// },
		// {
		// 	file:    "go_help.txt",
		// 	pattern: "go",
		// 	flags:   []string{"-v", "-n"},
		// },
		{
			file:    "go_help.txt",
			pattern: "bug",
			flags:   []string{"-A", "5", "-n"},
		},
	}

	for _, cases := range cases {
		command := []string{"run", "dev05.go"}
		command = append(command, cases.pattern)
		command = append(command, cases.file)

		dev05, err := exec.Command("go", command...).CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[Error] test failed: %v\n", err)
			os.Exit(1)
		}

		command = command[0:0]
		command = append(command, cases.pattern)
		command = append(command, cases.file)

		grep, err := exec.Command("grep", command...).CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "[Error] test failed: %v\n", err)
			os.Exit(1)
		}

		for i := range grep {
			if grep[i] != dev05[i] {
				fmt.Println("Dev05 test:\n", string(dev05))
				fmt.Println("Grep test:\n", string(grep))

				t.Errorf("[Error] expected:  %s\ngot: %s\n", string(grep), string(dev05))
				os.Exit(1)
			}
		}
	}
}
