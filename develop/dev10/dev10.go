// Утилита telnet
//
// Реализовать простейший telnet-клиент.
//
// Примеры вызовов:
// go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123
//
// Требования:
// Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
// После подключения STDIN программы должен записываться в сокет, а данные полученные из сокета должны выводиться в STDOUT
// Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
// При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
// При подключении к несуществующему сервер, программа должна завершаться через timeout

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

// Args strcut - representation of input arguments (host, port, timeout)
type Args struct {
	host    string
	port    string
	timeout time.Duration
}

// getArgs - returns parsed arguments
func getArgs() (*Args, error) {

	// checking checking the number of entered arguments
	if len(os.Args) < 3 {
		return nil, errors.New("[Error] check for required arguments: host, port\n")
	}

	// scanning timeout flag
	var timeout time.Duration
	flag.DurationVar(&timeout, "t", time.Second*10, "set time for cancellation telnet session in seconds if connection has not been established session will cessation affter passing provided time")

	flag.Parse()

	// avoiding error if user provide only 2 arguments
	var host, port string

	if strings.Contains(os.Args[1], "-t") {
		host = os.Args[3]
		port = os.Args[4]
	} else {
		host = os.Args[1]
		port = os.Args[2]
	}

	args := &Args{
		host:    host,
		port:    port,
		timeout: timeout,
	}

	return args, nil
}

// read func - reads from socket and prints to console output data
func readSocket(conn net.Conn, errChan chan error) {

	// read from socket
	reader := bufio.NewReader(conn)

	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			errChan <- fmt.Errorf("[Error] %v\n", err)
			return
		}
		fmt.Printf("%s\n", data)
	}
}

// write func - writes results of STDIN to the socket conn
func writeSocket(conn net.Conn, errChan chan error) {

	reader := bufio.NewReader(os.Stdin)

	for {
		// read from cli
		data, err := reader.ReadBytes('\n')
		if err != nil {
			errChan <- fmt.Errorf("[Error] %v\n", err)
			return
		}

		// send data to the socket
		_, err = conn.Write(data)
		if err != nil {
			errChan <- fmt.Errorf("[Error] %v\n", err)
			return
		}
	}

}

func run(args *Args) error {

	conf := fmt.Sprintf("%s:%s", args.host, args.port)

	// establishing a connection
	conn, err := net.DialTimeout("tcp", conf, args.timeout)
	if err != nil {
		return fmt.Errorf("[Error] failed connection: %v\n", err)
	}
	defer conn.Close()

	fmt.Printf("[Info] connection with: %s - established\n", conf)

	// work with signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	errChan := make(chan error)

	go writeSocket(conn, errChan)
	go readSocket(conn, errChan)

	// listening to the signals
	select {
	case v := <-sigCh:
		fmt.Printf("[Info] connection received sig ch: %v\n", v)
	case v := <-errChan:
		fmt.Printf("%v\n", v)
	}

	return nil
}

func main() {

	args, err := getArgs()
	if err != nil {
		log.Fatal(err)
	}

	err = run(args)
	if err != nil {
		log.Fatal(err)
	}

}
