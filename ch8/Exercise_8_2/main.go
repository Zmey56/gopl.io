//Implement a concurrent File Transfer Protocol (FTP) server. The server should interpret commands from
//each client such as cd to chagne directore, ls to list a directory, get to send the contents of a file,
//and close to close the connection. You can use the standard ftp command as the client, or write your own.

package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":2121")
	if err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("FTP server started on port 2121")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Println("New client connected:", conn.RemoteAddr())

	conn.Write([]byte("220 Welcome to the FTP server\r\n"))

	for {
		cmd, err := readCommand(conn)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading command:", err)
			}
			return
		}

		response := executeCommand(cmd)
		conn.Write([]byte(response))
	}
}

func readCommand(conn net.Conn) (string, error) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}

	command := string(buffer[:n])
	return strings.TrimSpace(command), nil
}

func executeCommand(cmd string) string {
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return "500 Invalid command\r\n"
	}

	switch parts[0] {
	case "cd":
		return changeDirectory(parts[1])
	case "ls":
		return listDirectory(parts[1])
	case "get":
		return getFile(parts[1])
	case "close":
		return "221 Goodbye\r\n"
	default:
		return "500 Unknown command\r\n"
	}
}

func changeDirectory(dir string) string {
	err := os.Chdir(dir)
	if err != nil {
		return fmt.Sprintf("550 Failed to change directory: %v\r\n", err)
	}

	return "250 Directory changed\r\n"
}

func listDirectory(dir string) string {
	fileList := ""

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path != dir {
			fileList += info.Name() + "\r\n"
		}

		return nil
	})

	if err != nil {
		return fmt.Sprintf("550 Failed to list directory: %v\r\n", err)
	}

	return "200 " + fileList + "\r\n"
}

func getFile(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Sprintf("550 Failed to read file: %v\r\n", err)
	}

	return "200 " + string(data) + "\r\n"
}
