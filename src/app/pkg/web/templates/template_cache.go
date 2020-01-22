package templates

import (
	"html/template"
	"os"
	"sync"
)


var cacheEnabled = os.Getenv("TEMPLATE_CACHE") == "true"
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
