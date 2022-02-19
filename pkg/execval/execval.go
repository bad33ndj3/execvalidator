package execval

import (
	"os/exec"
)

type Validator interface {
	Validate(executable string) error
}

type execValidator struct {
}

func (e *execValidator) Validate(executable string) error {
	_, err := exec.LookPath(executable)
	if err != nil {
		return err
	}

	return nil
}

// NewValidator returns a new Validator.
func NewValidator() Validator {
	return &execValidator{}
}
