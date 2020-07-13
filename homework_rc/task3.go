package main

import (
	"fmt"
)

var c chan string

func addLine(words string) {
	c <- words
}

func main() {

	c = make(chan string, 4)

	go addLine("I'll")
	go addLine(" be here")
	go addLine(" all day")
	go addLine(" and you'll be too")

	for i:=0; i<4; i++ {
		fmt.Print(<-c)
	}
}
