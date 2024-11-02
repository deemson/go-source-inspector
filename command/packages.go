package command

var _ Command = &Packages{}

type Packages struct {
}

func (c *Packages) Execute(args ...string) Output {
	return nil
}
