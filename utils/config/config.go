package config

import (
	"os"
	"os/user"
	"path/filepath"
	"sync"
)

type Configure struct {
	BaseDir        string
	StorageDir     string
	LogFileName    string
	CmdServicePort string
	CurrentVersion string
}

const cmdServicePort = "6080"
const currentVersion = "0.0.1"

var config *Configure
var once sync.Once

func GetConfig() *Configure {
	once.Do(func() {
		config = initConfig()
	})

	return config
}

func initConfig() *Configure {

	baseDir := getBaseDir()
	if !FileExists(baseDir) {
		createDir(baseDir)
	}

	storageDir := filepath.Join(baseDir, string(filepath.Separator), "database")

	logFileName := filepath.Join(baseDir, string(filepath.Separator), "nbs.log")

	return &Configure{
		BaseDir:        baseDir,
		StorageDir:     storageDir,
		LogFileName:    logFileName,
		CmdServicePort: cmdServicePort,
		CurrentVersion: currentVersion,
	}
}

func getBaseDir() string {

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	baseDir := filepath.Join(usr.HomeDir, string(filepath.Separator), ".nbs")

	return baseDir
}

func FileExists(fileName string) bool {

	fileInfo, err := os.Lstat(fileName)

	if fileInfo != nil || (err != nil && !os.IsNotExist(err)) {
		return true
	}

	return false
}

func createDir(baseDir string) {

	err := os.Mkdir(baseDir, os.ModePerm)
	if err != nil {
		panic(err)
	}
}
