package k1s

import (
	"context"
	"io"
)

type App interface {
	Context() context.Context
	OutOrStdout() io.Writer
}
