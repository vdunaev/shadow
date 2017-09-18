package dashboard

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/kihamo/gotypes"
	"github.com/kihamo/shadow/components/config"
	"github.com/kihamo/shadow/components/dashboard/auth"
	"github.com/kihamo/shadow/components/logger"
)

type Request struct {
	original *http.Request
}

func NewRequest(r *http.Request) *Request {
	return &Request{
		original: r,
	}
}

func (r *Request) Original() *http.Request {
	return r.original
}

func (r *Request) Context() context.Context {
	return r.original.Context()
}

func (r *Request) Config() *config.Component {
	return ConfigFromContext(r.Context())
}

func (r *Request) Logger() logger.Logger {
	return LoggerFromContext(r.Context())
}

func (r *Request) Render() *Renderer {
	return RenderFromContext(r.Context())
}

func (r *Request) Panic() *PanicError {
	return PanicFromContext(r.Context())
}

func (r *Request) Session() *Session {
	return SessionFromContext(r.Context())
}

func (r *Request) User() *auth.User {
	user := &auth.User{}
	r.Session().GetObject(SessionUser, user)
	return user
}

func (r *Request) URL() *url.URL {
	return r.original.URL
}

func (r *Request) IsGet() bool {
	return r.original.Method == http.MethodGet
}

func (r *Request) IsPost() bool {
	return r.original.Method == http.MethodPost
}

func (r *Request) IsAjax() bool {
	return r.original.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

func (r *Request) DecodeJSON(j interface{}) error {
	decoder := json.NewDecoder(r.original.Body)

	var in interface{}
	err := decoder.Decode(&in)

	if err != nil {
		return err
	}

	converter := gotypes.NewConverter(in, &j)

	if !converter.Valid() {
		return errors.New("Convert failed")
	}

	return nil
}
