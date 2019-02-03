package http

import (
	"bytes"
	"html/template"
)

// RenderTemplate .
func RenderTemplate(templateName string, data interface{}) ([]byte, error) {
	tmpl, err := template.ParseFiles("./http/templates/" + templateName + ".html")
	if err != nil {
		return nil, err
	}

	var result bytes.Buffer
	if err := tmpl.Execute(&result, data); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
