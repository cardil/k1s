package cli

import (
	"io"
	"os"

	"github.com/cardil/k1s/pkg/cli/retcode"
	"github.com/cardil/k1s/pkg/stack"
	"github.com/spf13/cobra"
)

// Cmd represents a command line application entrypoint.
type Cmd struct {
	root *cobra.Command
	exit func(code int)
}

// Execute will execute the application.
func (c *Cmd) Execute() {
	if err := c.execute(); err != nil {
		c.exit(retcode.Calc(err))
	}
}

// ExecuteWithOptions will execute the application with the provided options.
func (c *Cmd) ExecuteWithOptions(options ...CommandOption) error {
	return c.execute(options...)
}

// WithArgs creates an option which sets args.
func WithArgs(args ...string) CommandOption {
	return func(command *cobra.Command) {
		command.SetArgs(args)
	}
}

// WithOutput creates an option witch sets os.Stdout and os.Stderr.
func WithOutput(out io.Writer) CommandOption {
	return func(command *cobra.Command) {
		command.SetOut(out)
		command.SetErr(out)
	}
}

// CommandOption is used to configure a command in Cmd.ExecuteWithOptions.
type CommandOption func(*cobra.Command)

func (c *Cmd) execute(configs ...CommandOption) error {
	c.init()
	for _, config := range configs {
		config(c.root)
	}
	// cobra.Command should pass our own errors, no need to wrap them.
	return c.root.Execute() //nolint:wrapcheck
}

func (c *Cmd) init() {
	if c.root != nil {
		return
	}
	c.exit = os.Exit
	c.root = rootCmd(&Options{})
}

// Options holds a general args for all commands.
type Options struct {
	// Verbose tells does commands should display additional information about
	// what's happening? Verbose information is printed on stderr.
	Verbose bool

	stack.Stack

	Out io.Writer
	Err io.Writer
}

type subcommand interface {
	command() *cobra.Command
}
