//go:build wireinject
// +build wireinject

package resthttp

import (
	"database/sql"

	"github.com/gadhittana01/socialmedia/db"
	"github.com/gadhittana01/socialmedia/pkg/tag"
	"github.com/gadhittana01/socialmedia/services"
	"github.com/google/wire"
)

var dbTagSet = wire.NewSet(
	db.InitDB,
	wire.Bind(new(tag.DBTX), new(*sql.DB)),
)

var tagPkgSet = wire.NewSet(
	tag.New,
	wire.Bind(new(services.TagResource), new(*tag.Queries)),
)

var tagService = wire.NewSet(
	services.NewTagService,
	wire.Bind(new(TagService), new(services.TagService)),
)

func InitializedTagHandler() (*TagHandler, error) {
	wire.Build(dbTagSet, tagPkgSet, tagService, NewTagHandler)
	return nil, nil
}
