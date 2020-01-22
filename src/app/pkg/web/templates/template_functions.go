package templates

import (
	"bytes"
	"github.com/yuin/goldmark"
	"gonference/pkg/utils"
	"html/template"
	"os"
)

const customMetaFilepath = "custom/meta.html"
const customMetaPostFilepath = "custom/meta-post.html"
var customLogoStaticPath = os.Getenv("CUSTOM_LOGO_PATH")
var customBrandName = os.Getenv("CUSTOM_BRAND_NAME")

var templateFunctions = template.FuncMap{
	"markdown": markdownFunc,
	"custom_meta": fileIfExistsFunc(customMetaFilepath),
	"custom_meta_post": fileIfExistsFunc(customMetaPostFilepath),
	"custom_logo_path": customLogoFunc,
	"brand_name": brandNameFunc,
	"iterate": iterateFunc,
}

func markdownFunc(text string) template.HTML {
	var buf bytes.Buffer
	utils.Check(goldmark.Convert([]byte(text), &buf))
	return template.HTML(buf.String())
}

func fileIfExistsFunc(filepath string) func() template.HTML {
	return func() template.HTML {
		if utils.FileExists(filepath) {
			return template.HTML(utils.ReadFile(filepath))
		}
		return ""
	}
}

func customLogoFunc() string {
	return customLogoStaticPath
}

func brandNameFunc() string {
	if customBrandName != "" {
		return customBrandName
	}
	return "Gonference"
}

func iterateFunc(count int) []int {
	result := make([]int, 0)
	for i := 0; i < count; i++ {
		result = append(result, i)
	}
	return result
}
