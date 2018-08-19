package nbslog

import (
	"fmt"
	"github.com/op/go-logging"
	"os"
	"sync"
)

var instance *logging.Logger
var once sync.Once

func GetInstance() *logging.Logger {

	once.Do(func() {

		instance = newLogIns()
	})

	return instance
}

func newLogIns() *logging.Logger {

	log := logging.MustGetLogger("NBS")

	logFile, err := os.OpenFile("nbs.log",
		os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}

	fileBackend := logging.NewLogBackend(logFile, ">>>", 0)
	leveledFileBackend := logging.AddModuleLevel(fileBackend)
	leveledFileBackend.SetLevel(logging.DEBUG, "")

	cmdBackend := logging.NewLogBackend(os.Stderr, ">>>", 0)
	format := logging.MustStringFormatter(
		`%{color}%{time:01-02/15:04:05} %{longfunc} > %{level:.4s} %{message}%{color:reset}  <<< `,
	)
	formattedCmdBackend := logging.NewBackendFormatter(cmdBackend, format)

	logging.SetBackend(leveledFileBackend, formattedCmdBackend)

	return log
}

type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func SetLevel(level logging.Level, module string) {
	logging.SetLevel(level, module)
}

func Test() {

	log := GetInstance()

	log.Debugf("debug %s", Password("secret"))
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("error")
	log.Critical("critical")
}
