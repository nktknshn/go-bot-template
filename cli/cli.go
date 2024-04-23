package cli

import (
	"github.com/nktknshn/go-bot-template/cli/clibot"
	"github.com/spf13/cobra"
)

func init() {
	CmdRoot.AddCommand(clibot.CmdBot)
}

var CmdRoot = &cobra.Command{
	Use:   "cli",
	Short: "cli is a command line interface for twitter",
}
