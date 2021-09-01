package repositories

import (
	"context"
	"github.com/morticius/octo/internal/models"
)

type IUserRepository interface {
	GetByEmail(context.Context, string) (*models.User, error)
	Save(context.Context, models.User) error
}

type MockUserRepository struct{}

func (MockUserRepository) GetByEmail(ctx context.Context, s string) (*models.User, error) {
	panic("implement me")
}

func (MockUserRepository) Save(ctx context.Context, user models.User) error {
	panic("implement me")
}
