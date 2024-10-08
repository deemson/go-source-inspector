package importresolver

import (
	"errors"
	"fmt"
	"golang.org/x/mod/modfile"
	"os"
	"path"
)

func New(dir string) (*ImportResolver, error) {
	goModPath := path.Join(dir, "go.mod")
	goModData, err := os.ReadFile(goModPath)
	if err != nil {
		return nil, fmt.Errorf(`failed to read go.mod at "%s": %w`, goModPath, err)
	}
	_, err = modfile.Parse(goModPath, goModData, nil)
	if err != nil {
		return nil, fmt.Errorf(`failed to parse go.mod at "%s": %w`, goModPath, err)
	}
	goRootPath := os.Getenv("GOROOT")
	if goRootPath == "" {
		return nil, errors.New("GOROOT env var must be set")
	}
	return &ImportResolver{}, nil
}

type ImportResolver struct {
	goRootPath string
}

func (r *ImportResolver) Resolve(name string) (string, error) {
	return "", nil
}
