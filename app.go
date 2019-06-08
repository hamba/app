package app

import (
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
)

// Application represents the application.
type Application struct {
	logger  log.Logger
	statter stats.Statter
}

// NewApplication creates an instance of Application.
func NewApplication(l log.Logger, s stats.Statter) *Application {
	return &Application{
		logger:  l,
		statter: s,
	}
}

// YourFunction represents your application function.
func (a *Application) YourFunction() error {
	return nil
}

// IsHealthy checks the health of the Application.
func (a *Application) IsHealthy() error {
	return nil
}

// Logger returns the Logger attached to the Application.
func (a *Application) Logger() log.Logger {
	return a.logger
}

// Statter returns the Statter attached to the Application.
func (a *Application) Statter() stats.Statter {
	return a.statter
}
