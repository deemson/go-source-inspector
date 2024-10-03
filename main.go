package main

import (
	"fmt"
	"golang.org/x/mod/modfile"
	"os"
	"path"
)

func errPanic(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	workingDirPath, err := os.Getwd()
	errPanic(err)
	goModPath := path.Join(workingDirPath, "go.mod")
	data, err := os.ReadFile(goModPath)
	errPanic(err)
	goModFile, err := modfile.Parse(goModPath, data, nil)
	errPanic(err)
	fmt.Println(goModFile)
}
