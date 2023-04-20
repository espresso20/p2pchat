package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: client [server address] [port] [name]")
		os.Exit(1)
	}

	serverAddr := os.Args[1]
	port := os.Args[2]
	name := os.Args[3]

	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", serverAddr+":"+port, config)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Send the client's name to the server
	fmt.Fprintf(conn, "NAME:%s\n", name)

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		fmt.Fprintf(conn, "%s\n", msg)
	}
}
