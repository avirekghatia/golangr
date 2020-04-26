package gohelper

import (
	"os"
)

func getGoBin() string {
	return os.Getenv("GOROOT")
}

func getGoPath() string {
	return os.Getenv("GOPATH")
}
