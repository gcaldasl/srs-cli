package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CLI struct{}

func NewCLI() *CLI {
	return &CLI{}
}

func (c *CLI) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n1. Print message")
		fmt.Println("2. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Println("Hello, world!")
		case "2":
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
