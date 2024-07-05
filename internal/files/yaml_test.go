package files

import (
	"fmt"
	"os"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestItGetsTheValueFromYaml(t *testing.T) {
	type test struct {
		name        string
		file        string
		accessor    string
		expects     string
		shouldError bool
	}

	tcs := []test{
		{
			name:        "gets value from root level key",
			file:        "yaml_root_key.yaml",
			accessor:    ".bongo",
			expects:     "value",
			shouldError: false,
		},
		{
			name:        "gets value from array single match",
			file:        "yaml_array_key_once_match.yaml",
			accessor:    ".bongo[0].version",
			expects:     "value",
			shouldError: false,
		},
		{
			name:        "gets value from array multi match",
			file:        "yaml_array_key_multi_match.yaml",
			accessor:    ".bongo[].version",
			shouldError: true,
		},
	}

	for _, c := range tcs {
		t.Run(c.name, func(t *testing.T) {
			file, err := os.ReadFile(fmt.Sprintf("testdata/%s", c.file))
			if err != nil {
				t.Fatal(err)
			}
			val, err := getValue(file, c.accessor)
			if c.shouldError {
				if err == nil {
					t.Fatal("expected an error but didn't get one")
				}
				return
			}
			if string(val) != c.expects {
				t.Fatal("expected", c.expects, "got", string(val))
			}
		})
	}
}

func TestItUpdatesTheValueInYaml(t *testing.T) {
	type test struct {
		name        string
		file        string
		accessor    string
		with        string
		expects     string
		shouldError bool
	}

	tcs := []test{
		{
			name:     "updates value from root level key",
			file:     "yaml_root_key.yaml",
			accessor: ".bongo",
			with:     "new",
			expects: `---
bongo: new
`,
			shouldError: false,
		},
		{
			name:     "updates value from array single match",
			file:     "yaml_array_key_once_match.yaml",
			accessor: ".bongo[0].version",
			with:     "new",
			expects: `---
bongo:
  - version: new
`,
			shouldError: false,
		},
		{
			name:     "gets value from array multi match",
			file:     "yaml_array_key_multi_match.yaml",
			accessor: ".bongo[].version",
			with:     "new",
			expects: `---
bongo:
  - version: new
  - version: new
`,
			shouldError: true,
		},
	}

	for _, c := range tcs {
		t.Run(c.name, func(t *testing.T) {
			file, err := os.ReadFile(fmt.Sprintf("testdata/%s", c.file))
			if err != nil {
				t.Fatal(err)
			}
			out, err := updateValue(file, c.accessor, c.with)
			if err != nil {
				if c.shouldError {
					return
				}
				t.Fatal(err)
			}
			if string(out) != c.expects {
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(string(out), c.expects, false)
				fmt.Println(string(out))
				fmt.Println(c.expects)
				t.Fatal(dmp.DiffPrettyText(diffs))
			}
		})
	}
}
