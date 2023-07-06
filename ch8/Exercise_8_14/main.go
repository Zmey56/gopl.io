// “Exercise 8.14:
//Change the chat server’s network protocol so that each client provides
//its name on entering.
//
//Use that name instead of the network address when prefixing each message
//with its sender’s identity”

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

type client struct {
	name         string
	msgCh        chan<- string
	lastActivity time.Time
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	var mutex sync.Mutex

	for {
		select {
		case msg := <-messages:
			mutex.Lock()
			for cli := range clients {
				cli.msgCh <- msg
			}
			mutex.Unlock()

		case cli := <-entering:
			mutex.Lock()
			clients[cli] = true
			sendCurrentClients(clients, cli.msgCh)
			mutex.Unlock()

		case cli := <-leaving:
			mutex.Lock()
			delete(clients, cli)
			close(cli.msgCh)
			mutex.Unlock()
		}
	}
}

func sendCurrentClients(clients map[client]bool, msgCh chan<- string) {
	names := make([]string, 0, len(clients))
	for cli := range clients {
		names = append(names, cli.name)
	}
	msg := "Current clients: " + strings.Join(names, ", ")
	msgCh <- msg
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	ch <- "Enter your name: "
	nameInput := bufio.NewScanner(conn)
	nameInput.Scan()
	who := nameInput.Text()

	ch <- "Welcome, " + who
	messages <- who + " has arrived"

	cli := client{
		name:         who,
		msgCh:        ch,
		lastActivity: time.Now(),
	}
	entering <- cli

	input := bufio.NewScanner(conn)
	message := make(chan string)
	go func() {
		for input.Scan() {
			message <- input.Text()
		}
		close(message)
	}()

	idleTimeout := time.NewTimer(5 * time.Minute)

	for {
		select {
		case msg, ok := <-message:
			if !ok {
				// Клиент отключился
				leaving <- cli
				messages <- who + " has left"
				conn.Close()
				return
			}
			messages <- who + ": " + msg
			cli.lastActivity = time.Now()

		case <-idleTimeout.C:
			// Превышено время бездействия, отключаем клиента
			leaving <- cli
			messages <- who + " has been disconnected due to inactivity"
			conn.Close()
			return
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
