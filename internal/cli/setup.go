package cli

import (
	"github.com/cardil/k1s/pkg/pulumi"
	"github.com/spf13/cobra"
)

type setupCmd struct{}

func (c setupCmd) command() *cobra.Command {
	return &cobra.Command{
		Use:    "setup",
		Short:  "Setup the k1s environment",
		Hidden: true,
		RunE:   c.run,
	}
}

func (c setupCmd) run(cmd *cobra.Command, _ []string) error {
	return pulumi.Setup(cmd.Context()) //nolint:wrapcheck
}
