package main

import (
	"fmt"
	"net/http"

	"github.com/hamba/app"
	"github.com/hamba/app/server"
	"github.com/hamba/app/server/middleware"
	"github.com/hamba/cmd"
	"github.com/hamba/pkg/httpx"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
	"gopkg.in/urfave/cli.v2"
)

func runServer(c *cli.Context) error {
	ctx, err := cmd.NewContext(c)
	if err != nil {
		return err
	}

	app, err := newApplication(ctx)
	if err != nil {
		log.Fatal(ctx, err.Error())
	}

	port := c.String(cmd.FlagPort)
	s := newServer(ctx, app)
	log.Info(ctx, fmt.Sprintf("Starting server on port %s", port))
	if err := http.ListenAndServe(":"+port, s); err != nil {
		log.Fatal(ctx, "app: server error", "error", err.Error())
	}

	return nil
}

func newServer(sable stats.Statable, app *app.Application) http.Handler {
	health := httpx.NewHealthMux(app)
	met := httpx.NewStatsMux(sable.Statter())
	srv := server.NewMux(app)
	mux := httpx.CombineMuxes(health, met, srv)

	return middleware.Common(mux, app)
}
