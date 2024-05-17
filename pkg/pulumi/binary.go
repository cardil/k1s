package pulumi

import (
	"context"
	"os"
	"path"

	"github.com/cardil/ghet/pkg/ghet/download"
	"github.com/cardil/ghet/pkg/ghet/install"
	"github.com/mitchellh/go-homedir"
)

const (
	markerFilePerms      = 0o640
	dirPerm              = 0o750
	downloadedMarkerFile = ".downloaded"
)

func ensurePulumiBinary(ctx context.Context) (string, error) {
	pver, err := version()
	if err != nil {
		return "", err
	}
	var binDir string
	binDir, err = homedir.Expand(path.Join("~", ".cache", "pulumi", pver, "bin"))
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
