//go:build wireinject
// +build wireinject

// you can write your dependancy hear ratherthan create wire_gen.go by using google wire. so you need to maintain only single file also get a good idea.
// In my case i genereate wire_gen.go then , i hard core all depedancy rather than user wire auto create.

package di

import (
	http "sample/pkg/api"
	"sample/pkg/api/handler"
	config "sample/pkg/config"
	"sample/pkg/db"
	"sample/pkg/repository"
	"sample/pkg/usecase"

	"github.com/google/wire"
)

func InitializeAPI(cfg *config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase, repository.NewUserDataBase, usecase.NewuserUseCase, handler.NewUserHandler, repository.NewAdminRepository, usecase.NewAdminUseCase, handler.NewAdminHandler, http.NewServerHttp)

	return &http.ServerHTTP{}, nil
}
