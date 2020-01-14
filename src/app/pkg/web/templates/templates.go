package templates

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"gonference/pkg/assets"
	"gonference/pkg/utils"
	"html/template"
	"net/http"
)

func ReplyTemplate(c *gin.Context, templateName string, data interface{}) {
	result, err := renderTemplate(templateName, data); utils.Check(err)
	c.Data(http.StatusOK, "text/html", result)
}

// RenderTemplate .
func renderTemplate(templateName string, data interface{}) (result []byte, err error) {
	defer utils.HandlePanic(&err)

	content := getTemplateContent(templateName)
	tmpl, err := template.New(templateName).
		Funcs(templateFunctions()).
		Parse(content); utils.Check(err)

	includeCommonTemplates(tmpl)

	var buffer bytes.Buffer
	utils.Check(tmpl.ExecuteTemplate(&buffer, "layout", data))

	return buffer.Bytes(), nil
}

func templateFunctions() template.FuncMap {
	return template.FuncMap{
		"markdown": func(text string) template.HTML {
			return template.HTML(blackfriday.Run([]byte(text)))
		},
	}
}

func includeCommonTemplates(tmpl *template.Template)  {
	var err error
	_, err = tmpl.Parse(getTemplateContent("_layout")); utils.Check(err)
	_, err = tmpl.Parse(getTemplateContent("_icons")); utils.Check(err)
	_, err = tmpl.Parse(getTemplateContent("_components")); utils.Check(err)
}

func getTemplateContent(templateName string) string {
	content, err := assets.WebTemplates.FindString(templateName + ".html")
	utils.Check(err)
	return content
}
