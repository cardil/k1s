package pulumi

import (
	"fmt"
	"io/fs"
	"sync"

	"github.com/cardil/k1s"
	"golang.org/x/mod/modfile"
)

type versionCache struct {
	version string
	err     error
	once    sync.Once
}

var vc versionCache //nolint:gochecknoglobals

func version() (string, error) {
	vc.once.Do(func() {
		vc.version, vc.err = resolveVersion()
	})
	return vc.version, vc.err
}

func resolveVersion() (string, error) {
	gomodFilename := "go.mod"
	data, err := fs.ReadFile(k1s.GoMod, gomodFilename)
	if err != nil {
		return "", wrapErr(err, ErrBug)
	}
	mod, err := modfile.ParseLax(gomodFilename, data, nil)
	if err != nil {
		return "", wrapErr(err, ErrBug)
	}
	for _, dep := range mod.Require {
		if dep.Mod.Path == "github.com/pulumi/pulumi/sdk/v3" {
			return dep.Mod.Version, nil
		}
	}
	return "", fmt.Errorf("%w: can't find Pulumi SDK version", ErrBug)
}
