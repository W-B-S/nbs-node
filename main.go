package main

import (
	"github.com/W-B-S/nbs-node/storage/routing"
	"github.com/W-B-S/nbs-node/utils/nbslog"
)

func main() {
	nbslog.Test()
	router := routing.GetInstance()
	router.Running()
}
