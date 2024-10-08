package main

import (
	"bufio"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
	"strings"
)

var (
	help    bool
	workDir string
	goRoot  string
)

func init() {
	flag.CommandLine.SortFlags = false
	flag.BoolVarP(&help, "help", "h", false, "show help (this message)")
	flag.StringVarP(&workDir, "work-dir", "w", "", strings.Join([]string{
		"set working directory where go.mod is located;",
		"if empty then current directory is used",
	}, " "))
	flag.StringVarP(&goRoot, "goroot", "r", os.Getenv("GOROOT"), strings.Join([]string{
		"set go installation path for core packages look ups;",
		"defaults to the value of GOROOT env variable;",
		"will result in an error if both are empty",
	}, " "))
}

func main() {
	flag.Parse()
	if help {
		fmt.Println("go-source-inspector usage:")
		flag.PrintDefaults()
		os.Exit(0)
	}
	fmt.Print("> ")
	for {
		data, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			panic(err.Error())
		}
		if data == "exit" {
			break
		}
		fmt.Print(data + "> ")
	}

}
