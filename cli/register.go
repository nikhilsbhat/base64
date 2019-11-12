package cli

import (
	"github.com/spf13/cobra"
)

var (
	cmds map[string]*cobra.Command
)

type confcmds struct {
	commands []*cobra.Command
}

// SetConfigCmds helps in gathering all the subcommands so that it can be used while registering it with main command.
func SetConfigCmds() *cobra.Command {
	cmd := getConfigCmds()
	return cmd
}

func getConfigCmds() *cobra.Command {

	var configCmd = &cobra.Command{
		Use:   "base64 [command]",
		Short: "command to deal with base64 operations",
		Long:  `This command helps dealing with base64 operations such as encoding and decoding.`,
		Args:  cobra.MinimumNArgs(1),
		RunE:  cm.echoConfig,
	}
	configCmd.SetUsageTemplate(getUsageTemplate())

	var setDecodeCmd = &cobra.Command{
		Use:          "decode [flags]",
		Short:        "command to decode the string",
		Long:         `This will help user to base64 decode the string.`,
		Run:          decode,
		SilenceUsage: true,
	}

	var setEncodeCmd = &cobra.Command{
		Use:          "encode [flags]",
		Short:        "command to encode the string",
		Long:         `This will help user to base64 encode the string.`,
		Run:          encode,
		SilenceUsage: true,
	}
	configCmd.AddCommand(setDecodeCmd)
	configCmd.AddCommand(setEncodeCmd)
	registerFlags(configCmd)
	return configCmd
}

func (cm *cliMeta) echoConfig(cmd *cobra.Command, args []string) error {
	cmd.Usage()
	return nil
}

// This function will return the custom template for usage function,
// only functions/methods inside this package can call this.

func getUsageTemplate() string {
	return `{{printf "\n"}}Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if gt (len .Aliases) 0}}{{printf "\n" }}
Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}{{printf "\n" }}
Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}{{printf "\n"}}
Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}{{printf "\n"}}
Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}{{printf "\n"}}
Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}{{printf "\n"}}
Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}{{printf "\n"}}
Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}"
{{printf "\n"}}`
}
