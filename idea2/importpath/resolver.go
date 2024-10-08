package importpath

import (
	"path"
	"strings"
	"unicode"
)

type Resolver struct {
	WorkDir      string
	GoModuleName string
	GoRoot       string
	ModCacheDir  string
}

func (r *Resolver) Resolve(pkg string) []string {
	var result []string
	if r.ModCacheDir != "" {
		result = append(result, path.Join(r.ModCacheDir, NormalizeForGoModCache(pkg)))
	}
	if r.GoRoot != "" {
		result = append(result, path.Join(r.GoRoot, "src", pkg))
	}
	return result
}

func NormalizeForGoModCache(input string) string {
	var result strings.Builder
	for _, char := range input {
		if unicode.IsUpper(char) {
			result.WriteRune('!')
			result.WriteRune(unicode.ToLower(char))
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}
