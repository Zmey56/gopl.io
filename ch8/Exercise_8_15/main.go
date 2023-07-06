//Failure of any client program to read data in a timely manner ultimately causes all clients to get stuck.
//
//Modify the broadcaster to skip a message rather than wait if a client writer is not ready to accept it.
//
//Alternatively, add buffering to each client’s outgoing message channel so that most messages are not dropped;
//the broadcaster should use a non-blocking send to this channel.

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
	msgCh        chan string
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
				select {
				case cli.msgCh <- msg:
					// Message sent to the client successfully
				default:
					// Client writer is not ready to accept the message, skip it
				}
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

func sendCurrentClients(clients map[client]bool, msgCh chan string) {
	names := make([]string, 0, len(clients))
	for cli := range clients {
		names = append(names, cli.name)
	}
	msg := "Текущие клиенты: " + strings.Join(names, ", ")
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
		msgCh:        make(chan string, 100), // Add buffering to outgoing message channel
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
				leaving <- cli
				messages <- who + " has left"
				conn.Close()
				return
			}
			messages <- who + ": " + msg
			cli.lastActivity = time.Now()

		case <-idleTimeout.C:
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
