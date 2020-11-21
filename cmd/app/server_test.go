package main

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hamba/cmd"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
)

func TestServer_Health(t *testing.T) {
	ts := newTestServer(t)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func newTestContext() (*cmd.Context, *flag.FlagSet) {
	fs := new(flag.FlagSet)
	c := cli.NewContext(&cli.App{}, fs, nil)

	ctx, _ := cmd.NewContext(c)
	ctx.AttachLogger(func(log.Logger) log.Logger { return log.Null })
	ctx.AttachStatter(func(stats.Statter) stats.Statter { return stats.Null })

	return ctx, fs
}

func newTestServer(t *testing.T) *httptest.Server {
	ctx, fs := newTestContext()
	fs.String(flagSomeFlag, "some value", "doc")

	app, err := newApplication(ctx)
	if err != nil {
		t.Fatal(err.Error())
	}

	srv := newServer(app)

	return httptest.NewServer(srv)
}
