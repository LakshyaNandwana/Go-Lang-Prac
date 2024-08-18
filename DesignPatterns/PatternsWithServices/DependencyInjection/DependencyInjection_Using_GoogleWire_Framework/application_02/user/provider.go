package user

import (
	"DI-Wire/domain"
	"database/sql"
	"sync"

	"github.com/google/wire"
)

// Simple Implementation

/*
func ProvideRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}
*/

//Using singleton object as provider.

var (
	hdl     *handler
	hdlOnce sync.Once

	svc     *service
	svcOnce sync.Once

	repo     *repository
	repoOnce sync.Once
)

func ProvideHandler(svc domain.UserService) *handler {
	hdlOnce.Do(func() {
		hdl = &handler{
			svc: svc,
		}
	})

	return hdl
}

func ProvideService(repo domain.UserRepository) *service {
	svcOnce.Do(func() {
		svc = &service{
			repo: repo,
		}
	})

	return svc
}

func ProvideRepository(db *sql.DB) *repository {
	repoOnce.Do(func() {
		repo = &repository{
			db: db,
		}
	})

	return repo
}

var ProviderSet wire.ProviderSet = wire.NewSet(
	ProvideHandler,
	ProvideService,
	ProvideRepository,

	wire.Bind(new(domain.UserHandler), new(*handler)),
	wire.Bind(new(domain.UserService), new(*service)),
	wire.Bind(new(domain.UserRepository), new(*repository)),
)
