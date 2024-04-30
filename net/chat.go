package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string //An outgoing message channel

var (
	incomingClients = make(chan Client) //All incoming client messages
	leavingClients  = make(chan Client) //All leaving client messages
	messages        = make(chan string) //All client messages
)
var (
	host = flag.String("host", "localhost", "Host to connect to")
	port = flag.Int("port", 8080, "Port to connect to")
)

func HandleConnection(conn net.Conn) { //handles the connection to a specific client
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)
	message := make(chan string) //Outgoing client messages
	go MessageWrite(conn, message)

	clientName := conn.RemoteAddr().String()
	message <- "You are " + clientName + " Welcome to the server\n"
	messages <- clientName + " has arrived"

	incomingClients <- message
	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- clientName + ": " + inputMessage.Text()
	}
	leavingClients <- message
	messages <- clientName + " has left"

}

func MessageWrite(conn net.Conn, messages <-chan string) {
	for msg := range messages {
		_, err := fmt.Fprintln(conn, msg)
		if err != nil {
			return
		}
	}

}

func Broadcaster() {
	clients := make(map[Client]bool)
	for {
		select {
		case msg := <-messages: // A message is sent
			for cli := range clients {
				cli <- msg //the message will be sent to all clients
			}
		case cli := <-incomingClients: // When a new client arrives
			clients[cli] = true
		case cli := <-leavingClients: // When a client leaves
			delete(clients, cli)
			close(cli)
		}
	}

}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatalln(err)
	}
	go Broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go HandleConnection(conn)
	}

}
