package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
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
	port string
	host string
	ttl  time.Duration
}

func New() *TcpConnection {
	conn := TcpConnection{}
	flag.DurationVar(&conn.ttl, "timeout", 10*time.Second, "количество времени для соединения")
	flag.Parse()
	args := flag.Args()
	conn.host = args[0]
	conn.port = args[1]
	return &conn
}

func main() {
	c := New()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(c.host, c.port), c.ttl)
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
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
