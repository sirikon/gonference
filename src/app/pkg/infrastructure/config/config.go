package config

type DatabaseConfig struct {
	URL string
}

type WebConfig struct {
	Port string
	BaseURL string
	CookieSecret string
	TemplateCache bool
}

type CustomConfig struct {
	LogoPath string
	BrandName string
}

type Config struct {
	Database DatabaseConfig
	Web WebConfig
	Custom CustomConfig
}
