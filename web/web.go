package web

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type request struct {
	Key         string
	MatchString string
	Row         int
	Col         int
}

func StartServer() {

	server, err := net.Listen("tcp", ":"+"2556")
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	conns := clientConns(server)

	for {
		go handleConn(<-conns)

	}

}

func clientConns(listener net.Listener) chan net.Conn {
	ch := make(chan net.Conn)
	i := 0

	go func() {
		for {
			client, err := listener.Accept()
			if err != nil {
				fmt.Println("couldn't accept: " + err.Error())
				continue
			}
			i++
			fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

func handleConn(client net.Conn) {
	b := bufio.NewReader(client)
	for {
		line, err := b.ReadBytes('\n')
		if err != nil {
			client.Write([]byte("500"))
			break
		}
		if err == nil {
			parseErr := parseRequest(line)
			client.Write([]byte(parseErr))
		}

	}
}

func parseRequest(jsonRequst []byte) string {
	r := new(request)
	err := json.Unmarshal(jsonRequst, &r)
	if err != nil {
		return ("500 " + err.Error())
	}
	return ("200")

}
