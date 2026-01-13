package users

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// thats is for DB, writer is for modify data and create
type Writer interface {
	Create(user *User) error
	Upate(id int, attributes map[string]interface{}) error
	Delete(id int) error
}

// only for read and return some data in db
type Reader interface {
	GetByID(id int) (*User, error)
	Auth(email, password string) (*User, error)
}

// for DB operations
type Repository interface {
	// when a aplication is big (empresarial level), need more than one DB, so this is for that case
	//DBs only for readers and only for writers, this can help the DBmaster and DBslave concept
	Writer
	Reader
}

// create a contract for what the user can do
type Usecase interface {
	Create(user *User) error
	GetByID(id int) (*User, error)
	Update(id int, attributes map[string]interface{}) error //interface is the informations that we are passing to validate and update after the validation
	Delete(id int) error
	Auth(email, password string) (*User, error)
}