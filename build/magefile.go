//go:build mage

package main

import (
	"github.com/cardil/k1s/pkg/metadata"

	// mage:import
	"knative.dev/toolbox/magetasks"
	"knative.dev/toolbox/magetasks/config"
	"knative.dev/toolbox/magetasks/pkg/artifact"
	"knative.dev/toolbox/magetasks/pkg/artifact/platform"
	"knative.dev/toolbox/magetasks/pkg/checks"
	"knative.dev/toolbox/magetasks/pkg/git"
)

// Default target is set to binary.
//
//goland:noinspection GoUnusedGlobalVariable
var Default = magetasks.Build // nolint:deadcode,gochecknoglobals

func init() { //nolint:gochecknoinits
	cli := artifact.Binary{
		Metadata: config.Metadata{
			Name: "k1s",
		},
		Platforms: []artifact.Platform{
			{OS: platform.Linux, Architecture: platform.AMD64},
			{OS: platform.Linux, Architecture: platform.ARM64},
			{OS: platform.Mac, Architecture: platform.AMD64},
			{OS: platform.Mac, Architecture: platform.ARM64},
			{OS: platform.Windows, Architecture: platform.AMD64},
		},
	}
	magetasks.Configure(config.Config{
		Version: &config.Version{
			Path: metadata.VersionPath(),
			Resolver: git.NewVersionResolver(
				git.WithCache(config.Cache()),
			),
		},
		Artifacts: []config.Artifact{cli},
		Checks:    []config.Task{checks.GolangCiLint()},
	})
}
