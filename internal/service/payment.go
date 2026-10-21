package payment

import (
	"errors"
	"fmt"
)

var CreditCardBalance float64 = 2000
var DebitCardBalance float64 = 5000
var PaypalBalance float64 = 1000

func PaymentWithCreditCard(money float64) error {

	fmt.Println("Account Credit Balance", CreditCardBalance)
	fmt.Print("Enter payment amount: ")
	if money <= CreditCardBalance {
		fmt.Println("Payment with Credit Card successful", money)

		CreditCardBalance = CreditCardBalance - money

		fmt.Println("Current balance:", CreditCardBalance)
	} else {
		fmt.Println("Insufficient funds")
	}
	if money == 0 {
		return errors.New("Payment Fail")

	}
	if money <= 0 {
		return errors.New("Payment Fail")

	}
	return nil

}

func PaymentWithDebitCard(money float64) error {
	fmt.Println("Account Debit Balance", DebitCardBalance)
	fmt.Print("Enter payment amount: ")
	if money <= DebitCardBalance {
		fmt.Println("Payment with Credit Card successful", money)

		DebitCardBalance = DebitCardBalance - money

		fmt.Println("Current balance:", DebitCardBalance)
	} else {
		fmt.Println("Insufficient funds")
	}
	if money == 0 {
		return errors.New("Payment Fail")

	}
	if money <= 0 {
		return errors.New("Payment Fail")

	}
	return nil

}

func PaymentWithPayPal(money float64) error {
	fmt.Println("Account Paypal Balance", PaypalBalance)
	if money <= PaypalBalance {
		fmt.Println("Payment with Credit Card successful", money)

		PaypalBalance = PaypalBalance - money

		fmt.Println("Current balance:", PaypalBalance)
	} else {
		fmt.Println("Insufficient funds")
	}
	if money == 0 {
		return errors.New("Payment Fail")

	}
	if money <= 0 {
		return errors.New("Payment Fail")

	}
	return nil

}
