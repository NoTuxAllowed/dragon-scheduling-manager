package config

import (
	"encoding/json"
	"fmt"
	"reflect"

	"sigs.k8s.io/yaml"
)

type Object interface {
	GetKind() string
	GetName() string
}

func (m *ConfigManifest) GetKind() string { return m.Kind }
func (m *ConfigManifest) GetName() string { return m.Metadata.Name }

var registry = make(map[string]reflect.Type)

func Register[T any](kind string) {
	registry[kind] = reflect.TypeOf((*T)(nil)).Elem()
}

func LoadManifest(data []byte) (*ConfigManifest, error) {
	var peek struct {
		Kind string `json:"kind"`
	}
	if err := yaml.Unmarshal(data, &peek); err != nil {
		return nil, fmt.Errorf("failed to peek kind: %w", err)
	}
	specType, ok := registry[peek.Kind]
	if !ok {
		return nil, fmt.Errorf("unknown kind: %s", peek.Kind)
	}

	type tempManifest struct {
		Version  string          `json:"version"`
		Kind     string          `json:"kind"`
		Metadata ObjectMetadata  `json:"metadata"`
		Spec     json.RawMessage `json:"spec"`
	}
	var tm tempManifest
	if err := yaml.Unmarshal(data, &tm); err != nil {
		return nil, fmt.Errorf("failed to unmarshal manifest structure: %w", err)
	}

	m := &ConfigManifest{
		Version:  tm.Version,
		Kind:     tm.Kind,
		Metadata: tm.Metadata,
	}

	specPtr := reflect.New(specType).Interface()
	if err := yaml.Unmarshal(tm.Spec, specPtr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal spec into %s: %w", peek.Kind, err)
	}
	m.Spec = reflect.ValueOf(specPtr).Elem().Interface()

	return m, nil
}
