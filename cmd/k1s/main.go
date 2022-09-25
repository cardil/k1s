package main

import (
	"github.com/cardil/k1s/internal/cli"
	"github.com/wavesoftware/go-commandline"
)

func main() {
	commandline.New(&cli.App{}).ExecuteOrDie(cli.Opts...)
}

// RunMain is used for testing.
func RunMain() { //nolint:deadcode
	main()
}
