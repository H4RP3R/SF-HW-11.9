package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
)

func send(buffer *ring.Ring) {
	for {
		if buffer.Value == nil {
			break
		}
		fmt.Printf("%s ", buffer.Value)
		buffer.Value = nil
		buffer = buffer.Next()
	}
	fmt.Println()
}

func main() {
	buffer := ring.New(1024)
	fmt.Println("Enter a message. You can push enter, but only 'QRT' will finish the message and send it.")

	reader := bufio.NewReader(os.Stdin)
	start := buffer

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		input = input[:len(input)-1]

		if input == "QRT" {
			send(start)
			buffer = start
		} else {
			buffer.Value = input
			buffer = buffer.Next()
		}
	}
}
