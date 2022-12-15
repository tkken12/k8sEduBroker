package query

import (
	"context"
	pClient "k8sEduBroker/monitoring/prometheus/client"
	"k8sEduBroker/util"
	"time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type QueryResult struct {
	Value      model.Value
	Warn       v1.Warnings
	Error      error
	Timeseries []Timeseries `json:"timeseries"`
}

type Timeseries struct {
	Key   model.Metric  `json:"key"`
	Value []interface{} `json:"value"` // idx[0] => epoch time(sec), idx[1] => query result(str)
}

func promAPI() v1.API { return v1.NewAPI(pClient.GetPrometheusClient()) }
func ContextTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(
		context.Background(),
		time.Duration(util.GetBrokerConf().PromTimeoutSecond)*time.Second)
}

func (result *QueryResult) QueryCall(query string) {

	ctx, cancel := ContextTimeout()
	defer cancel()

	val, warn, err := promAPI().Query(
		ctx, query, time.Now(), v1.WithTimeout(15*time.Second))

	result.Value = val
	result.Warn = warn
	result.Error = err

	return
}

func (result *QueryResult) TimeseriesCall(query string, startTime, endTime time.Time, period time.Duration) {

	ctx, cancel := ContextTimeout()
	defer cancel()

	val, warn, err := promAPI().QueryRange(
		ctx, query, v1.Range{
			Start: startTime,
			End:   endTime,
			Step:  period,
		},
	)

	result.Value = val
	result.Warn = warn
	result.Error = err

	return
}
