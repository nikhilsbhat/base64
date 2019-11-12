package cli

import (
	"github.com/spf13/cobra"
)

var (
	baseString string
)

// Registering all the flags to the command neuron itself.
func registerFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVarP(&baseString, "string", "s", "", "string which has to be encoded or decoded")
}
