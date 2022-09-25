package pulumi

import (
	"fmt"
	"io/fs"

	"github.com/cardil/k1s"
	"golang.org/x/mod/modfile"
)

func version() (string, error) {
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
