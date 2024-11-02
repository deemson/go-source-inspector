package command

import "github.com/deemson/go-source-inspector/command/output"

type Output = output.Output

type Command interface {
	Execute(args ...string) Output
}
