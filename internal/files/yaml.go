package files

import (
	"errors"
	"fmt"
	"strings"

	"github.com/henrywhitaker3/ci-bump/internal/semver"
	"github.com/mikefarah/yq/v4/pkg/yqlib"
	glog "gopkg.in/op/go-logging.v1"
)

func init() {
	glog.SetLevel(glog.CRITICAL, "yq-lib")
}

func Patch(file []byte, field string) ([]byte, error) {
	raw, err := getValue(file, field)
	if err != nil {
		return nil, err
	}
	sv, err := semver.Parse(string(raw))
	if err != nil {
		return nil, err
	}
	sv.Patch()
	return updateValue(file, field, sv.String())
}

func Minor(file []byte, field string) ([]byte, error) {
	raw, err := getValue(file, field)
	if err != nil {
		return nil, err
	}
	sv, err := semver.Parse(string(raw))
	if err != nil {
		return nil, err
	}
	sv.Minor()
	return updateValue(file, field, sv.String())
}

func Major(file []byte, field string) ([]byte, error) {
	raw, err := getValue(file, field)
	if err != nil {
		return nil, err
	}
	sv, err := semver.Parse(string(raw))
	if err != nil {
		return nil, err
	}
	sv.Major()
	return updateValue(file, field, sv.String())
}

func Set(file []byte, field string, value string) ([]byte, error) {
	return updateValue(file, field, value)
}

func getValue(file []byte, field string) ([]byte, error) {
	eval := yqlib.NewStringEvaluator()
	out, err := eval.Evaluate(field, string(file), yqlib.NewYamlEncoder(yqlib.ConfiguredYamlPreferences), yqlib.NewYamlDecoder(yqlib.ConfiguredYamlPreferences))

	spl := strings.Split(out, "\n")
	if len(spl) > 2 {
		return nil, errors.New("more than one match returned")
	}

	return []byte(spl[0]), err
}

func updateValue(file []byte, field string, value string) ([]byte, error) {
	eval := yqlib.NewStringEvaluator()

	expr := fmt.Sprintf("%s = \"%s\"", field, value)

	out, err := eval.Evaluate(expr, string(file), yqlib.NewYamlEncoder(yqlib.ConfiguredYamlPreferences), yqlib.NewYamlDecoder(yqlib.ConfiguredYamlPreferences))
	return []byte(out), err
}
