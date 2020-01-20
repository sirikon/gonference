package templates

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/yuin/goldmark"
	"gonference/pkg/assets"
	"gonference/pkg/utils"
	"html/template"
	"net/http"
	"os"
)

const customMetaFilepath = "custom/meta.html"
const customMetaPostFilepath = "custom/meta-post.html"
var customLogoStaticPath = os.Getenv("CUSTOM_LOGO_PATH")
var customBrandName = os.Getenv("CUSTOM_BRAND_NAME")

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
			return template.HTML(markdownToHTML(text))
		},
		"custom_meta": fileIfExists(customMetaFilepath),
		"custom_meta_post": fileIfExists(customMetaPostFilepath),
		"custom_logo_path": func() string {
			return customLogoStaticPath
		},
		"brand_name": func() string {
			if customBrandName != "" {
				return customBrandName
			}
			return "Gonference"
		},
	}
}

func fileIfExists(filepath string) func() template.HTML {
	return func() template.HTML {
		if utils.FileExists(filepath) {
			return template.HTML(utils.ReadFile(filepath))
		}
		return ""
	}
}

func markdownToHTML(source string) string {
	var buf bytes.Buffer
	utils.Check(goldmark.Convert([]byte(source), &buf))
	return buf.String()
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
