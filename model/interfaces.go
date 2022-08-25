package model

type UserService interface {
	GetById(id int) (*User, error)
}

type UserRepository interface {
	FindById(id int) (*User, error)
}
