package dashboard

import (
	"context"
	"net/http"
)

type Handler struct {
	http.Handler
}

func (h *Handler) Redirect(url string, code int, w *Response, r *Request) {
	http.Redirect(w, r.Original(), url, code)
}

func (h *Handler) NotFound(w *Response, r *Request) {
	router := RouterFromContext(r.Context())
	if router == nil {
		panic("Router isn't set in context")
	}

	router.NotFoundServeHTTP(w, r.Original())
}

func (h *Handler) MethodNotAllowed(w *Response, r *Request) {
	router := RouterFromContext(r.Context())
	if router == nil {
		panic("Router isn't set in context")
	}

	router.MethodNotAllowedServeHTTP(w, r.Original())
}

func (h *Handler) InternalError(w *Response, r *Request, e error) {
	router := RouterFromContext(r.Context())
	if router == nil {
		panic("Router isn't set in context")
	}

	router.InternalErrorServeHTTP(w, r.Original(), e)
}

func (h *Handler) Render(ctx context.Context, view string, data map[string]interface{}) {
	render := RenderFromContext(ctx)
	if render == nil {
		panic("Render isn't set in context")
	}

	if err := render.Render(ResponseFromContext(ctx), ctx, TemplateNamespaceFromContext(ctx), view, data); err != nil {
		panic(err.Error())
	}
}

func (h *Handler) RenderLayout(ctx context.Context, view, layout string, data map[string]interface{}) {
	render := RenderFromContext(ctx)
	if render == nil {
		panic("Render isn't set in context")
	}

	if err := render.RenderLayout(ResponseFromContext(ctx), ctx, TemplateNamespaceFromContext(ctx), view, layout, data); err != nil {
		panic(err.Error())
	}
}
