-- t_alert_rule definition

CREATE TABLE t_alert_rule (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	"level" TEXT NOT NULL,
	"type" TEXT NOT NULL,
	"source" TEXT NOT NULL,
	promql_rule TEXT NOT NULL,
	content TEXT NOT NULL,
	"for" INTEGER NOT NULL,
	promql_query TEXT,
	custom_labels TEXT,
	notify_id TEXT,
	created_at DATETIME DEFAULT (CURRENT_TIMESTAMP),
	updated_at DATETIME DEFAULT (CURRENT_TIMESTAMP)
, enabled TEXT, prometheus_config_id INTEGER);

CREATE UNIQUE INDEX alert_rule_unique ON t_alert_rule(title, prometheus_config_id);


-- t_alert_type definition

CREATE TABLE t_alert_type (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  code TEXT NOT NULL,
  created_at DATETIME NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  updated_at DATETIME NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);

-- t_prometheus_config definition

CREATE TABLE t_prometheus_config (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  remark TEXT NOT NULL,
  address TEXT NOT NULL,
  username TEXT,
  password TEXT,
  ctrl_address TEXT NOT NULL,
  config TEXT,
  rule TEXT,
  enabled TEXT NOT NULL ,
  created_at DATETIME NOT NULL DEFAULT (CURRENT_TIMESTAMP),
  updated_at DATETIME NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);