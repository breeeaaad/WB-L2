package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	server, _ := net.Listen("tcp", ":80")
	defer server.Close()
	for {
		cl, _ := server.Accept()
		sc := bufio.NewScanner(cl)
		var data string
		for sc.Scan() {
			data = sc.Text()
			fmt.Print(data)
		}
		cl.Close()
	}
}
