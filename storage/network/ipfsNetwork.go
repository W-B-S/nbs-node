package network

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-host"
	"sync"
)

var p2pHost host.Host
var once sync.Once
var Context context.Context

func GetInstance() host.Host {

	once.Do(func() {
		Context = context.Background()
		newHost, err := libp2p.New(Context)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Create host  %s\n", newHost.ID())

		//--->convert gx version control to github one's
		p2pHost = newHost.(host.Host)
	})

	return p2pHost
}
