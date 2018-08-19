package application

import (
	"context"
	"fmt"
	"os"
	"sync"
)

type Application interface {
	addFile(file os.File) error
}

type NbsApplication struct {
}

var instance *NbsApplication
var once sync.Once
var parentContext context.Context

func GetInstance() *NbsApplication {

	once.Do(func() {

		parentContext = context.Background()

		app, err := newApplication()
		if err != nil {
			panic(err)
		}
		fmt.Printf("--->Create application to run......\n")

		instance = app
	})

	return instance
}

func newApplication() (*NbsApplication, error) {
	return &NbsApplication{}, nil
}

func (*NbsApplication) AddFile(file *os.File) error {
	return nil
}
