package internal

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/alexedwards/scs"
	"github.com/alexedwards/scs/stores/memstore"
	"github.com/kihamo/shadow"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/dashboard/auth"
	"github.com/kihamo/shadow/components/dashboard/auth/providers/password"
	"github.com/kihamo/shadow/components/i18n"
	"github.com/kihamo/shadow/components/logging"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/gitlab"
	"github.com/markbates/goth/providers/gplus"
)

type Component struct {
	application shadow.Application
	config      config.Component
	logger      logging.Logger
	renderer    *Renderer
	session     *scs.Manager
	routes      []dashboard.Route
	router      *Router
	server      *http.Server
}

func (c *Component) Name() string {
	return dashboard.ComponentName
}

func (c *Component) Version() string {
	return dashboard.ComponentVersion
}

func (c *Component) Dependencies() []shadow.Dependency {
	return []shadow.Dependency{
		{
			Name:     config.ComponentName,
			Required: true,
		},
		{
			Name: i18n.ComponentName,
		},
		{
			Name: logging.ComponentName,
		},
	}
}

func (c *Component) Init(a shadow.Application) error {
	c.application = a
	c.config = a.GetComponent(config.ComponentName).(config.Component)

	return nil
}

func (c *Component) Run() (err error) {
	c.logger = logging.NewOrNop(c.Name(), c.application)

	if err := c.loadTemplates(); err != nil {
		return err
	}

	if err := c.loadMenu(); err != nil {
		return err
	}

	c.initSession()

	if err := c.initAuth(); err != nil {
		return err
	}

	if c.router, err = c.getServeMux(); err != nil {
		return err
	}

	addr := net.JoinHostPort(c.config.String(dashboard.ConfigHost), c.config.String(dashboard.ConfigPort))
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		c.logger.Errorf("Failed to listen [%d]: %s\n", os.Getpid(), err.Error())
		return err
	}

	c.logger.Info("Running service", map[string]interface{}{
		"addr": addr,
		"pid":  os.Getpid(),
	})

	c.server = &http.Server{
		Handler: c.router,
	}

	if err := c.server.Serve(lis); err != nil {
		c.logger.Errorf("Failed to serve [%d]: %s\n", os.Getpid(), err.Error())
		return err
	}

	return nil
}

func (c *Component) Shutdown() error {
	if c.server != nil {
		return c.server.Shutdown(context.Background())
	}

	return nil
}

func (c *Component) initAuth() (err error) {
	auth.ClearProviders()
	providers := []goth.Provider{}

	var baseURL *url.URL
	baseURLFromConfig := c.config.String(dashboard.ConfigOAuth2BaseURL)
	if baseURLFromConfig != "" {
		if baseURL, err = url.Parse(baseURLFromConfig); err != nil {
			return err
		}
	}

	if baseURL == nil {
		return fmt.Errorf("Base path for auth callbacks is empty")
	}

	baseURL.Path = strings.Trim(baseURL.Path, "/")
	pathCallbackTpl := "%s/" + strings.Trim(dashboard.AuthPath, "/") + "/%s/callback"

	if c.config.Bool(dashboard.ConfigAuthEnabled) {
		passwordRedirectURL := new(url.URL)
		*passwordRedirectURL = *baseURL
		passwordRedirectURL.Path = fmt.Sprintf(pathCallbackTpl, passwordRedirectURL.Path, "github")

		providers = append(providers, password.New(
			passwordRedirectURL.String(),
			map[string]string{
				c.config.String(dashboard.ConfigAuthUser): c.config.String(dashboard.ConfigAuthPassword),
			},
		))
	}

	if c.config.Bool(dashboard.ConfigOAuth2GithubEnabled) {
		githubRedirectURL := new(url.URL)
		*githubRedirectURL = *baseURL
		githubRedirectURL.Path = fmt.Sprintf(pathCallbackTpl, githubRedirectURL.Path, "github")

		providers = append(providers, github.NewCustomisedURL(
			c.config.String(dashboard.ConfigOAuth2GithubID),
			c.config.String(dashboard.ConfigOAuth2GithubSecret),
			githubRedirectURL.String(),
			github.AuthURL,
			github.TokenURL,
			github.ProfileURL,
			github.EmailURL,
			strings.Split(c.config.String(dashboard.ConfigOAuth2GithubScopes), ",")...,
		))
	}

	if c.config.Bool(dashboard.ConfigOAuth2GitlabEnabled) {
		gitlabRedirectURL := new(url.URL)
		*gitlabRedirectURL = *baseURL
		gitlabRedirectURL.Path = fmt.Sprintf(pathCallbackTpl, gitlabRedirectURL.Path, "gitlab")

		providers = append(providers, gitlab.NewCustomisedURL(
			c.config.String(dashboard.ConfigOAuth2GitlabID),
			c.config.String(dashboard.ConfigOAuth2GitlabSecret),
			gitlabRedirectURL.String(),
			c.config.String(dashboard.ConfigOAuth2GitlabAuthURL),
			c.config.String(dashboard.ConfigOAuth2GitlabTokenURL),
			c.config.String(dashboard.ConfigOAuth2GitlabProfileURL),
			strings.Split(c.config.String(dashboard.ConfigOAuth2GitlabScopes), ",")...,
		))
	}

	if c.config.Bool(dashboard.ConfigOAuth2GplusEnabled) {
		gplusRedirectURL := new(url.URL)
		*gplusRedirectURL = *baseURL
		gplusRedirectURL.Path = fmt.Sprintf(pathCallbackTpl, gplusRedirectURL.Path, "gplus")

		providers = append(providers, gplus.New(
			c.config.String(dashboard.ConfigOAuth2GplusID),
			c.config.String(dashboard.ConfigOAuth2GplusSecret),
			gplusRedirectURL.String(),
			strings.Split(c.config.String(dashboard.ConfigOAuth2GplusScopes), ",")...,
		))
	}

	auth.UseProviders(providers...)

	return nil
}

func (c *Component) initSession() {
	store := memstore.New(0)
	c.session = scs.NewManager(store)
	c.session.Name(c.config.String(dashboard.ConfigSessionCookieName))

	c.session.Domain(c.config.String(dashboard.ConfigSessionDomain))
	c.session.HttpOnly(c.config.Bool(dashboard.ConfigSessionHttpOnly))
	c.session.IdleTimeout(c.config.Duration(dashboard.ConfigSessionIdleTimeout))
	c.session.Lifetime(c.config.Duration(dashboard.ConfigSessionLifetime))
	c.session.Path(c.config.String(dashboard.ConfigSessionPath))
	c.session.Persist(c.config.Bool(dashboard.ConfigSessionPersist))
	c.session.Secure(c.config.Bool(dashboard.ConfigSessionSecure))
}

func (c *Component) Renderer() dashboard.Renderer {
	return c.renderer
}
