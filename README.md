# Go TCP Chat
---
## 💬 Project Overview

A simple **concurrent TCP chat server and client** written in Go.  
Multiple clients can connect to the server and chat in real-time via the terminal.

This project demonstrates:

- Basic TCP networking in Go (`net` package)  
- Handling multiple clients concurrently using goroutines and mutexes  
- Bidirectional communication between client and server  
- Synchronizing shared data safely with `sync.Mutex`  
- Simple CLI input/output with `bufio.Scanner`  

---

## 🚀 Features

- Multi-client support — multiple clients connect simultaneously  
- Broadcast messages from server to all connected clients  
- Server prints client messages in real-time  
- Client and server chat interactively via terminal input  
- Graceful client disconnect handling  

---

## 🛠️ Getting Started

### Prerequisites

- Go installed
- Terminal / Command line access  

### Installation

Clone the repo:

```bash
git clone https://github.com/sobhanagh/go-chat-multiclient.git
cd go-chat-multiclient
```
---

## 🎯 Usage
### Run the Server
```bash
cd server
go run server.go
```
The server listens on port 8080.
### Run the Client(s)
In separate terminals, run:
```bash
cd client
go run client.go
```
Type messages in the client terminal to send to the server.<br> 
Server messages broadcast to all clients.

---

## 💡 How It Works
- Server accepts incoming connections in a loop

- Each client connection runs in its own goroutine

- Messages from clients are printed on the server terminal

- Server input broadcasts messages to all clients

- sync.Mutex protects shared client map from race conditions

---

## 🙌 Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

