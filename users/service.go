package users

// why not diretely a store?
// if has a change of db, we need to change only the repository that is implementing the interface Repository
type Service struct {
	repository Repository
}

func NewService(repository Repository) Usecase {
	return &Service{repository: repository}
}

func (s *Service) Create(user *User) error {
	return nil
}

func (s *Service) GetByID(id int) (*User, error) {
	return nil, nil
}

func (s *Service) Update(id int, attributes map[string]interface{}) error {
	return nil
}

func (s *Service) Delete(id int) error {
	return nil
}

func (s *Service) Auth(email, password string) (*User, error) {
	return nil, nil
}