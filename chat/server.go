package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	message  = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-message:
			fmt.Println("Broadcast ==> ", msg)
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)

	go writeToClient(conn, ch)

	who := conn.RemoteAddr().String()

	ch <- "You are " + who

	fmt.Printf("%s are arrived\n", who)

	message <- who + " are arrived"

	entering <- ch

	input := bufio.NewScanner(conn)

	for input.Scan() {
		message <- who + "==>" + input.Text()
	}

	leaving <- ch
	message <- who + " are left"

	//fmt.Printf("%s are left\n", who)
	conn.Close()
}

func writeToClient(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Printf("Received: %s\n", msg)
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
