package importresolver_test

import (
	"encoding/base64"
	"fmt"
	"github.com/deemson/go-source-inspector/importresolver"
	"os"
	"path"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	var createGoMod = func(t *testing.T, dir string, lines []string) {
		t.Helper()
		filePath := path.Join(dir, "go.mod")
		if err := os.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0755); err != nil {
			t.Fatalf(`failed to write "%s"`, filePath)
		}
	}
	testCases := map[string]struct {
		prepDir     func(t *testing.T, dir string)
		expectedErr func(dir string) string
	}{
		"no go.mod": {
			expectedErr: func(dir string) string {
				return fmt.Sprintf(`failed to read go.mod at "%s/go.mod": open %s/go.mod: no such file or directory`, dir, dir)
			},
		},
		"empty go.mod": {
			prepDir: func(t *testing.T, dir string) {
				createGoMod(t, dir, nil)
			},
		},
		"bad go.mod": {
			prepDir: func(t *testing.T, dir string) {
				createGoMod(t, dir, []string{
					"bad",
				})
			},
			expectedErr: func(dir string) string {
				return fmt.Sprintf(`failed to parse go.mod at "%s/go.mod": %s/go.mod:1: unknown directive: bad`, dir, dir)
			},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testDir := path.Join(
				os.TempDir(),
				base64.StdEncoding.EncodeToString([]byte(name)),
			)
			if err := os.Mkdir(testDir, 0755); err != nil {
				t.Fatalf(`failed to create dir "%s"`, testDir)
			}
			defer func(t *testing.T) {
				if err := os.RemoveAll(testDir); err != nil {
					t.Fatalf(`failed to remove dir "%s"`, testDir)
				}
			}(t)
			if testCase.prepDir != nil {
				testCase.prepDir(t, testDir)
			}
			importResolver, err := importresolver.New(testDir)
			if testCase.expectedErr != nil {
				expectedErr := testCase.expectedErr(testDir)
				if err == nil {
					t.Errorf(`expected error to say "%s", but it's nil`, expectedErr)
				}

				if err.Error() != expectedErr {
					t.Errorf(`expected error to say "%s", but it says "%s"`, expectedErr, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf(`got unexpected error "%s"`, err.Error())
				}
				if importResolver == nil {
					t.Error(`importResolver is nil`)
				}
			}
		})
	}
}

func TestImportResolver_Resolve(t *testing.T) {
	goRoot := os.Getenv("GOROOT")
	if goRoot == "" {
		t.Fatalf("GOROOT is not set")
	}
	var createDir = func(t *testing.T, testDir, dirPath string) {
		fullDirPath := path.Join(testDir, dirPath)
		if err := os.Mkdir(fullDirPath, 0755); err != nil {
			t.Fatalf(`failed to create dir "%s"`, testDir)
		}
	}
	var createFile = func(t *testing.T, testDir, filePath string, lines []string) {
		t.Helper()
		fullFilePath := path.Join(testDir, filePath)
		if err := os.WriteFile(fullFilePath, []byte(strings.Join(lines, "\n")), 0755); err != nil {
			t.Fatalf(`failed to write "%s"`, filePath)
		}
	}
	testCases := map[string]struct {
		goMod       []string
		prepDir     func(t *testing.T, dir string)
		expected    string
		expectedErr string
	}{
		"fmt": {
			goMod:    []string{},
			expected: path.Join(goRoot, "src", "fmt"),
		},
		"some": {
			prepDir: func(t *testing.T, dir string) {
				createDir(t, dir, "some")
				createFile(t, dir, "some/some.go", []string{"package some"})
			},
		},
	}
	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			testDir := path.Join(
				os.TempDir(),
				base64.StdEncoding.EncodeToString([]byte(name)),
			)
			if err := os.Mkdir(testDir, 0755); err != nil {
				t.Fatalf(`failed to create dir "%s"`, testDir)
			}
			defer func(t *testing.T) {
				if err := os.RemoveAll(testDir); err != nil {
					t.Fatalf(`failed to remove dir "%s"`, testDir)
				}
			}(t)
			createFile(t, testDir, "go.mod", testCase.goMod)
			if testCase.prepDir != nil {
				testCase.prepDir(t, testDir)
			}
			importResolver, err := importresolver.New(testDir)
			if err != nil {
				t.Fatalf(`failed to initialize import resolver: %s`, err.Error())
			}
			actual, err := importResolver.Resolve(name)
			if testCase.expected != actual {
				t.Errorf(`expected "%s", got "%s"`, testCase.expected, actual)
			}
		})
	}
}
