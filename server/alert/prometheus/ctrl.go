package prometheus

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qwganker/boring/comm/table"
)

type PrometheusCtrl struct {
	ctx *gin.Context
}

func NewPrometheusCtrl(c *gin.Context) *PrometheusCtrl {
	return &PrometheusCtrl{ctx: c}
}

func get(url, username, password string, timeoutSecond time.Duration, params io.Reader) error {
	req, err := http.NewRequest("GET", url, params)
	if err != nil {
		log.Printf("创建请求失败: %v", err)
		return err
	}

	if username != "" || password != "" {
		req.SetBasicAuth(username, password)
	}

	client := &http.Client{
		Timeout: timeoutSecond * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("发送请求失败: %v", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取响应失败: %v", err)
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("请求失败，状态码: %d, %s", resp.StatusCode, string(body))
	}

	return nil
}

func post(url, username, password string, timeoutSecond time.Duration, params io.Reader) error {
	req, err := http.NewRequest("POST", url, params)
	if err != nil {
		log.Printf("创建请求失败: %v", err)
		return err
	}

	if username != "" || password != "" {
		req.SetBasicAuth(username, password)
	}

	client := &http.Client{
		Timeout: timeoutSecond * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("发送请求失败: %v", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取响应失败: %v", err)
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("请求失败，状态码: %d, %s", resp.StatusCode, string(body))
	}

	return nil
}

func (p *PrometheusCtrl) CheckPrometheusStatus(config *table.TPrometheusConfig) error {
	if err := get(config.Address+"/query", config.Username, config.Password, 10, nil); err != nil {
		log.Printf("检查 Prometheus 运行状态失败: %v", err)
		return err
	}

	return nil
}

func (p *PrometheusCtrl) ReloadPrometheus(config *table.TPrometheusConfig) error {
	if err := post(config.Address+"/-/reload", config.Username, config.Password, 10, nil); err != nil {
		log.Printf("Reload Prometheus 失败: %v", err)
		return err
	}

	return nil
}

func (p *PrometheusCtrl) RewritePrometheusConfig(config *table.TPrometheusConfig) error {
	if err := post(config.CtrlAddress+"/-/rewrite_config", config.Username, config.Password, 10, bytes.NewBufferString(config.Config)); err != nil {
		log.Printf("Rewrite Prometheus Config 失败: %v", err)
		return err
	}

	return nil
}

func (p *PrometheusCtrl) RewritePrometheusRule(config *table.TPrometheusConfig) error {
	if err := post(config.CtrlAddress+"/-/rewrite_rule", config.Username, config.Password, 10, bytes.NewBufferString(config.Rule)); err != nil {
		log.Printf("Rewrite Prometheus Rule 失败: %v", err)
		return err
	}

	return nil
}
