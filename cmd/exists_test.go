package cmd

import (
	"fmt"
	"testing"

	"github.com/bad33ndj3/execvalidator/pkg/execval/mocks"

	"github.com/stretchr/testify/mock"
)

func Test_ExistsCmdHappyFlow(t *testing.T) {
	execValidator := &mocks.Validator{}
	execValidator.On("Validate", mock.Anything).Return(nil)
	cmd := NewExistsCmd(execValidator)

	// set args for testing
	cmd.SetArgs([]string{
		fmt.Sprintf("--executables=%s", "ls,vim"),
	})

	// check if Execute on cmd returns no error
	if err := cmd.Execute(); err != nil {
		t.Errorf("cmd.Execute() error = %v", err)
	}

	execValidator.AssertNumberOfCalls(t, "Validate", 2)
}

func Test_ExistsCmdUnnamedArgs(t *testing.T) {
	execValidator := &mocks.Validator{}
	execValidator.On("Validate", mock.Anything).Return(nil)
	cmd := NewExistsCmd(execValidator)

	// set args for testing
	cmd.SetArgs([]string{
		"ls",
		"vim",
	})

	// check if Execute on cmd returns no error
	if err := cmd.Execute(); err != nil {
		t.Errorf("cmd.Execute() error = %v", err)
	}

	execValidator.AssertNumberOfCalls(t, "Validate", 2)
}
