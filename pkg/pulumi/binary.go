package pulumi

import (
	"context"
	"os"
	"path"

	"github.com/cardil/ghet/pkg/ghet/download"
	"github.com/cardil/ghet/pkg/ghet/install"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
)

const (
	markerFilePerms      = 0o640
	dirPerm              = 0o750
	downloadedMarkerFile = ".downloaded"
)

func newBinaryOnPathStack(stackFn func() (Stack, error)) (Stack, error) {
	return &binaryOnPathStack{stackFn: stackFn}, nil
}

type binaryOnPathStack struct {
	old     string
	stackFn func() (Stack, error)
}

func (t *binaryOnPathStack) Up(ctx context.Context, opts ...optup.Option) (auto.UpResult, error) {
	if err := t.setBinaryOnPath(ctx); err != nil {
		return auto.UpResult{}, err
	}
	defer t.restorePath()
	s, err := t.stackFn()
	if err != nil {
		return auto.UpResult{}, err
	}
	return s.Up(ctx, opts...) //nolint:wrapcheck
}

func (t *binaryOnPathStack) setBinaryOnPath(ctx context.Context) error {
	bin, err := ensurePulumiBinary(ctx)
	if err != nil {
		return err
	}
	t.old = os.Getenv("PATH")
	if err = os.Setenv("PATH", bin+":"+t.old); err != nil {
		return wrapErr(err, ErrBug)
	}
	return nil
}

func (t *binaryOnPathStack) restorePath() {
	if err := os.Setenv("PATH", t.old); err != nil {
		panic(err)
	}
}

func ensurePulumiBinary(ctx context.Context) (string, error) {
	pver, err := version()
	if err != nil {
		return "", err
	}
	var binDir string
	binDir, err = homeResource("bin")
	if err != nil {
		return "", wrapErr(err, ErrUnexpected)
	}
	if err = os.MkdirAll(binDir, dirPerm); err != nil {
		return "", wrapErr(err, ErrUnexpected)
	}
	if _, err = os.Stat(path.Join(binDir, downloadedMarkerFile)); err == nil {
		return binDir, nil
	}
	args := install.Parse("pulumi/pulumi@" + pver)
	args.MultipleBinaries = true
	if err = download.Action(ctx, download.Args{
		Args: args, Destination: binDir,
	}); err != nil {
		return "", wrapErr(err, ErrUnexpected)
	}
	if err = os.WriteFile(path.Join(binDir, downloadedMarkerFile), []byte{}, markerFilePerms); err != nil {
		return "", wrapErr(err, ErrUnexpected)
	}
	return binDir, nil
}
