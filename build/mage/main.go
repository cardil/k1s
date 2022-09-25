package mage

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/magefile/mage/mage"
	"github.com/wavesoftware/go-retcode"
)

// Main is the main entry point for the mage command.
func Main() int {
	chdirBuild()
	return parseAndRun(os.Stdout, os.Stderr, os.Stdin, os.Args[1:])
}

func parseAndRun(stdout, stderr io.Writer, stdin io.Reader, args []string) int {
	inv, cmd, err := mage.Parse(stderr, stdout, args)
	inv.Stdin = stdin
	inv.Dir = "."
	inv.WorkDir = ".."
	return run(inv, cmd, err, stderr)
}

func run(inv mage.Invocation, cmd mage.Command, err error, stderr io.Writer) int {
	if errors.Is(err, flag.ErrHelp) {
		return 0
	}
	errlog := log.New(inv.Stderr, "", 0)
	if err != nil {
		errlog.Println("Error:", err)
		return retcode.Calc(err)
	}
	inv.Stderr = stderr

	switch cmd {
	case mage.Version, mage.Init, mage.Clean:
		return mage.ParseAndRun(inv.Stdout, inv.Stderr, inv.Stdin, inv.Args)
	case mage.CompileStatic:
		return mage.Invoke(inv)
	case mage.None:
		return mage.Invoke(inv)
	default:
		panic(fmt.Errorf("Unknown command type: %v", cmd))
	}
}

func chdirBuild() {
	if err := os.Chdir(builddir()); err != nil {
		panic(err)
	}
}

func builddir() string {
	_, file, _, _ := runtime.Caller(0) //nolint:dogsled
	return path.Dir(path.Dir(file))
}
