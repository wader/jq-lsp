package lsp_test

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/wader/jq-lsp/internal/difftest"
	"github.com/wader/jq-lsp/lsp"
)

var update = flag.Bool("update", false, "Update tests")

// replace "@path" with content of file
func replaceTextFile(t *testing.T, v any, readFile func(s string) ([]byte, error)) any {
	switch v := v.(type) {
	case []any:
		a := make([]any, len(v))
		for i, ve := range v {
			a[i] = replaceTextFile(t, ve, readFile)
		}
		return a
	case map[string]any:
		a := make(map[string]any, len(v))
		for k, ve := range v {
			a[k] = replaceTextFile(t, ve, readFile)
		}
		return a
	case string:
		if _, path, found := strings.Cut(v, "@"); found {
			if b, err := readFile(path); err == nil {
				return string(b)
			} else {
				t.Fatal(err)
			}
		}
		return v
	default:
		return v
	}
}

func TestLSP(t *testing.T) {
	difftest.TestWithOptions(t, difftest.Options{
		Path:        "testdata",
		Pattern:     "*.json",
		ColorDiff:   os.Getenv("TEST_COLOR") != "",
		WriteOutput: *update,
		Fn: func(t *testing.T, path string, input string) (string, string, error) {
			testBaseDir := filepath.Dir(path)

			readFile := func(s string) ([]byte, error) {
				return os.ReadFile(filepath.Join(testBaseDir, s))
			}
			readFileOrEmpty := func(s string) []byte {
				b, err := readFile(s)
				if err != nil {
					return nil
				}
				return b
			}

			type request struct {
				Request  any `json:"request"`
				Response any `json:"response"`
			}
			var requests []request
			err := json.Unmarshal([]byte(input), &requests)
			if err != nil {
				t.Fatal(err)
			}

			for i := range requests {
				requests[i].Response = map[string]any{}
			}

			stdinBuf := &bytes.Buffer{}
			for _, r := range requests {
				rr := replaceTextFile(t, r.Request, readFile)
				reqB, err := json.Marshal(&rr)
				if err != nil {
					t.Fatal(err)
				}
				fmt.Fprintf(stdinBuf, "Content-Length: %d\r\n\r\n", len(reqB))
				stdinBuf.Write(reqB)
			}

			actualStdout := &bytes.Buffer{}
			actualStderr := &bytes.Buffer{}
			err = lsp.Run(lsp.Env{
				Version:  "test-version",
				ReadFile: readFile,
				Stdin:    stdinBuf,
				Stdout:   actualStdout,
				Stderr:   actualStderr,
				Args:     strings.Split(string(readFileOrEmpty("args")), " "),
				Environ:  strings.Split(string(readFileOrEmpty("environ")), " "),
			})
			if err != nil {
				return "", "", err
			}

			s := actualStdout.String()
			i := 0
			for {
				h, rest, found := strings.Cut(s, "\r\n\r\n")
				if !found {
					break
				}
				kv := map[string]string{}
				for _, line := range strings.Split(h, "\r\n") {
					k, v, _ := strings.Cut(line, ": ")
					kv[k] = v
				}

				contentLength, _ := strconv.Atoi(kv["Content-Length"])

				var resp any
				if err := json.Unmarshal([]byte(rest[0:contentLength]), &resp); err != nil {
					t.Fatal(err)
				}

				requests[i].Response = resp
				i++

				s = rest[contentLength:]
			}

			actualRequests := &bytes.Buffer{}
			je := json.NewEncoder(actualRequests)
			je.SetIndent("", "    ")
			if err := je.Encode(&requests); err != nil {
				t.Fatal(err)
			}

			return path, actualRequests.String(), nil
		},
	})
}
