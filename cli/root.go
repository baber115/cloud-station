package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "cloud-station-cli",
	Short:   "cloud-station-cli 云中转站",
	Long:    "cloud-station-cli 云中转站",
	Example: "cloud-station-cli cmds",
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			fmt.Println("cloud-station-cli v0.0.1")
		}
		return nil
	},
}

var (
	version bool
)

func init() {
	f := RootCmd.PersistentFlags()
	f.BoolVarP(&version, "version", "v", false, "cloud station 版本信息")
}
