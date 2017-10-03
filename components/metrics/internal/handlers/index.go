package handlers

import (
	"sort"
	"time"

	"github.com/kihamo/shadow/components/dashboard"
	"github.com/kihamo/shadow/components/metrics"
)

type IndextHandler struct {
	dashboard.Handler

	Component metrics.Component
}

func (h *IndextHandler) ServeHTTP(_ *dashboard.Response, r *dashboard.Request) {
	var updated time.Time

	measures, err := h.Component.Registry().Gather()
	if err == nil {
		sort.Sort(measures)

		for _, measure := range measures {
			if measure.CreatedAt.After(updated) {
				updated = measure.CreatedAt
			}
		}
	}

	h.Render(r.Context(), h.Component.GetName(), "list", map[string]interface{}{
		"measures": measures,
		"error":    err,
		"updated":  updated,
	})
}