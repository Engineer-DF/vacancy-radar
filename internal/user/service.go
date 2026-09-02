package user

import "context"

// Интерфейс зависимости. Use case сам диктует, какие данные ему нужны.
// Реализация этого интерфейса будет находиться в слое инфраструктуры/БД (repository)
// Хотя хуй его знает....

type UserRepository interface {
	Create(ctx context.Context, user *User) error
}

type NewUserRepository struct {
	repo UserRepository
}
