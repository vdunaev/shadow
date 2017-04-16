package workers

import (
	"github.com/kihamo/go-workers/dispatcher"
	"github.com/kihamo/go-workers/worker"
	"github.com/kihamo/snitch"
)

const (
	MetricListenersTotal = ComponentName + ".listeners.total"
	MetricListenersTasks = ComponentName + ".listeners.tasks"
	MetricWorkersTotal   = ComponentName + ".workers.total"
	MetricWorkersStatus  = ComponentName + ".workers.status"
	MetricTasksTotal     = ComponentName + ".tasks.total"
	MetricTasksStatus    = ComponentName + ".tasks.status"
)

var (
	metricListenersTotal snitch.Gauge
	metricListenersTasks snitch.Gauge

	metricWorkersTotal        snitch.Counter
	metricWorkerStatusWait    snitch.Gauge
	metricWorkerStatusProcess snitch.Gauge
	metricWorkerStatusBusy    snitch.Gauge

	metricTasksTotal               snitch.Counter
	metricTasksStatusWait          snitch.Counter
	metricTasksStatusProcess       snitch.Counter
	metricTasksStatusSuccess       snitch.Counter
	metricTasksStatusFail          snitch.Counter
	metricTasksStatusFailByTimeout snitch.Counter
	metricTasksStatusKill          snitch.Counter
	metricTasksStatusRepeatWait    snitch.Counter
)

type metricsCollector struct {
	dispatcher *dispatcher.Dispatcher
}

func (c *metricsCollector) Describe(ch chan<- *snitch.Description) {
	ch <- metricListenersTotal.Description()
	ch <- metricListenersTasks.Description()

	ch <- metricWorkersTotal.Description()
	ch <- metricWorkerStatusWait.Description()
	ch <- metricWorkerStatusProcess.Description()
	ch <- metricWorkerStatusBusy.Description()

	ch <- metricTasksTotal.Description()
	ch <- metricTasksStatusWait.Description()
	ch <- metricTasksStatusProcess.Description()
	ch <- metricTasksStatusSuccess.Description()
	ch <- metricTasksStatusFail.Description()
	ch <- metricTasksStatusFailByTimeout.Description()
	ch <- metricTasksStatusKill.Description()
	ch <- metricTasksStatusRepeatWait.Description()
}

func (c *metricsCollector) Collect(ch chan<- snitch.Metric) {
	metricListenersTotal.Set(float64(len(c.dispatcher.GetListeners())))
	metricListenersTasks.Set(float64(len(c.dispatcher.GetListenersTasks())))

	var (
		workerStatusWait    float64
		workerStatusProcess float64
		workerStatusBusy    float64
	)

	for _, w := range c.dispatcher.GetWorkers().GetItems() {
		switch w.GetStatus() {
		case worker.WorkerStatusWait:
			workerStatusWait++
		case worker.WorkerStatusProcess:
			workerStatusProcess++
		case worker.WorkerStatusBusy:
			workerStatusBusy++
		}
	}

	metricWorkerStatusWait.Set(workerStatusWait)
	metricWorkerStatusProcess.Set(workerStatusProcess)
	metricWorkerStatusBusy.Set(workerStatusBusy)

	ch <- metricListenersTotal
	ch <- metricListenersTasks

	ch <- metricWorkersTotal
	ch <- metricWorkerStatusWait
	ch <- metricWorkerStatusProcess
	ch <- metricWorkerStatusBusy

	ch <- metricTasksTotal
	ch <- metricTasksStatusWait
	ch <- metricTasksStatusProcess
	ch <- metricTasksStatusSuccess
	ch <- metricTasksStatusFail
	ch <- metricTasksStatusFailByTimeout
	ch <- metricTasksStatusKill
	ch <- metricTasksStatusRepeatWait
}

func (c *Component) Metrics() snitch.Collector {
	metricListenersTotal = snitch.NewGauge(MetricListenersTotal)
	metricListenersTasks = snitch.NewGauge(MetricListenersTasks)

	metricWorkersTotal = snitch.NewCounter(MetricWorkersTotal)
	metricWorkerStatusWait = snitch.NewGauge(MetricWorkersStatus, "status", "wait")
	metricWorkerStatusProcess = snitch.NewGauge(MetricWorkersStatus, "status", "process")
	metricWorkerStatusBusy = snitch.NewGauge(MetricWorkersStatus, "status", "busy")

	metricTasksTotal = snitch.NewCounter(MetricTasksTotal)
	metricTasksStatusWait = snitch.NewCounter(MetricTasksStatus, "status", "wait")
	metricTasksStatusProcess = snitch.NewCounter(MetricTasksStatus, "status", "process")
	metricTasksStatusSuccess = snitch.NewCounter(MetricTasksStatus, "status", "success")
	metricTasksStatusFail = snitch.NewCounter(MetricTasksStatus, "status", "fail")
	metricTasksStatusFailByTimeout = snitch.NewCounter(MetricTasksStatus, "status", "fail-by-timeout")
	metricTasksStatusKill = snitch.NewCounter(MetricTasksStatus, "status", "kill")
	metricTasksStatusRepeatWait = snitch.NewCounter(MetricTasksStatus, "status", "repeat-wait")

	return &metricsCollector{
		dispatcher: c.dispatcher,
	}
}
