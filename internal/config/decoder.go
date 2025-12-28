package config

import (
	"fmt"
	"reflect"

	"sigs.k8s.io/yaml"
)

type Object interface {
	GetKind() string
	GetName() string
}

func (m ConfigManifest[T]) GetKind() string { return m.Kind }

var registry = make(map[string]reflect.Type)
var registryCloudSpec = make(map[string]reflect.Type)

func Register[T any](kind string) {
	registry[kind] = reflect.TypeOf(ConfigManifest[T]{})
}

func RegisterCloudSpec[T any](kind string) {
	registryCloudSpec[kind] = reflect.TypeOf(ConfigManifest[T]{})
}

func LoadManifest(data []byte) (Object, error) {
	var peek struct {
		Kind string `json:"kind"`
	}
	if err := yaml.Unmarshal(data, &peek); err != nil {
		return nil, fmt.Errorf("failed to peek kind: %w", err)
	}
	combinedType, ok := registry[peek.Kind]
	if !ok {
		return nil, fmt.Errorf("unknown kind: %s", peek.Kind)
	}
	
	ptr := reflect.New(combinedType).Interface()
	if err := yaml.Unmarshal(data, ptr); err != nil {
		return nil, fmt.Errorf("failed to unmarshal into %s: %w", peek.Kind, err)
	}
	return reflect.ValueOf(ptr).Elem().Interface().(Object), nil
}
