package pulumi

import (
	"context"

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
}

func CreateStack(ctx context.Context, pc ProjectConfig) (Stack, error) {
	return createPulumiStack(func() (Stack, error) {
		s, err := auto.UpsertStackInlineSource(ctx, pc.Stack, pc.Project, pc.RunFunc)
		if err != nil {
			return nil, wrapErr(err, ErrBug)
		}
		return &s, nil
	})
}

type stackFn func() (Stack, error)

func createPulumiStack(fn stackFn) (Stack, error) {
	creators := []func(func() (Stack, error)) (Stack, error){
		newBinaryOnPathStack,
		newLoggedInStack,
		newManagedPassphraseStack,
	}

	for i := range creators {
		creator := creators[len(creators)-1-i]
		fn = func(currentFn func() (Stack, error)) func() (Stack, error) {
			return func() (Stack, error) {
				return creator(currentFn)
			}
		}(fn)
	}

	return fn()
}
