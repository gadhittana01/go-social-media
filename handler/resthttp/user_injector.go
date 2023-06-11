//go:build wireinject
// +build wireinject

package resthttp

import (
	"database/sql"

	"github.com/gadhittana01/socialmedia/db"
	"github.com/gadhittana01/socialmedia/pkg/user"
	"github.com/gadhittana01/socialmedia/services"
	"github.com/google/wire"
)

var dbUsrSet = wire.NewSet(
	db.InitDB,
	wire.Bind(new(user.DBTX), new(*sql.DB)),
)

var userPkgSet = wire.NewSet(
	user.New,
	wire.Bind(new(services.UserResource), new(*user.Queries)),
)

var userService = wire.NewSet(
	services.NewUserService,
	wire.Bind(new(UserService), new(services.UserService)),
)

func InitializedUserHandler() (*UserHandler, error) {
	wire.Build(dbUsrSet, userPkgSet, userService, NewUserHandler)
	return nil, nil
}
