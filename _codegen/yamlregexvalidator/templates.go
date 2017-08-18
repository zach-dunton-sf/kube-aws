package yamlregexvalidator

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	regex,
}

var regex = &typewriter.Template{
	Name: "Regex",
	Text: `
func (v {{.Pointer}}{{ .Name }}) UnmarshalYAML(f func(interface{}) error) error {
	var data string
	if err := f(&data); err != nil {
		return err
	}

	re := {{ .Name }}_regex

	if !regexp.MustCompile(re).MatchString(data) {
		return fmt.Errorf("{{.Name}}='%s' doesn't match validation regex='%s'", data, re)
	}

	*v = {{.Name}}(data)
	return nil
}
`}
