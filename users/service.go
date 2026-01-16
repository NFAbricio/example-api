package users

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/NFAbricio/example-api/internal/payments"
)

// why not diretely a store?
// if has a change of db, we need to change only the repository that is implementing the interface Repository
type Service struct {
	repository Repository
	payments   payments.Payment
}

func NewService(repository Repository) Usecase {
	return &Service{repository: repository}
}

func (s *Service) Create(user *User) error {
	
	customer, err := s.payments.CreateCustomer(user.Name, user.Email, user.CustomerID)
	if err != nil {
		return fmt.Errorf("fail to create customer in payment plataform: %w", err)
	}

	user.CustomerID = customer.ID

	err = s.repository.Create(user)
	if err != nil {
		return fmt.Errorf("fail to create user: %w", err)	
	}
	
	return nil
}

func (s *Service) GetByID(id int) (*User, error) {
	user, err := s.repository.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound){
		return nil, fmt.Errorf("user not found: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("fail to get user: %w", err)
	}
	
	return user, nil
}

func (s *Service) Update(id int, attributes map[string]interface{}) error {
	return nil
}

func (s *Service) Delete(id int) error {
	user, err := s.repository.GetByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound){
		return fmt.Errorf("user not found: %w", err)
	}
	if err != nil {
		return fmt.Errorf("fail to get user: %w", err)
	}

	err = s.repository.Delete(id)
	if errors.Is(err, gorm.ErrRecordNotFound){
		return fmt.Errorf("user not found: %w", err)
	}
	if err != nil {
		return fmt.Errorf("fail to delete user: %w", err)
	}
	
	if user.CustomerID == "" {
		return nil
	}

	_, err = s.payments.DeleteCustomer(user.CustomerID)
	if err != nil {
		return err
	}

	
	return nil
}

func (s *Service) Auth(email, password string) (*User, error) {
	user, err := s.repository.Auth(email, password)
	if errors.Is(err, gorm.ErrRecordNotFound){
		return nil, fmt.Errorf("user not found: %w", err)
	}
	if err != nil {
		return nil, fmt.Errorf("fail to get user: %w", err)
	}
	
	return nil, nil
}