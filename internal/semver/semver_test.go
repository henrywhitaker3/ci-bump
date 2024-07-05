package semver

import "testing"

func TestItParsesInputs(t *testing.T) {
	type test struct {
		name        string
		input       string
		final       string
		shouldError bool
	}

	tcs := []test{
		{
			name:        "errors when parsing invalid verison",
			input:       "bongo",
			final:       "",
			shouldError: true,
		},
		{
			name:        "parses a semver without v prefix",
			input:       "1.0.4",
			final:       "1.0.4",
			shouldError: false,
		},
		{
			name:        "parses a semver with a v prefix",
			input:       "v1.0.4",
			final:       "v1.0.4",
			shouldError: false,
		},
	}

	for _, c := range tcs {
		t.Run(c.name, func(t *testing.T) {
			s, err := Parse(c.input)
			if c.shouldError {
				if err == nil {
					t.Fatal("should have errored but didn't")
				}
				return
			}

			if c.final != s.String() {
				t.Fatal("did not get the expected output, expected", c.final, "got", s.String())
			}
		})
	}
}

func TestItBumpsPatchVersion(t *testing.T) {
	type test struct {
		name  string
		input string
		final string
	}

	tcs := []test{
		{
			name:  "increments patch without v prefix",
			input: "1.0.4",
			final: "1.0.5",
		},
		{
			name:  "increments patch with v prefix",
			input: "v1.0.4",
			final: "v1.0.5",
		},
	}

	for _, c := range tcs {
		t.Run(c.name, func(t *testing.T) {
			s, err := Parse(c.input)
			if err != nil {
				t.Fatal(err)
			}
			s.Patch()
			if s.String() != c.final {
				t.Fatal("expected", c.final, "got", s.String())
			}
		})
	}
}

func TestItBumpsMinorVersion(t *testing.T) {
	type test struct {
		name  string
		input string
		final string
	}

	tcs := []test{
		{
			name:  "increments minor without v prefix",
			input: "1.0.4",
			final: "1.1.0",
		},
		{
			name:  "increments minor with v prefix",
			input: "v1.0.4",
			final: "v1.1.0",
		},
	}

	for _, c := range tcs {
		t.Run(c.name, func(t *testing.T) {
			s, err := Parse(c.input)
			if err != nil {
				t.Fatal(err)
			}
			s.Minor()
			if s.String() != c.final {
				t.Fatal("expected", c.final, "got", s.String())
			}
		})
	}
}

func TestItBumpsMajorVersion(t *testing.T) {
	type test struct {
		name  string
		input string
		final string
	}

	tcs := []test{
		{
			name:  "increments major without v prefix",
			input: "1.0.4",
			final: "2.0.0",
		},
		{
			name:  "increments major with v prefix",
			input: "v1.0.4",
			final: "v2.0.0",
		},
	}

	for _, c := range tcs {
		t.Run(c.name, func(t *testing.T) {
			s, err := Parse(c.input)
			if err != nil {
				t.Fatal(err)
			}
			s.Major()
			if s.String() != c.final {
				t.Fatal("expected", c.final, "got", s.String())
			}
		})
	}
}
