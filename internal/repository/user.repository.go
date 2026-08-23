package repository

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetUsersRepo() string {
	return "Hanoidev"
}
