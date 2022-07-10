package main

import (
	"bytes"
	"testing"

	"github.com/cardil/k1s/internal/cli"
	"github.com/stretchr/testify/assert"
	"github.com/wavesoftware/go-commandline"
)

func TestTheMain(t *testing.T) {
	s := capture(func() {
		main()
	})

	assert.Equal(t, 0, s.exitCode)
}

type state struct {
	exitCode int
	out      bytes.Buffer
}

func (s *state) opts() []commandline.Option {
	return []commandline.Option{
		commandline.WithArgs("--help"),
		commandline.WithOutput(&s.out),
		commandline.WithExit(func(code int) {
			s.exitCode = code
		}),
	}
}

func capture(fn func()) state {
	var s state
	withOpts(fn, s.opts())
	return s
}

func withOpts(fn func(), opts []commandline.Option) {
	keep := cli.Opts
	defer func() {
		cli.Opts = keep
	}()
	cli.Opts = opts
	fn()
}
