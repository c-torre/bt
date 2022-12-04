package bluetooth

import (
	"github.com/spf13/cobra"
	"github.com/c-torre/bt/pkg/bt"
)

var disconnectCmd = &cobra.Command{
    Use:   "disconnect",
    Aliases: []string{"discon"},
    Short:  "Disconnect from a Bluetooth device",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
        bluetooth.Disconnect()
    },
}

func init() {
    rootCmd.AddCommand(disconnectCmd)
}
