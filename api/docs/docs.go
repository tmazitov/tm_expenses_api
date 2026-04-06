package docs

import (
	"html/template"

	"github.com/gofiber/contrib/v3/swaggo"
	"github.com/gofiber/fiber/v3"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Register(a *fiber.App) {
	a.Get("/docs/swagger.json", r.Spec())
	a.Get("/docs/*", r.Entrypoint())
}

func (r *Router) Spec() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		return ctx.SendFile("./api/docs/swagger.json")
	}
}

func (r *Router) Entrypoint() fiber.Handler {
	return swaggo.New(swaggo.Config{
		Title:              "Swagger UI",
		Layout:             "StandaloneLayout",
		URL:                "/docs/swagger.json",
		DeepLinking:        true,
		ShowMutatedRequest: true,
		Plugins: []template.JS{
			template.JS("SwaggerUIBundle.plugins.DownloadUrl"),
		},
		Presets: []template.JS{
			template.JS("SwaggerUIBundle.presets.apis"),
			template.JS("SwaggerUIStandalonePreset"),
		},
	})
}
