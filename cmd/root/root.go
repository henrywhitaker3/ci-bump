package root

import (
	"github.com/henrywhitaker3/ci-bump/cmd/yaml"
	"github.com/spf13/cobra"
)

func NewCommand(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ci-bump",
		Version: version,
	}

	cmd.AddCommand(yaml.NewCommand())

	return cmd
}
