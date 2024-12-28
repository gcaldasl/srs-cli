package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gcaldasl/srs-cli/internal/core/services"
)

type CLI struct {
	service *services.SRSService
}

func NewCLI(service *services.SRSService) *CLI {
	return &CLI{service: service}
}

func (c *CLI) Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n1. Create a new card")
		fmt.Println("2. Review a card")
		fmt.Println("3. List due cards")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			c.createCard(reader)
		case "2":
			c.reviewCard(reader)
		case "3":
			c.listDueCards()
		case "4":
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}

func (c *CLI) createCard(reader *bufio.Reader) {
	fmt.Print("Enter front side of the card: ")
	frontSide, _ := reader.ReadString('\n')
	frontSide = strings.TrimSpace(frontSide)

	fmt.Print("Enter back side of the card: ")
	backSide, _ := reader.ReadString('\n')
	backSide = strings.TrimSpace(backSide)

	err := c.service.CreateCard(frontSide, backSide)
	if err != nil {
		fmt.Println("Error creating card:", err)
	} else {
		fmt.Println("Card created successfully!")
	}
}

func (c *CLI) reviewCard(reader *bufio.Reader) {
	fmt.Print("Enter card ID to review: ")
	idStr, _ := reader.ReadString('\n')
	idStr = strings.TrimSpace(idStr)
	cardID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		fmt.Println("Invalid card ID")
		return
	}

	fmt.Print("Enter review quality (0-5): ")
	qualityStr, _ := reader.ReadString('\n')
	qualityStr = strings.TrimSpace(qualityStr)
	quality, err := strconv.Atoi(qualityStr)
	if err != nil || quality < 0 || quality > 5 {
		fmt.Println("Invalid quality score")
		return
	}

	err = c.service.ReviewCard(cardID, quality)
	if err != nil {
		fmt.Println("Error reviewing card:", err)
	} else {
		fmt.Println("Card reviewed successfully!")
	}
}

func (c *CLI) listDueCards() {
	cards, err := c.service.ListDueCards()
	if err != nil {
		fmt.Println("Error listing due cards:", err)
		return
	}

	if len(cards) == 0 {
		fmt.Println("No cards due for review.")
		return
	}

	fmt.Println("Due cards:")
	for _, card := range cards {
		fmt.Printf("ID: %d, Front: %s, Back: %s\n", card.ID, card.FrontSide, card.BackSide)
	}
}