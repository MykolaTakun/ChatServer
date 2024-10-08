package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
)

func read2(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	for {
		reader := bufio.NewReader(conn)
		msg, _ := reader.ReadString('\n')
		fmt.Printf(msg)
	}
}

func write2(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	//TODO Try to connect to the server
	conn, _ := net.Dial("tcp", *addrPtr)
	//TODO Start asynchronously reading and displaying messages
	go read2(conn)
	//TODO Start getting and sending user messages.
	go write2(conn)
	for {
	}
}
