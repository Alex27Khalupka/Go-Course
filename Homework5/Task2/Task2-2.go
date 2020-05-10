package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	/*
	TCP Client connects to a TCP server on a port.
	Client reads STDIN for any input.
	On hit enter Client sends input it got to a server.
	Client shows response from S to STDOUT and waits for another input from STDIN.
	 */

	port := 8081
	fmt.Printf("Connecting to server %d\n", port)
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))

	if err != nil{
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for{
		var message string
		fmt.Println("Waiting for input ")

		// reading STDIN for any input
		_, err := fmt.Scanln(&message)
		if err != nil{
			fmt.Println("Incorrect input", err)
			continue
		}
		// checking client close condition
		if message == "exit"{
			fmt.Println("Client closed")
			os.Exit(0)
		}

		// sending message to the server
		if _, err = conn.Write([]byte(message + "\n")); err!=nil {
			log.Fatal(err)
		}

		fmt.Println("Answer from server: ")

		// reading answer from the server
		answer, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {break}
		fmt.Println(answer)

	}
}
