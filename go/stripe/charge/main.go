package main

import (
	"fmt"
	"log"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func main() {
	params := &stripe.ChargeParams{
		Amount:      stripe.Int64(999),
		Currency:    stripe.String(string(stripe.CurrencyJPY)),
		Description: stripe.String("test"),
	}

	token := "tok_mastercard"

	if err := params.SetSource(token); err != nil {
		log.Fatal(err)
	}

	c, err := charge.New(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c.ID)
}

func init() {
	stripe.Key = os.Getenv("STRIPE_TEST_KEY")
}
