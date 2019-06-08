package main

import (
	"github.com/hamba/app"
	"github.com/hamba/cmd"
)

// Application =============================

func newApplication(c *cmd.Context) (*app.Application, error) {
	app := app.NewApplication(
		c.Logger(),
		c.Statter(),
	)

	// Setup your application here

	return app, nil
}
