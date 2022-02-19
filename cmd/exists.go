package cmd

import (
	"github.com/bad33ndj3/execvalidator/pkg/execval"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func NewExistsCmd(val execval.Validator) *cobra.Command {
	var executables []string
	existsCmd := &cobra.Command{
		Use:   "exists",
		Short: "Validate an executable exists",
		Long: `Exists searches for an executable named file in the
directories named by the PATH environment variable.
If file contains a slash, it is tried directly and the PATH is not consulted.
The result may be an absolute path or a path relative to the current directory.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(executables) == 0 {
				if len(args) == 0 {
					return cmd.Help()
				}
				executables = append(executables, args...)
			}

			for i := range executables {
				err := val.Validate(executables[i])
				if err != nil {
					return errors.Wrap(err, "failed to validate executable")
				}
			}

			return nil
		},
	}

	existsCmd.PersistentFlags().StringSliceVarP(&executables, "executables", "e", []string{}, "tags to be added to the build")

	return existsCmd
}
