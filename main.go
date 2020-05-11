package main

import (
	"net/http"

	"github.com/kulti/otus_open_lesson/internal/router"
	"github.com/kulti/otus_open_lesson/internal/storages/memstore"
)

func main() {
	r := router.New(memstore.New())
	http.ListenAndServe(":8080", r.RootHandler())
}
