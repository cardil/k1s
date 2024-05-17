package mage

import (
	"path"
	"runtime"

	"knative.dev/toolbox/magetasks/entrypoint"
)

// Main is the main entry point for the mage command.
func Main() int {
	bd := builddir()
	ctx := entrypoint.Context{
		Directories: entrypoint.Directories{
			BuildDir:   bd,
			ProjectDir: path.Dir(bd),
			CacheDir:   path.Join(bd, "_output"),
		},
	}
	return entrypoint.Execute(ctx)
}

func builddir() string {
	_, file, _, _ := runtime.Caller(0) //nolint:dogsled
	return path.Dir(path.Dir(file))
}
