package config

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type ConfigManifest struct {
	Version  string
	Kind     string
	Metadata ObjectMetadata
	Spec     any
}

type ObjectMetadata struct {
	Name string
}

type ClusterSpecV1 struct {
	CheckTimer    string
	CloudProvider string
	Spec          any
}

func (c *ClusterSpecV1) UnmarshalJSON(data []byte) error {
	type tempSpec struct {
		CheckTimer    string          `json:"checkTimer"`
		CloudProvider string          `json:"cloudProvider"`
		Spec          json.RawMessage `json:"spec"`
	}
	var temp tempSpec
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	c.CheckTimer = temp.CheckTimer
	c.CloudProvider = temp.CloudProvider

	specType, ok := cloudProviders[c.CloudProvider]
	if !ok {
		return fmt.Errorf("unknown cloud provider: %s", c.CloudProvider)
	}

	specPtr := reflect.New(specType).Interface()

	if err := json.Unmarshal(temp.Spec, specPtr); err != nil {
		return fmt.Errorf("failed to unmarshal cloud spec for %s: %w", c.CloudProvider, err)
	}

	c.Spec = specPtr

	return nil
}

type Metadata struct {
	Name string `json:"name"`
}

var cloudProviders = make(map[string]reflect.Type)
var cloudFactories = make(map[string]func(any) (any, error))

func RegisterCloudProvider[T any](name string, factory func(spec *T) (any, error)) {
	cloudProviders[name] = reflect.TypeOf(*new(T))
	cloudFactories[name] = func(rawSpec any) (any, error) {
		s, ok := rawSpec.(*T)
		if !ok {
			return nil, fmt.Errorf("invalid spec type for provider %s", name)
		}
		return factory(s)
	}
}

func GetProviderFactory(name string) (func(any) (any, error), bool) {
	f, ok := cloudFactories[name]
	return f, ok
}

func init() {
	Register[ClusterSpecV1]("Cluster")
}
