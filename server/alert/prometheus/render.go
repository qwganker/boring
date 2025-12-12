package prometheus

import (
	"bytes"
	"context"
	"fmt"

	"github.com/qwganker/boring/comm/constant"
	"github.com/qwganker/boring/comm/table"
	"github.com/qwganker/boring/storage"
	yaml "gopkg.in/yaml.v3"
)

type Annotations struct {
	Content string `yaml:"content,omitempty"`
	Level   string `yaml:"level,omitempty"`
}

type AlertRuleYaml struct {
	Title       string      `yaml:"alert"`
	For         string      `yaml:"for,omitempty"`
	Expr        string      `yaml:"expr,omitempty"`
	Annotations Annotations `yaml:"annotations,omitempty"`
}

type AlertGroupYaml struct {
	Name  string          `yaml:"name"`
	Rules []AlertRuleYaml `yaml:"rules"`
}

type AlertGroupsYaml struct {
	Groups []AlertGroupYaml `yaml:"groups"`
}

func RenderAlertRule(ctx context.Context, prometheusID int64) (string, error) {

	gormDB := storage.GetDBInstance()

	rules := []table.TAlertRule{}
	if err := gormDB.WithContext(ctx).Where(&table.TAlertRule{Enabled: constant.Enabled, PrometheusConfigID: prometheusID}, prometheusID).Find(&rules).Error; err != nil {
		return "", err
	}

	tmpRuleMap := make(map[string][]AlertRuleYaml)

	for _, r := range rules {
		if tmpRuleMap[r.Type] == nil {
			tmpRuleMap[r.Type] = []AlertRuleYaml{}
		}

		tmpRuleMap[r.Type] = append(tmpRuleMap[r.Type], AlertRuleYaml{
			Title: r.Title,
			For:   fmt.Sprintf("%ds", r.For),
			Expr:  r.PromQLRule,
			Annotations: Annotations{
				Level:   r.Level,
				Content: r.Content,
			},
		})
	}

	groupsYaml := AlertGroupsYaml{}

	for k, v := range tmpRuleMap {
		yamlGroups := AlertGroupYaml{
			Name:  k,
			Rules: []AlertRuleYaml{},
		}

		for _, rule := range v {
			yamlGroups.Rules = append(yamlGroups.Rules, rule)
		}

		groupsYaml.Groups = append(groupsYaml.Groups, yamlGroups)
	}

	ruleYaml := new(bytes.Buffer)

	encoder := yaml.NewEncoder(ruleYaml)
	err := encoder.Encode(&groupsYaml)
	if err != nil {
		return "", err
	}

	return ruleYaml.String(), nil
}
