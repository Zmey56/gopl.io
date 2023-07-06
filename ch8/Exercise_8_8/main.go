//Exercise 8.8:
//Using a select statement, add a timeout to the echo server from Section 8.3 so that it disconnects any
//client that shouts nothing within 10 seconds.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	timeout := time.After(10 * time.Second) // Set the timeout to 10 seconds

	for {
		select {
		case <-timeout:
			fmt.Fprintln(c, "Timeout: No input received. Disconnecting...")
			c.Close()
			return
		default:
			if input.Scan() {
				go echo(c, input.Text(), 1*time.Second)
			} else {
				// NOTE: ignoring potential errors from input.Err()
				c.Close()
				return
			}
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}

		go handleConn(conn)
	}
}
