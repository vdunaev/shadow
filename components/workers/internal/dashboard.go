package internal

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/workers/internal/handlers"
)

func (c *Component) DashboardTemplates() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "templates",
	}
}

func (c *Component) DashboardMenu() dashboard.Menu {
	routes := c.DashboardRoutes()

	return dashboard.NewMenuWithRoute("Workers", routes[1], "tasks", nil, nil)
}

func (c *Component) DashboardRoutes() []dashboard.Route {
	if c.routes == nil {
		c.routes = []dashboard.Route{
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet},
				"/"+c.Name()+"/assets/*filepath",
				&assetfs.AssetFS{
					Asset:     Asset,
					AssetDir:  AssetDir,
					AssetInfo: AssetInfo,
					Prefix:    "assets",
				},
				"",
				false),
			dashboard.NewRoute(
				c.Name(),
				[]string{http.MethodGet, http.MethodPost},
				"/"+c.Name()+"/",
				&handlers.ManagerHandler{
					Application: c.application,
					Component:   c,
				},
				"",
				true),
		}
	}

	return c.routes
}
