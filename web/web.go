package web

import (
	"bufio"
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

	server, err := new.Listen("tcp", ":"+"2556")
	if err != nil {
		log.Fatal("couldn't start server: " + err.String())
	}

	conn, err := clientConns(server)

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
				fmt.Println("couldn't accept: " + err.String())
				continue
			}
			i++
			mt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
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
			break
		}
		client.Write(line)
	}
}
