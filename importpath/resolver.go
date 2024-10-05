package importpath

import "github.com/spf13/afero"

type Resolver struct {
	WorkDir     string
	GoRoot      string
	ModCacheDir string
	Fs          afero.Fs
}

func (r *Resolver) Resolve(pkg string) (string, error) {
	return "", nil
}
