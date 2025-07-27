package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("💬 GoChat - Terminal Based Encrypted Chat")
	fmt.Println("1. Start as Server")
	fmt.Println("2. Start as Client")
	fmt.Print("Choose option: ")

	choice, _ := reader.ReadString('\n')

	switch choice[0] {
	case '1':
		startServer()
	case '2':
		startClient()
	default:
		fmt.Println("Invalid choice")
	}
}