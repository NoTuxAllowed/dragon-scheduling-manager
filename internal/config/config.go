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
	// 1. Unmarshal into a temporary struct to get the CloudProvider and the raw Spec
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

	// 2. Look up the correct type in the registry
	specType, ok := cloudProviders[c.CloudProvider]
	if !ok {
		return fmt.Errorf("unknown cloud provider: %s", c.CloudProvider)
	}

	// 3. Create a new instance of that type
	specPtr := reflect.New(specType).Interface()

	// 4. Unmarshal the raw spec into that instance
	if err := json.Unmarshal(temp.Spec, specPtr); err != nil {
		return fmt.Errorf("failed to unmarshal cloud spec for %s: %w", c.CloudProvider, err)
	}

	// 5. Assign the value to c.Spec
	c.Spec = reflect.ValueOf(specPtr).Elem().Interface()

	return nil
}

type Metadata struct {
	Name string `json:"name"`
}

func init() {
	Register[ClusterSpecV1]("Cluster")
}