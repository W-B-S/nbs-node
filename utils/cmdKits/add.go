package cmdKits

import (
	"github.com/W-B-S/nbs-node/storage/application"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add file to nbs network",
	Long:  `Add file to cache and find the peers to store it.`,
	Run:   addFile,
}

func addFile(cmd *cobra.Command, args []string) {

	DialToCmdService()

	app := application.GetInstance()

	file, err := os.Open(args[0])
	if err != nil {
		logger.Fatal("Failed to open file.")
		return
	}

	app.AddFile(file)
}

func test() {

}
