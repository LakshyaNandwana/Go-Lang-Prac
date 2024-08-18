package user

import (
	"DI-Wire/domain"
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func (r *repository) FetchByUsername(ctx context.Context, username string) (*domain.UserEntity, error) {
	panic("Implement me please :')") // Hint: xD
}
