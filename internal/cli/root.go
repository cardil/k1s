package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/cardil/k1s/pkg/stack"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"
)

type App struct {
	Options
}

func (a *App) Command() *cobra.Command {
	r := &cobra.Command{
		Use:          "k1s",
		Short:        "Kubernetes as an API using k3s",
		SilenceUsage: true,
	}
	a.flags(r)
	sc := []subcommand{
		&deployCmd{&a.Options},
		&purgeCmd{&a.Options},
		&setupCmd{},
	}
	for _, s := range sc {
		r.AddCommand(s.command())
	}
	r.SetOut(os.Stdout)
	return r
}

func (a *App) flags(r *cobra.Command) {
	r.PersistentFlags().BoolVarP(
		&a.Verbose, "verbose", "v",
		false, "verbose output",
	)
	r.PersistentFlags().VarP(
		enumflag.New(&a.Stack, "stack", stack.Mapping(), enumflag.EnumCaseInsensitive),
		"stack", "s",
		fmt.Sprintf("Stack. One of: %s.", strings.Join(stacks(), "|")),
	)
}

func stacks() []string {
	m := stack.Mapping()
	ss := make([]string, 0, len(m))
	for s := range m {
		ss = append(ss, s.String())
	}
	return ss
}
