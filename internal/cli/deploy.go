package cli

import (
	"github.com/cardil/k1s/pkg/k1s"
	"github.com/spf13/cobra"
)

type deployCmd struct {
	*Options
}

func (c deployCmd) command() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy",
		Short: "Deploy a Kubernetes cluster",
		RunE:  c.run,
	}
}

func (c deployCmd) run(cmd *cobra.Command, _ []string) error {
	return k1s.Deploy(cmd, c.Stack.String()) //nolint:wrapcheck
}
