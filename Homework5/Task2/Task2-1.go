package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
)

func main() {
	/*
		TCP Server listens for a connections on a port.
		Reads input (split by ‘\n’) and checks if it’s an int,
		returns input multiplied by 2. If it’s not an integer,
		return uppercased input string.
	*/

	port := 8081
	fmt.Printf("Launching server on port: %d \n", port)
	ln, _ := net.Listen("tcp", fmt.Sprintf(":%d", port))
	conn, _ := ln.Accept()
	fmt.Println("Accepted connection")
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')

		// if client disconnected - waiting for another client to connect
		if err != nil {
			if err == io.EOF {
				fmt.Println("Waiting for connection")
				conn, err = ln.Accept()
				fmt.Println("Accepted connection")
				continue
			}
			fmt.Println(err)
			return
		}

		fmt.Print("Message Received: ", message)

		// trying to convert message to int.
		// If successful - multiply by 2.
		// Else - uppercase the sting.
		if intMessage, err := strconv.Atoi(message[:len(message) - 1]); err==nil{
			answer := intMessage * 2
			// sending answer to client
			if _, err = conn.Write([]byte(strconv.Itoa(answer) + "\n")); err!=nil{
				log.Fatal(err)
			}
			fmt.Printf("Answer sent: %d\n\n", answer)
		} else {
			answer := strings.ToUpper(message)
			// sending answer to client
			if _, err = conn.Write([]byte(answer + "\n")); err !=nil{
				log.Fatal(err)
			}
			fmt.Println("Answer sent: ", answer)
		}
	}
}
