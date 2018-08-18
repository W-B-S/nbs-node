package network

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"gx/ipfs/QmNmJZL7FQySMtE2BQuLMuZg2EB2CLEunJJUSVSc9YnnbV/go-libp2p-host"
	"sync"
)

var p2pHost host.Host
var once sync.Once
var Context context.Context

func GetInstance() host.Host {

	once.Do(func() {
		Context = context.Background()
		host, err := libp2p.New(Context)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Create host  %s\n", host.ID())

		p2pHost = host
	})

	return p2pHost
}
