package cli

import (
	"github.com/cardil/k1s/pkg/k1s"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/spf13/cobra"
)

type deployCmd struct {
	*Options
	root *cobra.Command
}

func (c deployCmd) command() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy",
		Short: "Deploy a Kubernetes cluster",
		RunE:  c.run,
	}
}

func (c deployCmd) run(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()
	s, err := auto.UpsertStackInlineSource(ctx, c.Stack.String(), "k1s", k1s.Up)
	if err != nil {
		return err
	}
	_, err = s.Up(ctx, optup.ProgressStreams(cmd.OutOrStdout()))
	return err
}
