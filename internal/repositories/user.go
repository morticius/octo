package repositories

import (
	"context"
	"github.com/morticius/octo/internal/models"
)

type IUserRepository interface {
	GetByEmail(context.Context, string) (*models.User, error)
	Save(context.Context, models.User) error
}
