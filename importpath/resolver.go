package importpath

import "path"

type Resolver struct {
	WorkDir      string
	GoModuleName string
	GoRoot       string
	ModCacheDir  string
}

func (r *Resolver) Resolve(pkg string) []string {
	var result []string
	if r.GoRoot != "" {
		result = append(result, path.Join(r.GoRoot, "src", pkg))
	}
	return result
}
