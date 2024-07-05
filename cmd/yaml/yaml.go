package yaml

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/henrywhitaker3/ci-bump/internal/files"
	"github.com/spf13/cobra"
)

var (
	dryRun  bool
	patches = []string{}
	minors  = []string{}
	majors  = []string{}
	sets    = []string{}
)

func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "yaml [file]",
		Short: "Update yaml files",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			setFmt := [][2]string{}
			for _, sv := range sets {
				if !strings.Contains(sv, "=") {
					return errors.New("invaid --set format, must contain = separator")
				}
				spl := strings.Split(sv, "=")
				setFmt = append(setFmt, [2]string{spl[0], spl[1]})
			}
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
			for _, p := range setFmt {
				up, err := files.Set(file, p[0], p[1])
				if err != nil {
					return err
				}
				file = up
			}

			if !dryRun {
				if err := os.WriteFile(args[0], file, 0644); err != nil {
					return err
				}
			}

			fmt.Println(string(file))

			return nil
		},
	}

	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "If set, just prints the updated file to stdout")
	cmd.Flags().StringSliceVar(&patches, "patch", nil, "The key of fields to increment the patch version of")
	cmd.Flags().StringSliceVar(&minors, "minor", nil, "The key of fields to increment the minor version of")
	cmd.Flags().StringSliceVar(&majors, "major", nil, "The key of fields to increment the major version of")
	cmd.Flags().StringSliceVar(&sets, "set", nil, "The field to update and a value to set it to e.g. .appVersion=v1.0.0")

	return cmd
}
