package promclient

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/snappy"
	"github.com/prometheus/prometheus/prompb"
)

type Pusher struct {
	PromURL  string
	Username string
	Password string
	Timeout  time.Duration
}

func NewPusher(promURL string, username string, password string, timeout time.Duration) *Pusher {
	return &Pusher{
		PromURL:  promURL,
		Username: username,
		Password: password,
		Timeout:  timeout,
	}
}

func (p *Pusher) doPush(series []prompb.TimeSeries) error {

	// TODO: 添加 HELPer 和 Type
	req := &prompb.WriteRequest{Timeseries: series}
	data, err := proto.Marshal(req)
	if err != nil {
		return err
	}
	compressed := snappy.Encode(nil, data)

	url := p.PromURL + "/api/v1/write"

	httpReq, err := http.NewRequest("POST", url, bytes.NewReader(compressed))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Encoding", "snappy")
	httpReq.Header.Set("Content-Type", "application/x-protobuf")
	httpReq.Header.Set("X-Prometheus-Remote-Write-Version", "0.1.0")

	if p.Username != "" && p.Password != "" {
		httpReq.SetBasicAuth(p.Username, p.Password)
	}

	client := &http.Client{Timeout: p.Timeout * time.Second}

	resp, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("server returned HTTP status %s", resp.Status)
	}

	return nil
}

func (p *Pusher) Push(series []prompb.TimeSeries) error {
	return p.doPush(series)
}
