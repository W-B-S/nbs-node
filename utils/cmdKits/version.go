package cmdKits

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CurrentVersion = "0.01"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the current software's version.",
	Long:  `show the current software's version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Next BlockChain System version:" + CurrentVersion)
	},
}
