package cmd

import (
	"testing"

	"github.com/bad33ndj3/execvalidator/pkg/execval/mocks"

	"github.com/stretchr/testify/mock"
)

func Test_ExecuteCommand(t *testing.T) {
	execValidator := &mocks.Validator{}
	execValidator.On("Validate", mock.Anything).Return(nil)
	cmd := NewExecCommand(execValidator)

	// set args for testing
	cmd.SetArgs([]string{"ls"})

	// check if Execute on cmd returns no error
	if err := cmd.Execute(); err != nil {
		t.Errorf("cmd.Execute() error = %v", err)
	}

	execValidator.AssertNumberOfCalls(t, "Validate", 1)
}
