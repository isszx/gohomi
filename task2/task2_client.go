package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Println(err)
		return
	}
	for {
		readData := bufio.NewReader(os.Stdin)
		fmt.Print("Input message: ")
		message, err := readData.ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		if message == "exit\n" {
			connection.Close()
			return
		}
		fmt.Fprintf(connection, message+"\n")
		msg, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			log.Println(err)
		}
		fmt.Print("Received from server: " + msg)
	}
}
