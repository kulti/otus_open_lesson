package router

import (
	"net/http"

	"github.com/kulti/otus_open_lesson/internal/storages"
)

type Router struct {
	rootHandler rootHandler
}

func New(store storages.Store) *Router {
	return &Router{
		rootHandler: newRootHandler(store),
	}
}

func (r *Router) RootHandler() http.Handler {
	return r.rootHandler
}
