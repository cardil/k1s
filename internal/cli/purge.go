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

func (c purgeCmd) run(cmd *cobra.Command, args []string) error {
	return nil
}
