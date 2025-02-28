package pulumi

import (
	"path"

	"github.com/mitchellh/go-homedir"
)

func homeResource(subdir ...string) (string, error) {
	var pver string
	{
		v, err := version()
		if err != nil {
			return "", err
		}
		pver = v
	}
	homedirPaths := []string{"~", ".cache", "k1s", "pulumi", pver}
	paths := append(
		append(
			make([]string, 0, len(homedirPaths)+len(subdir)),
			homedirPaths...,
		),
		subdir...,
	)
	dir, err := homedir.Expand(path.Join(paths...))
	if err != nil {
		return "", wrapErr(err, ErrUnexpected)
	}
	return dir, nil
}
