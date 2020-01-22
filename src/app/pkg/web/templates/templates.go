package templates

import (
	"gonference/pkg/assets"
	"gonference/pkg/utils"
	"html/template"
)

func getTemplate(name string) *template.Template {
	if name == "" {
		return nil
	}
	if tmpl := getTemplateFromCache(name); tmpl != nil {
		return tmpl
	}
	content := getTemplateContent(name)
	tmpl := template.New(name)
	_, err := tmpl.Funcs(templateFunctions).Parse(content); utils.Check(err)
	if name[0] != '_' {
		includeCommonTemplates(tmpl)
	}
	setTemplateToCache(name, tmpl)
	return tmpl
}

func getTemplateContent(templateName string) string {
	content, err := assets.WebTemplates.FindString(templateName + ".html")
	utils.Check(err)
	return content
}

func includeCommonTemplates(tmpl *template.Template)  {
	extend(tmpl, "_layout")
	extend(tmpl, "_icons")
	extend(tmpl, "_components")
}

func extend(base *template.Template, name string) {
	tmpl, err := getTemplate(name).Clone(); utils.Check(err)
	for _, t := range tmpl.Templates() {
		_, err := base.AddParseTree(t.Name(), t.Tree); utils.Check(err)
	}
}
