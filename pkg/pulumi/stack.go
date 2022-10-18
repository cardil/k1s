package pulumi

import (
	"context"
	"os"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ProjectConfig struct {
	Stack   string
	Project string
	pulumi.RunFunc
}

type Stack interface {
	Up(ctx context.Context, opts ...optup.Option) (auto.UpResult, error)
	CleanUp()
}

func CreateStack(ctx context.Context, pc ProjectConfig) (Stack, error) {
	bin, err := ensurePulumiBinary()
	if err != nil {
		return nil, err
	}
	oldPath := os.Getenv("PATH")
	if err = os.Setenv("PATH", bin+":"+oldPath); err != nil {
		return nil, wrapErr(err, ErrBug)
	}
	var s auto.Stack
	s, err = auto.UpsertStackInlineSource(ctx, pc.Stack, pc.Project, pc.RunFunc)
	if err != nil {
		return nil, wrapErr(err, ErrBug)
	}
	return tempPathStack{oldPath, &s}, nil
}

type tempPathStack struct {
	old string
	*auto.Stack
}

func (t tempPathStack) CleanUp() {
	if err := os.Setenv("PATH", t.old); err != nil {
		panic(err)
	}
}
