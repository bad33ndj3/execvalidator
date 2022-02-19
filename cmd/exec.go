package cmd

import (
	"log"

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
				log.Println("args", args[i])
				err := val.Validate(args[i])
				if err != nil {
					return errors.Wrap(err, "failed to validate executable")
				}
			}

			return nil
		},
	}

	exec.PersistentFlags().StringSliceP("executables", "e", []string{}, "tags to be added to the build")

	return exec
}
