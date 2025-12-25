package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
)

// DBConfig 描述 db 段配置
type DBConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"`
}

// ServerConfig 描述 server 段配置
type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

type JobWorkerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

// Config 对应整个 conf.yaml
type Config struct {
	DB               DBConfig        `yaml:"db"`
	Server           ServerConfig    `yaml:"server"`
	JobWorkerAddress string          `yaml:"job_worker_address"`
	JobWorkerConfig  JobWorkerConfig `yaml:"job_worker"`
}

var config Config

func GetConfig() *Config {
	return &config
}

// Load 从指定路径加载并解析 YAML 配置文件
func Load(path string) *Config {
	// 如果传入相对路径，统一转为绝对路径，方便调试排查
	if !filepath.IsAbs(path) {
		abs, err := filepath.Abs(path)
		if err == nil {
			path = abs
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	return nil
}

// PrintLoadedConfig 打印当前内存中的配置内容，方便诊断配置加载问题。
func PrintLoadedConfig() {
	out, err := yaml.Marshal(config)
	if err != nil {
		fmt.Printf("conf: failed to marshal config: %v\n", err)
		return
	}
	fmt.Println("conf: loaded configuration:\n" + string(out))
}
