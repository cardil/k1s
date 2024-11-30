package pulumi

import (
	"context"
	"errors"
	"os"

	"github.com/google/uuid"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/optup"
	"github.com/zalando/go-keyring"
	"knative.dev/pkg/logging"
)

func newManagedPassphaseStack(fn func() (Stack, error)) (Stack, error) {
	return &managedPassphaseStack{fn: fn}, nil
}

type managedPassphaseStack struct {
	fn func() (Stack, error)
}

func (m *managedPassphaseStack) Up(ctx context.Context, opts ...optup.Option) (auto.UpResult, error) {
	if err := m.setupPassphase(); err != nil {
		return auto.UpResult{}, err
	}
	defer m.cleanUpPassphase(ctx)
	s, err := m.fn()
	if err != nil {
		return auto.UpResult{}, err
	}
	return s.Up(ctx, opts...) //nolint:wrapcheck
}

func (m *managedPassphaseStack) setupPassphase() error {
	secret, err := m.generateOrGetSecret()
	if err != nil {
		return err
	}
	if err = os.Setenv("PULUMI_CONFIG_PASSPHRASE", secret); err != nil {
		return wrapErr(err, ErrUnexpected)
	}
	return nil
}

func (m *managedPassphaseStack) cleanUpPassphase(ctx context.Context) {
	if err := os.Unsetenv("PULUMI_CONFIG_PASSPHRASE"); err != nil {
		log := logging.FromContext(ctx)
		log.Fatal(wrapErr(err, ErrUnexpected))
	}
}

func (m *managedPassphaseStack) generateOrGetSecret() (string, error) {
	if secret := os.Getenv("K1S_PULUMI_CONFIG_PASSPHRASE"); secret != "" {
		return secret, nil
	}

	service := "k1s"
	user := "anon"

	secret, err := keyring.Get(service, user)
	if err != nil {
		if !errors.Is(err, keyring.ErrNotFound) {
			return "", wrapErr(err, ErrUnexpected)
		}
		secret, err = generateSecret()
		if err != nil {
			return "", err
		}
		if serr := keyring.Set(service, user, secret); serr != nil {
			return "", wrapErr(serr, ErrUnexpected)
		}
	}
	return secret, nil
}

func generateSecret() (string, error) {
	secret, err := uuid.NewV7()
	if err != nil {
		return "", wrapErr(err, ErrUnexpected)
	}
	return secret.String(), nil
}
