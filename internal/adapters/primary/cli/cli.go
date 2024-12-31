package cli

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"

	"github.com/gcaldasl/srs-cli/internal/core/services"
)

var (
	ErrInvalidInput      = errors.New("Invalid input provided.")
	ErrEmptyInput        = errors.New("Input cannot be empty.")
	ErrInvalidQuality    = errors.New("Quality score must be between 0 and 5.")
	ErrInvalidCardID     = errors.New("Invalid card ID.")
	ErrOperationCanceled = errors.New("Operation canceled by user.")
)

type CLI struct {
	service *services.CardService
}

func NewCLI(service *services.CardService) *CLI {
	if service == nil {
		panic("Card service can't be nil")
	}
	return &CLI{service: service}
}

func (c *CLI) Run() {
	for {
		fmt.Println("\n=== SRS CLI Application ===")
		prompt := promptui.Select{
			Label: "Select an option",
			Items: []string{
				"Review",
				"Create Deck",
				"Create Card",
				"Add Card to Deck",
				"Remove Card From Deck",
				"Delete Deck",
				"Exit",
			},
		}

		_, result, err := prompt.Run()
		if err != nil {
			fmt.Println("Error displaying menu:", err)
			continue
		}

		if err := c.handleMenuChoice(result); err != nil {
			if errors.Is(err, ErrOperationCanceled) {
				fmt.Println("Operation canceled by user.")
				continue
			}

			fmt.Printf("Error: %v\n", err)
		}

		if result == "Exit" {
			fmt.Println("Goodbye!")
			return
		}
	}
}

func (c *CLI) handleMenuChoice(choice string) error {
	switch choice {
	case "Review":
		return nil
	case "Create Deck":
		return nil
	case "Create Card":
		return c.handleCreateCard()
	case "Add Card to Deck":
		return nil
	case "Remove Card from Deck":
		return nil
	case "Delete Deck":
		return nil
	case "Exit":
		return nil
	default:
		return ErrInvalidInput
	}
}

func (c *CLI) handleCreateCard() error {
	frontSide, err := c.promptForInput("Enter front side of the card", true)
	if err != nil {
		return fmt.Errorf("Error getting front side of the card: %w", err)
	}

	backSide, err := c.promptForInput("Enter back side of the card", true)
	if err != nil {
		return fmt.Errorf("Error getting back side of the card: %w", err)
	}

	confirm := promptui.Prompt{
		Label:   "Create card with these details? (y/n)",
		Default: "y",
		Validate: func(input string) error {
			input = strings.ToLower(input)
			if input != "y" && input != "n" {
				return errors.New("Please enter 'y' or 'n'")
			}

			return nil
		},
	}

	result, err := confirm.Run()
	if err != nil {
		return fmt.Errorf("Error confirming card creation: %w", err)
	}

	if strings.ToLower(result) != "y" {
		return ErrOperationCanceled
	}

	if err := c.service.CreateCard(frontSide, backSide); err != nil {
		return fmt.Errorf("Error creating card: %w", err)
	}

	fmt.Println("Card created successfully!")
	return nil
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
		fmt.Printf("quality %d", quality)
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

func (c *CLI) promptForInput(label string, required bool) (string, error) {
	prompt := promptui.Prompt{
		Label: label,
		Validate: func(input string) error {
			if required && strings.TrimSpace(input) == "" {
				return ErrEmptyInput
			}
			return nil
		},
	}

	result, err := prompt.Run()
	if err != nil {
		if err == promptui.ErrInterrupt {
			return "", ErrOperationCanceled
		}
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return strings.TrimSpace(result), nil
}
