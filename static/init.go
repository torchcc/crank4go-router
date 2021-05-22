package test

import (
	"errors"
	"fmt"
	"path"
	"runtime"
)

var StaticDir = ""

func getStaticDir() {
	StaticDir = path.Dir(getCurrentFile())

	if StaticDir == "" {
		panic(errors.New("can not get current file info"))
	} else {
		fmt.Println("static dir is " + StaticDir)

	}
}

func getCurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(errors.New("can not get current file info"))
	}
	return file
}

func init() {
	getStaticDir()
}
