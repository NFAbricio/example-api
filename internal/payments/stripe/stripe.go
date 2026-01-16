package stripe

import (
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/customer"

	"github.com/NFAbricio/example-api/internal/payments"
)

type PaymentService struct {
	stripekey string
}

func NewPaymentService(sk string) payments.Payment {
	return &PaymentService{
		stripekey: sk,
	}
}

func (ps *PaymentService) CreateCustomer(name, email, phone string) (*stripe.Customer, error) {
	stripe.Key = ps.stripekey

	params := &stripe.CustomerParams{
		Name: &name,
		Email: &email,
		Phone: &phone,
	}

	return customer.New(params)
}

func (ps *PaymentService) DeleteCustomer(customerID string) (*stripe.Customer, error) {
	stripe.Key = ps.stripekey

	return customer.Del(customerID, nil)
}