package user

import (
	"DI-Wire/domain"
	"context"
)



type service struct{
	repo domain.UserRepository
}

func (s *service) FetchByUsername(ctx context.Context, username string) (*domain.User, error){
	panic("IMPLEMENT MEE!")
}