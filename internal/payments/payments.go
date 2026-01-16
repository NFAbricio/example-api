package payments

import "github.com/stripe/stripe-go/v79"

type Payment interface {
	CreateCustomer(name, email, phone string) (*stripe.Customer, error)
	DeleteCustomer(customerID string) (*stripe.Customer, error) 
}