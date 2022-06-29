package cli

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cardil/k1s/pkg/stack"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"
)

func rootCmd(opts *Options) *cobra.Command {
	r := &cobra.Command{
		Use:   "k1s",
		Short: "Kubernetes as an API using k3s",
	}
	r.SetOut(coalesceWriter(opts.Out, os.Stdout))
	r.SetErr(coalesceWriter(opts.Err, os.Stderr))
	r.PersistentFlags().BoolVarP(
		&opts.Verbose, "verbose", "v",
		false, "verbose output",
	)
	r.PersistentFlags().VarP(
		enumflag.New(&opts.Stack, "stack", stack.Mapping(), enumflag.EnumCaseInsensitive),
		"stack", "s",
		fmt.Sprintf("Stack. One of: %s.", strings.Join(stacks(), "|")),
	)
	sc := []subcommand{
		&deployCmd{opts, r},
		&purgeCmd{opts, r},
	}
	for _, s := range sc {
		r.AddCommand(s.command())
	}
	return r
}

func coalesceWriter(writters ...io.Writer) io.Writer {
	for _, w := range writters {
		if w != nil {
			return w
		}
	}
	return nil
}

func stacks() []string {
	var ss []string
	for s := range stack.Mapping() {
		ss = append(ss, s.String())
	}
	return ss
}
