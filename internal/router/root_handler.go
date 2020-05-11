package router

import (
	"net/http"

	"github.com/kulti/otus_open_lesson/internal/storages"
)

type rootHandler struct {
	tasksHandler tasksHandler
	taskHandler  taskHandler
}

func newRootHandler(store storages.Store) rootHandler {
	return rootHandler{
		tasksHandler: newTasksHandler(store),
		taskHandler:  newTaskHandler(store),
	}
}

func (h rootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var head string
	head, r.URL.Path = shiftPath(r.URL.Path)

	switch head {
	case "tasks":
		h.tasksHandler.ServeHTTP(w, r)
	case "task":
		h.taskHandler.ServeHTTP(w, r)
	default:
		http.NotFound(w, r)
	}
}
