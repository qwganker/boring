package series

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/prometheus/prometheus/prompb"
)

// 将 JSON 数据转换为 Prometheus TimeSeries 格式
type JsonConverter struct {
	rawJsonData  []byte
	metricDefine *MetricDefine
}

func NewJsonConverter(data []byte, metricDefine *MetricDefine) *JsonConverter {
	return &JsonConverter{
		rawJsonData:  data,
		metricDefine: metricDefine,
	}
}

func (s *JsonConverter) EncodeMetricMetadata() ([]prompb.MetricMetadata, error) {

	metadata := prompb.MetricMetadata{
		Type: prompb.MetricMetadata_MetricType(prompb.MetricMetadata_MetricType_value[s.metricDefine.Type]),
		Help: s.metricDefine.Help,
	}
	return []prompb.MetricMetadata{metadata}, nil
}

func (s *JsonConverter) EncodeTimeSeries() ([]prompb.TimeSeries, error) {
	var data []map[string]interface{}

	err := json.Unmarshal(s.rawJsonData, &data)
	if err != nil {
		return nil, err
	}

	series := []prompb.TimeSeries{}

	for _, item := range data {

		metricValue := 0.0

		labels := []prompb.Label{}
		labels = append(labels, prompb.Label{Name: "__name__", Value: s.metricDefine.Name})

		for key, value := range item {
			if key == s.metricDefine.ValueKey {
				metricValue = value.(float64)
			} else {
				for _, labelKey := range s.metricDefine.LabelKeys {
					if key == labelKey {
						labels = append(labels, prompb.Label{Name: key, Value: fmt.Sprintf("%v", value)})
					}
				}
			}
		}

		series = append(series, prompb.TimeSeries{
			Labels: labels,
			Samples: []prompb.Sample{
				{
					Value:     metricValue,
					Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
				},
			},
		})

	}

	return series, nil
}
