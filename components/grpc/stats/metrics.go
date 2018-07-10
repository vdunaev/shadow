package stats

import (
	"context"
	"fmt"

	"github.com/kihamo/shadow/components/grpc"
	"github.com/kihamo/snitch"
	"google.golang.org/grpc/stats"
	"google.golang.org/grpc/status"
)

const (
	MetaDataClientNameKey = "user-agent"

	// GRPC specific
	MetricNameHandledTotal  = grpc.ComponentName + "_handled_total"
	MetricNameReceivedTotal = grpc.ComponentName + "_received_total"
	MetricNameSentTotal     = grpc.ComponentName + "_sent_total"
	MetricNameStartedTotal  = grpc.ComponentName + "_started_total"

	// For all requests
	MetricNameResponseTimeSeconds = grpc.ComponentName + "_response_time_seconds"
	MetricNameRequestSizeBytes    = grpc.ComponentName + "_request_size_bytes"
	MetricNameResponseSizeBytes   = grpc.ComponentName + "_response_size_bytes"
	MetricNameRequestsTotal       = grpc.ComponentName + "_requests_total"
)

var (
	MetricHandledTotal  = snitch.NewCounter(MetricNameHandledTotal, "GRPC handled requests total")
	MetricReceivedTotal = snitch.NewCounter(MetricNameReceivedTotal, "GRPC received requests total")
	MetricSentTotal     = snitch.NewCounter(MetricNameSentTotal, "GRPC sent responses total")
	MetricStartedTotal  = snitch.NewCounter(MetricNameStartedTotal, "GRPC started requests total")

	MetricResponseTimeSeconds = snitch.NewTimer(MetricNameResponseTimeSeconds, "GRPC response time in total")
	MetricRequestSizeBytes    = snitch.NewHistogram(MetricNameRequestSizeBytes, "GRPC size of requests in bytes")
	MetricResponseSizeBytes   = snitch.NewHistogram(MetricNameResponseSizeBytes, "GRPC size of responses in bytes")
	MetricRequestsTotal       = snitch.NewCounter(MetricNameRequestsTotal, "GRPC requests total")
)

type MetricsHandler struct {
	Handler
}

func NewMetricHandler() *MetricsHandler {
	return &MetricsHandler{}
}

func (h *MetricsHandler) HandleRPC(ctx context.Context, stat stats.RPCStats) {
	ctxValue := h.RPCValueFromContext(ctx)

	switch s := stat.(type) {
	case *stats.Begin:
		//fmt.Println("HandleRPC::Begin")

		MetricStartedTotal.With(
			"grpc_service", ctxValue.Service,
			"grpc_method", ctxValue.Method,
			"grpc_type", ctxValue.Type,
			"client_name", ctxValue.ClientName).Inc()

		MetricRequestsTotal.With(
			"handler", fmt.Sprintf("%s/%s", ctxValue.Service, ctxValue.Method),
			"protocol", grpc.ProtocolGRPC,
			"client_name", ctxValue.ClientName).Inc()

	case *stats.End:
		//fmt.Println("HandleRPC::End")

		if !s.IsClient() {
			code, _ := status.FromError(s.Error)
			MetricHandledTotal.With(
				"grpc_service", ctxValue.Service,
				"grpc_method", ctxValue.Method,
				"grpc_type", ctxValue.Type,
				"client_name", ctxValue.ClientName,
				"grpc_code", CodeAsString(code.Code())).Inc()

			st := grpc.StatusOK
			if s.Error != nil {
				st = grpc.StatusError
			}
			MetricResponseTimeSeconds.With(
				"handler", fmt.Sprintf("%s/%s", ctxValue.Service, ctxValue.Method),
				"protocol", grpc.ProtocolGRPC,
				"client_name", ctxValue.ClientName,
				"status", st).UpdateSince(s.BeginTime)
		}

	case *stats.InPayload:
		//fmt.Println("HandleRPC::InPayload")

		if !s.IsClient() {
			MetricReceivedTotal.With(
				"grpc_service", ctxValue.Service,
				"grpc_method", ctxValue.Method,
				"grpc_type", ctxValue.Type,
				"client_name", ctxValue.ClientName).Inc()

			MetricRequestSizeBytes.With(
				"handler", fmt.Sprintf("%s/%s", ctxValue.Service, ctxValue.Method),
				"protocol", grpc.ProtocolGRPC,
				"client_name", ctxValue.ClientName,
			).Add(float64(s.WireLength))
		}

	case *stats.OutPayload:
		//fmt.Println("HandleRPC::OutPayload")

		if !s.IsClient() {
			MetricSentTotal.With(
				"grpc_service", ctxValue.Service,
				"grpc_method", ctxValue.Method,
				"grpc_type", ctxValue.Type,
				"client_name", ctxValue.ClientName).Inc()

			MetricResponseSizeBytes.With(
				"handler", fmt.Sprintf("%s/%s", ctxValue.Service, ctxValue.Method),
				"protocol", grpc.ProtocolGRPC,
				"client_name", ctxValue.ClientName,
				"status", grpc.StatusOK).Add(float64(s.WireLength))
		}

	case *stats.InTrailer:
		//fmt.Println("HandleRPC::InTrailer")

	case *stats.OutTrailer:
		//fmt.Println("HandleRPC::OutTrailer")

	case *stats.InHeader:
		//fmt.Println("HandleRPC::InHeader")

	case *stats.OutHeader:
		//fmt.Println("HandleRPC::OutHeader")

	default:
		//fmt.Println("HandleRPC::default")
	}
}