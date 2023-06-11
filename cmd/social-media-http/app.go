package main

import (
	"github.com/gadhittana01/socialmedia/config"
	"github.com/gadhittana01/socialmedia/db"
	"github.com/gadhittana01/socialmedia/handler/resthttp"
	"github.com/gadhittana01/socialmedia/pkg/post"
	"github.com/gadhittana01/socialmedia/pkg/post_tags"
	"github.com/gadhittana01/socialmedia/pkg/tag"
	"github.com/gadhittana01/socialmedia/services"
)

func initApp(c *config.GlobalConfig) error {
	db := db.InitDB()
	postPkg := post.New(db)
	tagPkg := tag.New(db)
	postTagPkg := post_tags.New(db)

	ps, err := services.NewPostService(postPkg, tagPkg, postTagPkg)
	if err != nil {
		return err
	}
	return startHTTPServer(resthttp.NewRoutes(resthttp.RouterDependencies{
		PR: ps,
	}), c)
}
