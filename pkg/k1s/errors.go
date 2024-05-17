package k1s

import (
	"errors"
	"fmt"

	"github.com/cardil/k1s/pkg/pulumi"
)

var (
	// ErrBug is returned when a bug is detected in the code.
	ErrBug = pulumi.ErrBug

	// ErrUnexpected is returned when an unexpected error is encountered.
	ErrUnexpected = pulumi.ErrUnexpected
)

func wrapErr(err, target error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, target) {
		return err
	}
	return fmt.Errorf("%w: %w", target, err)
}
