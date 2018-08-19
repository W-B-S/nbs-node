package main

import (
	"github.com/W-B-S/nbs-node/storage/routing"
)

func main() {
	router := routing.GetInstance()
	router.Running()
}
