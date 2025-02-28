package pulumi

import (
	"context"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
)

func Setup(ctx context.Context) error {
	st, err := createPulumiStack(func() (Stack, error) {
		return &setupStack{}, nil
	})
	if err != nil {
		return err
	}
	_, err = st.Up(ctx)
	return err
}

type setupStack struct{}

func (s *setupStack) Up(context.Context, ...optup.Option) (auto.UpResult, error) {
	return auto.UpResult{}, nil
}
