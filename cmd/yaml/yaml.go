package yaml

import (
	"fmt"
	"os"

	"github.com/henrywhitaker3/ci-bump/internal/files"
	"github.com/spf13/cobra"
)

var (
	dryRun  bool
	patches = []string{}
	minors  = []string{}
	majors  = []string{}
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "yaml [file]",
		Short: "Update yaml files",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			file, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}

			for _, p := range patches {
				up, err := files.Patch(file, p)
				if err != nil {
					return err
				}
				file = up
			}
			for _, p := range minors {
				up, err := files.Minor(file, p)
				if err != nil {
					return err
				}
				file = up
			}
			for _, p := range majors {
				up, err := files.Major(file, p)
				if err != nil {
					return err
				}
				file = up
			}

			if dryRun {
				fmt.Println(string(file))
				return nil
			}

			if err := os.WriteFile(args[0], file, 0644); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "If set, just prints the updated file to stdout")
	cmd.Flags().StringSliceVar(&patches, "patch", nil, "The key of fields to increment the patch version of")
	cmd.Flags().StringSliceVar(&minors, "minor", nil, "The key of fields to increment the minor version of")
	cmd.Flags().StringSliceVar(&majors, "major", nil, "The key of fields to increment the major version of")

	return cmd
}
