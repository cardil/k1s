package cli

import (
	"github.com/cardil/k1s/pkg/stack"
	"github.com/spf13/cobra"
	"github.com/wavesoftware/go-commandline"
)

var Opts []commandline.Option //nolint:gochecknoglobals

// Options holds a general args for all commands.
type Options struct {
	// Verbose tells does commands should display additional information about
	// what's happening? Verbose information is printed on stderr.
	Verbose bool

	stack.Stack
}

type subcommand interface {
	command() *cobra.Command
}
