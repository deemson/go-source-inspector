package importpath_test

import (
	"github.com/deemson/go-source-inspector/idea2/importpath"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolver_Resolve(t *testing.T) {
	testCases := map[string][]string{
		"fmt": {},
	}
	for pkg, expected := range testCases {
		t.Run(pkg, func(t *testing.T) {
			resolver := importpath.Resolver{
				WorkDir:      "work-dir",
				GoModuleName: "test-module",
				GoRoot:       "go-root",
				ModCacheDir:  "go-mod-cache",
			}
			actual := resolver.Resolve(pkg)
			assert.Equal(t, expected, actual)
		})
	}
}

func TestNormalizeForGoModCache(t *testing.T) {
	testCases := map[string]string{
		"hello":       "hello",
		"Hello":       "!hello",
		"HELLO":       "!h!e!l!l!o",
		"Hello-World": "!hello-!world",
	}
	for input, expected := range testCases {
		t.Run(input, func(t *testing.T) {
			actual := importpath.NormalizeForGoModCache(input)
			assert.Equal(t, expected, actual)
		})
	}
}
