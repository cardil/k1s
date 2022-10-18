package pulumi

import (
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
)

const dirPerm = 0o750

func ensurePulumiBinary() (string, error) {
	pver, err := version()
	if err != nil {
		return "", err
	}
	var binDir string
	binDir, err = homedir.Expand(path.Join("~", ".cache", "pulumi", pver, "bin"))
	if err != nil {
		return "", wrapErr(err, ErrUnexpected)
	}
	err = os.MkdirAll(binDir, dirPerm)
	if err != nil {
		return "", wrapErr(err, ErrUnexpected)
	}
	return binDir, nil
}
