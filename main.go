package main

import (
	"github.com/W-B-S/nbs-node/routing"
)

func main() {

	router := routing.GetInstance()
	router.Running()
}
