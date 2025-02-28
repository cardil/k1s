package cli

import "github.com/spf13/cobra"

type purgeCmd struct {
	*Options
}

func (c purgeCmd) command() *cobra.Command {
	return &cobra.Command{
		Use:   "purge",
		Short: "Purge a Kubernetes cluster",
		RunE:  c.run,
	}
}

func (c purgeCmd) run(_ *cobra.Command, _ []string) error {
	return nil
}
