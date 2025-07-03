package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

type Client struct {
	conn net.Conn
}

var (
	clients   = make(map[net.Conn]Client) // track all connected clients
	clientsMu sync.Mutex                  // protect clients map concurrent access
)

func main() {

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server is listening on port 8080...")

	go func() {
		for {
			conn, err := listener.Accept()

			if err != nil {
				fmt.Println("Error accepting connection:", err)

			}

			clientsMu.Lock()
			clients[conn] = Client{conn: conn}
			clientsMu.Unlock()

			fmt.Println("Client connected from:", conn.RemoteAddr())

			go handleClient(conn)
		}
	}()

	// Read from server terminal and send to client
	serverReader := bufio.NewScanner(os.Stdin)
	for serverReader.Scan() {
		msg := serverReader.Text()
		broadcast("Server : " + msg)

	}

}

// to read from client and print to terminal
func handleClient(conn net.Conn) {
	clientReader := bufio.NewScanner(conn)
	for clientReader.Scan() {
		text := clientReader.Text()
		fmt.Printf("Message from %s: %s\n", conn.RemoteAddr(), text)
	}

	fmt.Println("Client disconnected:", conn.RemoteAddr())

	clientsMu.Lock()
	delete(clients, conn)
	clientsMu.Unlock()
	conn.Close()
}

func broadcast(message string) {
	clientsMu.Lock()
	defer clientsMu.Unlock()

	for _, client := range clients {
		_, err := fmt.Fprintln(client.conn, message)
		if err != nil {
			fmt.Println("Error broadcasting to", client.conn.RemoteAddr(), ":", err)
		}
	}
}
