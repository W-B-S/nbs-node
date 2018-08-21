package cmdKits

import (
	"github.com/W-B-S/nbs-node/utils/cmdKits/cmdRpc"
	"github.com/spf13/cobra"
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

	logger.Info("add command args:", args)

	cmdRpc.DialToCmdService("nbs add nbs.log\n")

	logger.Info("Reading success......")

}

func addFileTask(fileName string) {
	//app := application.GetInstance()
	//
	//file, err := os.Open(fileName)
	//if err != nil {
	//	logger.Error("Failed to open file.")
	//	return
	//}
	//
	//app.AddFile(file)
}
