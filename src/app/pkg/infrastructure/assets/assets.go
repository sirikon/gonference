package assets

import "github.com/gobuffalo/packr/v2"

var (
	WebTemplates = packr.New("assets-web-templates", "../../../resources/templates")
	DatabaseMigrations = packr.New("assets-database-migrations", "../../../resources/migrations")
	BackofficeUI = packr.New("assets-backoffice-ui", "../../../../backoffice-ui/dist")
	FrontStyle = packr.New("assets-front-style", "../../../../front-style/dist")
)
