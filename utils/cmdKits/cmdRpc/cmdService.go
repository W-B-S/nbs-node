package cmdRpc

import (
	"github.com/W-B-S/nbs-node/utils/config"
	"github.com/W-B-S/nbs-node/utils/nbsLog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

var logger = nbsLog.GetInstance()

type server struct{}

func (s *server) RpcTask(ctx context.Context, req *CmdRequest) (*CmdReply, error) {
	return handleInputCmd(ctx, req)
}

func DialToCmdService(cmdStr string) string {

	var address = "127.0.0.1:" + config.GetConfig().CmdServicePort

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		logger.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewCmdTaskClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	response, err := client.RpcTask(ctx, &CmdRequest{Name: "add"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	logger.Info("Greeting: %s", response.Message)

	return response.Message
	/*


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
	*/
}

func StartCmdService() {

	var address = "127.0.0.1:" + config.GetConfig().CmdServicePort

	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	theServer := grpc.NewServer()

	RegisterCmdTaskServer(theServer, &server{})

	reflection.Register(theServer)
	if err := theServer.Serve(listener); err != nil {
		logger.Fatalf("failed to serve: %v", err)
	}

	/*
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
	*/
}

func handleInputCmd(ctx context.Context, req *CmdRequest) (*CmdReply, error) {

	return &CmdReply{Message: "Hello " + req.Name}, nil
}
