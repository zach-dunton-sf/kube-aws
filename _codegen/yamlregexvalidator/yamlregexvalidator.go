package yamlregexvalidator

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewYAMLRegexValidator())
	if err != nil {
		panic(err)
	}
}

// typewriter to validate strings against regex at yaml parsing time
type YAMLRegexValidator struct{}

func NewYAMLRegexValidator() *YAMLRegexValidator {
	return &YAMLRegexValidator{}
}

func (sw *YAMLRegexValidator) Name() string {
	return "regex"
}

func (sw *YAMLRegexValidator) Imports(t typewriter.Type) ([]typewriter.ImportSpec) {
	return []typewriter.ImportSpec{ {"", "regexp"} }
}

func (sw *YAMLRegexValidator) Write(w io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(sw)

	if !found {
		// nothing to be done
		return nil
	}

	tmpl, err := templates.ByTag(t, tag)

	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, t); err != nil {
		return err
	}

	return nil
}
