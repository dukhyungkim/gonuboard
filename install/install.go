package install

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func DefaultRouter(r chi.Router) {
	r.Get("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("install index"))
}
