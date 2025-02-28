package pulumi

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path"

	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/sirupsen/logrus"
	"knative.dev/pkg/logging"
)

func newLoggedInStack(stackFn func() (Stack, error)) (Stack, error) {
	return &loggedInStack{stackFn: stackFn}, nil
}

type loggedInStack struct {
	stackFn func() (Stack, error)
}

func (l *loggedInStack) Up(ctx context.Context, opts ...optup.Option) (auto.UpResult, error) {
	if err := ensureLoggedInLocally(); err != nil {
		return auto.UpResult{}, err
	}
	defer l.cleanUp(ctx)
	s, err := l.stackFn()
	if err != nil {
		return auto.UpResult{}, err
	}
	return s.Up(ctx, opts...) //nolint:wrapcheck
}

func (l *loggedInStack) cleanUp(ctx context.Context) {
	log := logging.FromContext(ctx)
	if err := os.Unsetenv("PULUMI_HOME"); err != nil {
		log.Fatal(wrapErr(err, ErrBug))
	}
}

func ensureLoggedInLocally() error {
	var homeDir string

	{
		h, err := homeResource("config")
		if err != nil {
			return err
		}
		homeDir = h
	}

	if err := os.Setenv("PULUMI_HOME", homeDir); err != nil {
		return wrapErr(err, ErrBug)
	}

	loginTarget := "file://" + homeDir
	logrus.Debug("Checking Pulumi credentials file")
	credsFile := path.Join(homeDir, "credentials.json")
	if current, err := isPulumiCredentialsFileCurrent(credsFile, loginTarget); err != nil {
		return err
	} else if current {
		return nil
	}

	logrus.Info("Logging in to Pulumi")
	binDir := path.Join(path.Dir(homeDir), "bin")
	c := exec.Command("pulumi", "login", loginTarget)
	c.Path = path.Join(binDir, "pulumi")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		return wrapErr(err, ErrUnexpected)
	}
	logrus.Info("Logged in to Pulumi in local mode: ", homeDir)
	return nil
}

func isPulumiCredentialsFileCurrent(credsFile string, loginTarget string) (bool, error) {
	if _, err := os.Stat(credsFile); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, wrapErr(err, ErrUnexpected)
	}
	bytes, err := os.ReadFile(credsFile)
	if err != nil {
		return false, wrapErr(err, ErrUnexpected)
	}
	logrus.Debug("Found Pulumi credentials file: ", credsFile)
	var creds pulumiCredsFile
	if err = json.Unmarshal(bytes, &creds); err != nil {
		return false, wrapErr(err, ErrUnexpected)
	}
	if creds.Current == loginTarget {
		logrus.Info("Already logged in to Pulumi for: ", loginTarget)
		return true, nil
	}
	return false, nil
}

type pulumiCredsFile struct {
	Current string `json:"current"`
}
