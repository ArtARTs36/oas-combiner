package internal

import (
	"bytes"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Spec struct {
	OpenAPI    any                       `yaml:"openapi"`
	Info       any                       `yaml:"info"`
	Servers    []any                     `yaml:"servers"`
	Paths      map[string]any            `yaml:"paths"`
	Components map[string]map[string]any `yaml:"components"`

	Include []Include `yaml:"include,omitempty"`
}

type Include struct {
	Ref string `yaml:"$ref"`
}

func LoadSpec(path string) (*Spec, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}

	var spec Spec

	if err = yaml.Unmarshal(content, &spec); err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	return &spec, nil
}

func MarshalSpec(spec *Spec) ([]byte, error) {
	buf := &bytes.Buffer{}

	enc := yaml.NewEncoder(buf)
	enc.SetIndent(1)

	err := enc.Encode(spec)
	if err != nil {
		return nil, fmt.Errorf("encode: %w", err)
	}

	return buf.Bytes(), nil
}
