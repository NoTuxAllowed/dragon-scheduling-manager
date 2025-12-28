package config

type ConfigManifest[T any] struct {
	Version  string
	Kind     string
	Metadata ObjectMetadata
	Spec     T
}

type ObjectMetadata struct {
	Name string
}

type ClusterSpecV1[T any] struct {
	CheckTimer    string
	CloudProvider string
	Spec          T
}
type Metadata struct {
	Name string `json:"name"`
}

func init() {
	Register[ClusterSpecV1[any]]("aws")
}