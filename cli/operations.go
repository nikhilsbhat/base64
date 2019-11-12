package cli

import (
	"github.com/spf13/cobra"
)

func encode(cmd *cobra.Command, args []string) {
	if len(baseString) == 0 {
		cm.NeuronSaysItsError("The string is not passed hence cannot proceed further")
	} else {
		if err := base64Encode(baseString); err != nil {
			cm.NeuronSaysItsError(getStringOfMessage(err))
		}
	}
}

func decode(cmd *cobra.Command, args []string) {
	if len(baseString) == 0 {
		cm.NeuronSaysItsError("The string is not passed hence cannot proceed further")
	} else {
		if err := base64Dcode(baseString); err != nil {
			cm.NeuronSaysItsError(getStringOfMessage(err))
		}
	}
}
