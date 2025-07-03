package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	// Start goroutine to read from server and print to terminal
	go func() {
		serverReader := bufio.NewScanner(conn)
		for serverReader.Scan() {
			fmt.Println(serverReader.Text())
		}
		fmt.Println("Server disconnected.")
		os.Exit(0)
	}()

	// Read from client terminal and send to server
	clientReader := bufio.NewScanner(os.Stdin)
	for clientReader.Scan() {
		text := clientReader.Text()
		_, err := fmt.Fprintln(conn, text)
		if err != nil {
			fmt.Println("Error writing to server:", err)
			break
		}
	}
}
