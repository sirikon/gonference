package templates

import (
	"bytes"
	"github.com/gobuffalo/packr/v2"
	"html/template"
)

// RenderTemplate .
func RenderTemplate(templateName string, data interface{}) ([]byte, error) {
	box := packr.New("template-files", "./files")
	content, err := box.FindString(templateName + ".html")
	if err != nil {
		return nil, err
	}
	tmpl, err := template.New(templateName).Parse(content)
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer
	if err := tmpl.Execute(&result, data); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
