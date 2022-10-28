package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// AppConfig contains site-wide configuration settings
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	Session       *scs.SessionManager
}
