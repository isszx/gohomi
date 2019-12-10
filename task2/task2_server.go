package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Running server on port :8081")
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("Start TCP error:", err)
	}
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal("TCP Accept error:", err)
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	message, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		log.Println("Сlient left TCP channel")
		connection.Close()
		return
	}

	fmt.Print("Server got message: ", string(message))
	clearMessage := strings.TrimRight(string(message), "\n")
	res, err := strconv.Atoi(clearMessage)
	if err == nil {
		connection.Write([]byte(strconv.Itoa(res*2) + "\n"))
	} else {
		msg := strings.ToUpper(message)
		connection.Write([]byte(msg + "\n"))
	}
	handleConnection(connection)
}
