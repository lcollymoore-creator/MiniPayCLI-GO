package ui

import (
	"fmt"

	payment "github.com/lcollymoore-creator/MiniPayCLI-GO/internal/service"
)

func Menu() {
menu:
	for {
		var option int

		fmt.Println("\n1. Payment with Card")
		fmt.Println("2. Payment with Debit Card")
		fmt.Println("3. Payment with PayPal")
		fmt.Println("4. Exit")
		fmt.Print("Select option: ")

		n, err := fmt.Scanln(&option)

		if err != nil {
			fmt.Println("Error reading option:", err)
			continue
		}

		fmt.Printf("Option entered: %d (values read: %d)\n", option, n)

		switch option {
		case 1:
			payment.PaymentWithCreditCard(20)
		case 2:
			payment.PaymentWithDebitCard(10)
		case 3:
			payment.PaymentWithPayPal(14)
		case 4:
			fmt.Println("Exiting...")
			break menu
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}
