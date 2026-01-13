package store

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/NFAbricio/example-api/users"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) users.Repository {
	return Repository{db: db}
}

func (r Repository) Create(user *users.User) error {
	return r.db.Create(user).Error
}

func (r Repository) Upate(id int, attributes map[string]interface{}) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to begin transaction: %w", tx.Error)
	}

	defer func() {
		if re := recover(); re != nil { //recover resecute the function in runtime
			tx.Rollback() //if something wrong happen, rollback the transaction, dont commit
		}
	}()

	if err := tx.Model(&users.User{}).Where("id = ?", id).Updates(attributes).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("update school error")
	}

	return tx.Commit().Error
}

func (r Repository) Delete(id int) error {
	return r.db.Delete(&users.User{}, "id = ?", id).Error // id = ? is to avoid SQL injection
}

func (r Repository) GetByID(id int) (*User, error) {}

func (r Repository) Auth(email, password string) (*User, error) {}