package server

import (
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/hamba/pkg/httpx"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
)

// Application represents the main application.
type Application interface {
	log.Loggable
	stats.Statable

	// YourFunction represents your application function.
	YourFunction() error
}

// NewMux creates a new server mux.
func NewMux(app Application) *bone.Mux {
	mux := httpx.NewMux()

	mux.GetFunc("/", SomeHandler(app))

	return mux
}

// SomeHandler handles requests of some type.
func SomeHandler(app Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// You can call your application functions from here
	}
}
