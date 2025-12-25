# Boring | [English](./README.en.md)
**Boring** 是基于 Prometheus 的轻量级告警系统。可统一管理分散在不同地方的 Prometheus 服务，并支持定时SQL采集

# 功能项
- 管理多 Prometheus 服务配置
- 图形化配置告警规则
- 定时采集 SQL 数据

# 架构图
![架构图](doc/design.png)

- **boring-server**: 管理端核心程序，实现定时任务、告警配置、数据采集配置等
- **boring-agent**: 部署在 Prometheus 节点上，接收 **boring-server** 下发的配置并写入文件
- **boring-jobworker**: 接收定时器下发的采集任务，并将采集数据推送到指定 Prometheus

# 部署运行

### 编译后端
```
cd ./server && ./build.sh
```

### 示例 `conf.yaml`:
```yaml
server:
  port: 7832
  host: 0.0.0.0

db:
 type: sqlite
 dsn: file:./boring.db?cache=shared&_pragma=foreign_keys(1)

job_worker_address: http://0.0.0.0:7855

job_worker:
  port: 7855
  host: 0.0.0.0
```

### 1. 启动 boring-agent
```
# 在部署 Prometheus 服务的主机上
./boring-agent --port 7767 --web.config.file ~/path/web-config.yml --config.file ~/path/prometheus.yml --rule.file ~/path/rule.yml
```

### 2. 启动 boring-jobworker
```
# 可使用 nginx 反向代理，部署多个实例
./boring-jobworker --config ./conf.yaml
```

### 3. 启动 boring-server
```
./boring-server --config ./conf.yaml
```

### 4. 运行前端
```
cd ./web
npm install
npm run dev
```

# 使用说明

### 管理 Prometheus 实例
- 在管理界面新增或修改 Prometheus 实例，填写访问地址、控制地址(**boring-agent** 监听地址)、用户名/密码等信息
- 规则通过告警规则页面设置，在提交后会同步到对应的 Prometheus 实例

![](doc/prom.png)

### 管理告警规则
告警规则使用 PromQL 标准语法
![](doc/rule.png)

### 创建 SQL 采集任务
指标配置示例
```
{
    "name": "alert_type_count", # 存入 prometheus 的指标名
    "help": "test",
    "type": "GAUGE",
    "label_keys": [             # 匹配 SQl 语句采集字段，并设置为指标标签
        "name",
        "code"
    ],
    "value_key":"count"         # 匹配 SQl 数量值，并设置为指标值
}
```

![](doc/sql.png)


# CHANGELOG
[CHANGELOG.md](./CHANGELOG.md)
