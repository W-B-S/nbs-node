package cmdKits

import (
	"github.com/W-B-S/nbs-node/utils/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the current software's version.",
	Long:  `show the current software's version.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Next BlockChain System version:" + config.GetConfig().CurrentVersion)
	},
}
