package yamlsimpleparse

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewYAMLSimpleParse())
	if err != nil {
		panic(err)
	}
}

// typewriter to validate strings against regex at yaml parsing time
type YAMLSimpleParse struct{}

func NewYAMLSimpleParse() *YAMLSimpleParse {
	return &YAMLSimpleParse{}
}

func (sw *YAMLSimpleParse) Name() string {
	return "parse"
}

func (sw *YAMLSimpleParse) Imports(t typewriter.Type) ([]typewriter.ImportSpec) {
	return []typewriter.ImportSpec{}
}

func (sw *YAMLSimpleParse) Write(w io.Writer, t typewriter.Type) error {
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
