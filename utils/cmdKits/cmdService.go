package cmdKits

import (
	"github.com/W-B-S/nbs-node/utils/config"
	"net"
	"os"
	"syscall"
)

func IsAddrInUse(err error) bool {
	e2, ok := err.(*net.OpError)
	if !ok {
		return false
	}
	e3, ok := e2.Err.(*os.SyscallError)
	if !ok {
		return false
	}
	return e3.Err == syscall.EADDRINUSE
}

func DialToCmdService() {
	test()
}

func StartCmdService() {

	var address = "127.0.0.1:" + config.GetConfig().CmdServicePort
	listener, err := net.Listen("tcp", address)

	if err != nil {

		logger.Fatal(err)
	}

	logger.Info("Starting to listening the incoming command...")

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

}
