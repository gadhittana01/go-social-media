//go:build wireinject
// +build wireinject

package resthttp

import "github.com/google/wire"

func InitializedPostHandler(ps PostService) (*PostHandler, error) {
	wire.Build(NewPostHandler)
	return nil, nil
}
