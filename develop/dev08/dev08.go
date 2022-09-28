// Взаимодействие с ОС
//
// Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:
//
// - cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
// - pwd - показать путь до текущего каталога
// - echo <args> - вывод аргумента в STDOUT
// - kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
// - ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*
//
//
// Так же требуется поддерживать функционал fork/exec-команд
//
// Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).
//
// *Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
// в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
// и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).

package main

// "bufio"
// "fmt"
// "log"
// "os"
// "os/exec"
// "path/filepath"
// "strconv"
// "strings"

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

func main() {

	scaner := bufio.NewScanner(os.Stdin)

	for {

		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		// current absolute path
		abs, err := filepath.Abs(pwd)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(abs, ":~$ ")

		// scan STDIN
		scaner.Scan()
		action := scaner.Text()
		fmt.Println()

		// spliting cli input for handy work
		command := strings.Fields(action)

		args := command[1:]

		// main logic
		switch command[0] {
		case "quit":
			fmt.Println("[Info] godbye")
			os.Exit(0)
		case "pwd":
			fmt.Printf("%s\n\n", abs)
		case "echo":
			line := strings.Join(args, " ")
			line += "\n"
			fmt.Println(line)
		case "cd":
			if len(args) < 1 {
				err := errors.New("check for argument correctnes")
				log.Printf("[Error] %v\n\n", err)
				continue
			}

			path := args[0]
			err := os.Chdir(path)
			if err != nil {
				log.Printf("[Error] %v\n\n", err)
			}

		case "ps":
			// V1========================================
			// matches, err := filepath.Glob("/proc/*/exe")
			// if err != nil {
			// 	fmt.Println(err)
			// 	continue
			// }
			// for _, file := range matches {
			// 	target, _ := os.Readlink(file)
			// 	// process := filepath.Base(target)
			// 	if len(target) > 0 {
			// 		fmt.Printf("%+v\n", target)
			// 	}
			// }
			// fmt.Println()

			// V2========================================
			// process, err := process.Processes()
			// if err != nil {
			// 	log.Printf("[Error] %v\n\n", err)
			// 	continue
			// }
			// fmt.Printf("PID\tNAME\t\tTIME\n")
			// for _, ps := range process {
			// 	psName, _ := ps.Name()

			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
			// 	fmt.Printf("%d\t%s\n", ps.Pid, psName)
			// }

			// V3========================================
			processList, err := ps.Processes()
			if err != nil {
				log.Printf("[Error] %v\n\n", err)
				continue
			}

			fmt.Printf("PID\tNAME\n")

			for x := range processList {
				var process ps.Process
				process = processList[x]
				fmt.Printf("%d\t%s\n", process.Pid(), process.Executable())
			}
			fmt.Println()
		case "kill":
			if len(args) < 1 {
				err := errors.New("check for argument correctnes")
				log.Printf("[Error] %v\n\n", err)
				continue
			}

			pid, err := strconv.Atoi(args[0])
			if err != nil {
				log.Printf("[Error] %v\n\n", err)
				continue
			}

			ps, err := os.FindProcess(pid)
			if err != nil {
				log.Printf("[Error] %v\n\n", err)
				continue
			}

			err = ps.Kill()
			if err != nil {
				log.Printf("[Error] %v\n\n", err)
				continue
			}
		default:
			fmt.Printf("[Error] %s: command not found\n\n", command[0])
		}

	}

}
