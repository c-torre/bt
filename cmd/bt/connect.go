package bt

import (
	"github.com/spf13/cobra"
	"github.com/c-torre/bt/pkg/bt"
)

var connectCmd = &cobra.Command{
    Use:   "connect",
    Aliases: []string{"con"},
    Short:  "Connects to a Bluetooth device",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
        bt.Connect()
    },
}

func init() {
    rootCmd.AddCommand(connectCmd)
}
