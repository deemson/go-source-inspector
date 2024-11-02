package main

import (
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
	flag.StringVarP(&goRoot, "go-root", "r", os.Getenv("GOROOT"), strings.Join([]string{
		"set go install path for core packages look ups;",
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
	if workDir == "" {
		var err error
		workDir, err = os.Getwd()
		if err != nil {
			fmt.Printf("failed to get current working directory: %s\n", err.Error())
			os.Exit(1)
		}
	}
	if goRoot == "" {
		fmt.Println("go-root flag is not provided and GOROOT env variable is empty")
		os.Exit(1)
	}
	//for {
	//	data, err := bufio.NewReader(os.Stdin).ReadString('\n')
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	data = strings.TrimSpace(data)
	//	if data == "exit" {
	//		break
	//	}
	//	fmt.Print(data + "> ")
	//}
}
