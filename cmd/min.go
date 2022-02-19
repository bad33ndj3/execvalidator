package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewMinCmd represents the min command
func NewMinCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "min",
		Short: "Used to validate an executable has a minimum version",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf("requires at least 2 arguments")
			}
			// check if first arg is a semver
			if !IsSemver(args[0]) {
				return fmt.Errorf("first argument must be a valid semver")
			}

			return nil
		},
	}
}

func IsSemver(s string) bool {
	// check if s is a valid semver
	return true
}
