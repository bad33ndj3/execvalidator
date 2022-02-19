package cmd

import (
	"github.com/bad33ndj3/execvalidator/pkg/execval"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewExecCommand(val execval.Validator) *cobra.Command {
	exec := &cobra.Command{
		Use:   "exec",
		Short: "Validate an executable exists",
		RunE: func(cmd *cobra.Command, args []string) error {
			for i := range args {
				err := val.Validate(args[i])
				if err != nil {
					return errors.Wrap(err, "failed to validate executable")
				}
			}

			return nil
		},
	}

	return exec
}
