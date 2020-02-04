package config

type DatabaseConfig struct {
	URL string
}

type WebConfig struct {
	Port int
	BaseURL string
	CookieSecret string
	TemplateCache bool
}

type CustomConfig struct {
	LogoPath string
	BrandName string
}

type LoggingConfig struct {
	AccessLog bool
}

type RootConfig struct {
	Database DatabaseConfig
	Web WebConfig
	Custom CustomConfig
	Logging  LoggingConfig
}
