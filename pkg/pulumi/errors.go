package pulumi

import (
	"errors"
	"fmt"
)

var (
	// ErrBug is returned when a bug is detected in the code.
	ErrBug = errors.New("probably a bug")

	// ErrUnexpected is returned when an unexpected error is encountered.
	ErrUnexpected = errors.New("unexpected")
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
