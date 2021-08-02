package repositories

import (
	"context"
	"github.com/morticius/accuracy/internal/models"
)

type IUserRepository interface {
	GetByEmail(context.Context, string) (*models.User, error)
	Save(context.Context, models.User) error
}
