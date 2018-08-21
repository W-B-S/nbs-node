package cmdRpc

import (
	"bufio"
	"github.com/W-B-S/nbs-node/utils/config"
	"github.com/W-B-S/nbs-node/utils/nbsLog"
	"net"
)

var logger = nbsLog.GetInstance()

func readFromRPC(connection net.Conn) (string, error) {
	return bufio.NewReader(connection).ReadString('\n')
}

func DialToCmdService(cmdStr string) string {
	var address = "127.0.0.1:" + config.GetConfig().CmdServicePort

	logger.Info("Start to dial address:", address)

	connection, err := net.Dial("tcp", address)
	if err != nil {
		logger.Fatal("The nbs daemon is not running.")
	}

	defer connection.Close()
	logger.Info("Dial server success......")

	_, err = connection.Write([]byte(cmdStr))
	if err != nil {
		logger.Fatal("Send command to nbs node failed.")
	}
	logger.Info("Write request success......")

	result, err := readFromRPC(connection)
	if err != nil {
		return err.Error()
	} else {
		return result
	}
}

func StartCmdService() {

	var address = "127.0.0.1:" + config.GetConfig().CmdServicePort
	listener, err := net.Listen("tcp", address)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Starting to listening the incoming command at:", address)

	for {
		connection, err := listener.Accept()

		if err != nil {
			logger.Warning(err)
			continue
		}

		go handleInputCmd(connection)
	}
}

func handleInputCmd(connection net.Conn) {

	defer connection.Close()

	logger.Info("Received new connection success......")

	cmdStr, err := readFromRPC(connection)
	if err != nil {
		logger.Error("Failed to receive command from client")
		return
	}

	logger.Info("Read cmd string success:", cmdStr)

	_, err = connection.Write([]byte("success"))
	if err != nil {
		logger.Error("Failed to send action result to client")
		return
	}
	logger.Info("Write result back success")
}
