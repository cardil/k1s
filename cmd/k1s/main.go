package main

import (
	"github.com/cardil/k1s/internal/cli"
	"github.com/wavesoftware/go-commandline"
)

func main() {
	app := new(cli.App)
	commandline.New(app).ExecuteOrDie(cli.Opts...)
}
