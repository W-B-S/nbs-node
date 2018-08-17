package main

import (
	"context"
	"fmt"

	"github.com/W-B-S/nbs-routing/nbsimpl"
	"github.com/libp2p/go-libp2p"
)

func main() {

	context, cancel := context.WithCancel(context.Background())
	defer cancel()

	host, err := libp2p.New(context)
	if err != nil {
		panic(err)
	}
	fmt.Printf("peer host  %s\n", host.ID())

	router, err := nbsrouting.NewNbsDht(context, host)
	if err != nil {
		panic(err)
	}
	fmt.Printf("router start to run......\n")

	router.Running()
}
