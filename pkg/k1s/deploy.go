package k1s

import (
	"github.com/cardil/k1s/pkg/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
)

func Deploy(core App, stack string) error {
	ctx := core.Context()
	s, err := pulumi.CreateStack(ctx, pulumi.ProjectConfig{
		Stack:   stack,
		Project: "k1s",
		RunFunc: Project,
	})
	if err != nil {
		return wrapErr(err, ErrInvalidCode)
	}
	_, err = s.Up(ctx, optup.ProgressStreams(core.OutOrStdout()))
	return wrapErr(err, ErrUnexpected)
}
