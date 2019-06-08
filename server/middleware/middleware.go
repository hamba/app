package middleware

import (
	"net/http"

	"github.com/hamba/pkg/httpx/middleware"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
)

// Monitorable represents a Loggable, Statable object.
type Monitorable interface {
	log.Loggable
	stats.Statable
}

// Common wraps the handler with common middleware.
func Common(h http.Handler, m Monitorable) http.Handler {
	h = middleware.WithRequestStats(h, m)
	return middleware.WithRecovery(h, m)
}
