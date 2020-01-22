package templates

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"gonference/pkg/utils"
	"net/http"
)

func ReplyTemplate(c *gin.Context, templateName string, data interface{}) {
	result := renderTemplate(templateName, data)
	c.Data(http.StatusOK, "text/html", result)
}

// RenderTemplate .
func renderTemplate(templateName string, data interface{}) (result []byte) {
	tmpl := getTemplate(templateName)
	var buffer bytes.Buffer
	utils.Check(tmpl.ExecuteTemplate(&buffer, "layout", data))
	return buffer.Bytes()
}
