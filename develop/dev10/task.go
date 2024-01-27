package main

import (
	"bufio"
	"errors"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

type TcpConnection struct {
	port int
	host string
	ttl  time.Duration
}

func flags() (TcpConnection, error) {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		return TcpConnection{}, errors.New("Недостаточно аргументов")
	}
	var (
		host      string
		port, ttl int
	)
	for _, arg := range args {
		a := strings.Split(arg, "=")
		if len(a) == 1 {
			return TcpConnection{}, errors.New("Отсутствует или неправльный аргумент")
		}
		if a[0] == "host" {
			host = a[1]
			continue
		}
		if a[0] == "port" {
			p, err := strconv.Atoi(a[1])
			if err != nil {
				return TcpConnection{}, errors.New("Невозможно распарсить порт")
			}
			port = p
			continue
		}
		if a[0] == "timeout" {
			t, err := strconv.Atoi(a[1])
			if err != nil {
				return TcpConnection{}, errors.New("Невозможно распарсить timeout")
			}
			ttl = t
			continue
		}
		return TcpConnection{}, errors.New("Неизветные аргументы")
	}
	if host == "" {
		return TcpConnection{}, errors.New("Нет аргумента хоста")
	}
	return TcpConnection{
		port: port,
		host: host,
		ttl:  time.Duration(ttl) * time.Second,
	}, nil
}

func main() {
	c, err := flags()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(c.host, strconv.Itoa(c.port)), c.ttl)
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	go func(conn net.Conn) {
		input := bufio.NewReader(os.Stdin)
		for {
			data, err := input.ReadString('\n')
			if err != nil {
				if err = conn.Close(); err != nil {
					log.Fatal(err)
				}
				log.Fatal(err)
			}
			if _, err = conn.Write([]byte(data)); err != nil {
				conn.Close()
				log.Fatal(err)
			}
		}
	}(conn)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGQUIT)
	<-sig
}
