package templates

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
	"github.com/sirikon/gonference/src/utils"
	"html/template"
	"net/http"
)

func ReplyTemplate(c *gin.Context, templateName string, data interface{}) {
	result, err := renderTemplate(templateName, data)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.Data(http.StatusOK, "text/html", result)
}

// RenderTemplate .
func renderTemplate(templateName string, data interface{}) (result []byte, err error) {
	defer utils.HandlePanic(&err)

	content := getTemplateContent(templateName)
	tmpl, err := template.New(templateName).Parse(content); utils.HandleErr(err)

	includeCommonTemplates(tmpl)

	var buffer bytes.Buffer
	utils.HandleErr(tmpl.ExecuteTemplate(&buffer, "layout", data))

	return buffer.Bytes(), nil
}

func includeCommonTemplates(tmpl *template.Template)  {
	var err error
	_, err = tmpl.Parse(getTemplateContent("_layout")); utils.HandleErr(err)
	_, err = tmpl.Parse(getTemplateContent("_icons")); utils.HandleErr(err)
}

func getTemplateContent(templateName string) string {
	box := packr.New("template-files", "./files")
	content, err := box.FindString(templateName + ".html")
	utils.HandleErr(err)
	return content
}
