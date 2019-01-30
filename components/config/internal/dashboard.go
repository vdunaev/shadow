package internal

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/config/internal/handlers"
	"github.com/kihamo/shadow/components/dashboard"
)

func (c *Component) DashboardTemplates() *assetfs.AssetFS {
	return dashboard.TemplatesFromAssetFS(c)
}

func (c *Component) DashboardMenu() dashboard.Menu {
	routes := c.DashboardRoutes()

	return dashboard.NewMenu("Configuration").WithRoute(routes[1]).WithIcon("cog")
}

func (c *Component) DashboardRoutes() []dashboard.Route {
	if c.routes == nil {
		c.routes = []dashboard.Route{
			dashboard.RouteFromAssetFS(c),
			dashboard.NewRoute("/"+c.Name()+"/", handlers.NewManagerHandler(c)).
				WithMethods([]string{http.MethodGet, http.MethodPost}).
				WithAuth(true),
		}
	}

	return c.routes
}

func (c *Component) DashboardTemplateFunctions() map[string]interface{} {
	return map[string]interface{}{
		"config": c.templateFunctionConfig,
	}
}

func (c *Component) DashboardMiddleware() []func(http.Handler) http.Handler {
	return []func(http.Handler) http.Handler{
		func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// save config in context
				if request := dashboard.RequestFromContext(r.Context()); request != nil {
					request.WithContext(config.ContextWithConfig(r.Context(), c))
					r = request.Original()
				}

				next.ServeHTTP(w, r)
			})
		},
	}
}

func (c *Component) templateFunctionConfig(key string, def ...interface{}) interface{} {
	var defValue interface{}

	if len(def) > 0 {
		defValue = def[0]
	}

	if c.Has(key) {
		return c.Get(key)
	}

	return defValue
}
