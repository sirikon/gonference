package templates

import (
	"bytes"
	"github.com/yuin/goldmark"
	"gonference/pkg/infrastructure/config"
	"gonference/pkg/utils"
	"html/template"
)

const customMetaFilepath = "custom/meta.html"
const customMetaPostFilepath = "custom/meta-post.html"
var customLogoStaticPath = config.Config.Custom.LogoPath
var customBrandName = config.Config.Custom.BrandName

var templateFunctions = template.FuncMap{
	"markdown": markdownFunc,
	"custom_meta": fileIfExistsFunc(customMetaFilepath),
	"custom_meta_post": fileIfExistsFunc(customMetaPostFilepath),
	"custom_logo_path": customLogoFunc,
	"brand_name": brandNameFunc,
	"base_url": baseUrlFunc,
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

func baseUrlFunc() string {
	return config.Config.Web.BaseURL
}

func iterateFunc(count int) []int {
	result := make([]int, 0)
	for i := 0; i < count; i++ {
		result = append(result, i)
	}
	return result
}
