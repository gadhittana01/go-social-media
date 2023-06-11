package resthttp

import (
	"log"

	"github.com/go-chi/chi"
)

type RouterDependencies struct {
	PR PostService
}

func NewRoutes(rd RouterDependencies) *chi.Mux {
	router := chi.NewRouter()

	uh, err := InitializedUserHandler()
	if err != nil {
		log.Println(err)
	}

	th, err := InitializedTagHandler()
	if err != nil {
		log.Println(err)
	}

	ph, err := InitializedPostHandler(rd.PR)
	if err != nil {
		log.Println(err)
	}

	// user
	router.Get("/users", uh.GetUsers)
	router.Post("/user", uh.CreateUser)
	router.Put("/user", uh.UpdateUser)
	router.Delete("/user", uh.DeleteUser)

	// tag
	router.Get("/tags", th.GetTags)
	router.Post("/tag", th.CreateTag)
	router.Put("/tag", th.UpdateTag)
	router.Delete("/tag", th.DeleteTag)

	// post
	router.Get("/posts", ph.GetPosts)
	router.Post("/post", ph.CreatePost)
	router.Put("/post", ph.UpdatePost)
	router.Delete("/post", ph.DeletePost)

	return router
}
