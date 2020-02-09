package templates

import (
	"gonference/pkg/infrastructure/config"
	"html/template"
	"sync"
)

var cacheEnabled = config.Config.Web.TemplateCache
var templateCache = make(map[string]*template.Template)
var templateCacheMutex = sync.Mutex{}

func setTemplateToCache(name string, tmpl *template.Template)  {
	templateCacheMutex.Lock()
	defer templateCacheMutex.Unlock()
	if cacheEnabled {
		templateCache[name] = tmpl
	}
}

func getTemplateFromCache(name string) *template.Template {
	templateCacheMutex.Lock()
	defer templateCacheMutex.Unlock()
	if value, ok := templateCache[name]; ok {
		return value
	}
	return nil
}
