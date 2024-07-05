package semver

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
)

type Semver struct {
	semver  semver.Version
	vPrefix bool
}

func Parse(input string) (Semver, error) {
	var out Semver
	sv, err := semver.NewVersion(input)
	if err != nil {
		return out, fmt.Errorf("unable to parse version %s: %w", input, err)
	}
	out.semver = *sv

	if input[0] == 'v' {
		out.vPrefix = true
	}

	return out, nil
}

func (s *Semver) Patch() {
	s.semver = s.semver.IncPatch()
}

func (s *Semver) Minor() {
	s.semver = s.semver.IncMinor()
}

func (s *Semver) Major() {
	s.semver = s.semver.IncMajor()
}

func (s *Semver) String() string {
	ver := s.semver.String()
	if s.vPrefix {
		ver = fmt.Sprintf("v%s", ver)
	}
	return ver
}
