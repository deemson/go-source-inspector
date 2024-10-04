package main

import (
	"fmt"
	"go/parser"
	"go/token"
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
	fileSet := token.NewFileSet()
	pkgs, err := parser.ParseDir(fileSet, workingDirPath, nil, 0)
	errPanic(err)
	fmt.Println(goModFile)
	fmt.Println(pkgs)
}
