package yamlsimpleparse

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	parse,
}

var parse = &typewriter.Template{
	Name: "Parse",
	Text: `
func (v *{{ .Name }}) UnmarshalYAML(f func(interface{}) error) error {
	var data string
	if err := f(&data); err != nil {
		return err
	}

        res, err := {{ .Name }}_parse(data)
	if err != nil {
              return err
	}

        *v = {{.Name}}(res)
        return nil

}
`}
