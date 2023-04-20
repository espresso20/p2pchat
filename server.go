package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

type Client struct {
	id   string
	name string
	conn net.Conn
}

var clients sync.Map

func handleConnection(client *Client) {
	defer client.conn.Close()

	scanner := bufio.NewScanner(client.conn)
	for scanner.Scan() {
		msg := scanner.Text()

		if strings.HasPrefix(msg, "NAME:") {
			client.name = strings.TrimPrefix(msg, "NAME:")
			continue
		}

		broadcast(client.id, client.name, msg)
	}
}

func broadcast(senderID, senderName, msg string) {
	clients.Range(func(key, value interface{}) bool {
		client := value.(*Client)
		if client.id != senderID {
			fmt.Fprintf(client.conn, "[%s]: %s\n", senderName, msg)
		}
		return true
	})
}

func main() {
	port := "8080"

	cert, err := tls.LoadX509KeyPair("../cert.pem", "../key.pem")
	if err != nil {
		fmt.Println("Error loading certificate and key:", err)
		os.Exit(1)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	listener, err := tls.Listen("tcp", ":"+port, config)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Printf("Secure chat server started on port %s...\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		clientID := strings.TrimSpace(conn.RemoteAddr().String())
		client := &Client{
			id:   clientID,
			conn: conn,
		}
		clients.Store(clientID, client)

		go handleConnection(client)
	}
}
