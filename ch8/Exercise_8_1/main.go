//“Exercise 8.1:
//Modify clock2 to accept a port number, and
//write a program, clockwall, that acts as a client of several
//clock servers at once, reading the times from each one and displaying
//the results in a table, akin to the wall of clocks seen in some
//business offices.
//
//If you have access to geographically distributed computers, run
//instances remotely; otherwise run local instances on different ports
//with fake time zones.
//
//
//Click here to view code image
//$ TZ=US/Eastern    ./clock2 -port 8010 &
//$ TZ=Asia/Tokyo    ./clock2 -port 8020 &
//$ TZ=Europe/London ./clock2 -port 8030 &
//$ clockwall NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030”

package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

// parameter: NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
func main() {
	for _, zone := range os.Args[2:] {
		link := strings.Split(zone, "=")
		go getTime(link[1])
	}

	s := strings.Split(os.Args[1], "=")
	getTime(s[1])
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func getTime(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	mustCopy(os.Stdout, conn)
}

//$ TZ=US/Eastern    ./clock -port 8010 &
//$ TZ=Asia/Tokyo    ./clock -port 8020 &
//$ TZ=Europe/London ./clock -port 8030 &

//./main NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
