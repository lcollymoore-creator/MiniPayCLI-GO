package ui

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	payment "github.com/lcollymoore-creator/MiniPayCLI-GO/internal/service"
)

type Transaction struct {
	Method string
	Amount float64
}

func Menu() {
	options := []string{
		"Payment with Card",
		"Payment with Debit Card",
		"Payment with PayPal",
		"Transaction History",
		"Exit",
	}

	var history []Transaction

	for {
		var selected string

		prompt := &survey.Select{
			Message: "Select an option:",
			Options: options,
		}

		err := survey.AskOne(prompt, &selected)
		if err != nil {
			fmt.Println("Error selecting option:", err)
			return
		}

		if selected == "Exit" {
			fmt.Println("Exiting...")
			return
		}

		// Show transaction history
		if selected == "Transaction History" {
			fmt.Println("\n--- Transaction History ---")

			if len(history) == 0 {
				fmt.Println("No transactions yet.")
			} else {
				for i, t := range history {
					fmt.Printf("%d. %s: $%.2f\n",
						i+1, t.Method, t.Amount)
				}
			}

			continue
		}

		// Read payment amount
		var money float64

		fmt.Print("Enter the amount: ")
		_, err = fmt.Scanln(&money)
		if err != nil || money <= 0 {
			fmt.Println("Invalid amount.")
			continue
		}

		// Process payment and save transaction
		switch selected {
		case "Payment with Card":
			payment.PaymentWithCreditCard(money)
			history = append(history, Transaction{
				Method: "Credit Card",
				Amount: money,
			})

		case "Payment with Debit Card":
			payment.PaymentWithDebitCard(money)
			history = append(history, Transaction{
				Method: "Debit Card",
				Amount: money,
			})

		case "Payment with PayPal":
			payment.PaymentWithPayPal(money)
			history = append(history, Transaction{
				Method: "PayPal",
				Amount: money,
			})

		}
	}
}
