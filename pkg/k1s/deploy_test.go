package k1s_test

import (
	"bytes"
	"context"
	"io"
	"testing"

	"github.com/cardil/k1s/pkg/k1s"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	tc := &testCmd{}
	tc.ctx = context.Background()
	err := k1s.Deploy(tc, "test")
	require.NoError(t, err)
}

type testCmd struct {
	ctx context.Context //nolint:containedctx
	buf bytes.Buffer
}

func (t *testCmd) Context() context.Context {
	return t.ctx
}

func (t *testCmd) OutOrStdout() io.Writer {
	return &t.buf
}
