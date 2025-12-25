package series

import (
    "encoding/json"

    "github.com/go-playground/validator"
)

type MetricDefine struct {
    Name      string   `json:"name" validate:"required"`
    Help      string   `json:"help" validate:"required"`
    Type      string   `json:"type" validate:"required"`
    LabelKeys []string `json:"label_keys" validate:"required"`
    ValueKey  string   `json:"value_key" validate:"required"`
}

func ValidateMetricDefine(jsonStr string) error {
    var m MetricDefine
    err := json.Unmarshal([]byte(jsonStr), &m)
    if err != nil {
        return err
    }

    validate := validator.New()
    return validate.Struct(m)
}

func ParseMetricDefine(jsonStr string) (*MetricDefine, error) {
    var m MetricDefine
    err := json.Unmarshal([]byte(jsonStr), &m)
    if err != nil {
        return nil, err
    }
    return &m, nil
}
