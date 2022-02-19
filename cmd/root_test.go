package cmd

import (
	"testing"
)

func Test_RootCommand(t *testing.T) {
	cmd := NewRootCmd()
	if err := cmd.Execute(); err != nil {
		t.Errorf("cmd.Execute() error = %v", err)
	}
}
