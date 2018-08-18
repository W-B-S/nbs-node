package routing

import (
	"context"
	"fmt"
	"github.com/W-B-S/nbs-node/network"
	"github.com/libp2p/go-libp2p-peer"
	"github.com/libp2p/go-libp2p-peerstore"
	"sync"
)

type NbsDHT struct {
	peerId peer.ID
}

var instance *NbsDHT
var once sync.Once
var parentContext context.Context

func GetInstance() *NbsDHT {
	once.Do(func() {
		parentContext = context.Background()
		router, err := NewNbsDht()
		if err != nil {
			panic(err)
		}
		fmt.Printf("router start to run......\n")
		instance = router
	})

	return instance
}

func NewNbsDht() (*NbsDHT, error) {

	host := network.GetInstance()

	distributeTable := &NbsDHT{
		peerId: peer.ID(host.ID()),
	}

	return distributeTable, nil
}

//----------->routing interface implementation<-----------//
func (*NbsDHT) Ping(context.Context, peer.ID) error {
	return nil
}

func (*NbsDHT) FindPeer(context.Context, peer.ID) (peerstore.PeerInfo, error) {
	return peerstore.PeerInfo{}, nil
}

func (*NbsDHT) PutValue(context.Context, string, []byte) error {
	return nil
}

func (*NbsDHT) GetValue(context.Context, string) ([]byte, error) {
	return nil, nil
}

//-----------> nbs distributed hash table functions<-----------//
func (router *NbsDHT) Running() {

	select {
	case <-parentContext.Done():
		fmt.Printf("routing node done!\n")
	}
}
