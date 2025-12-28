package config

import "reflect"

var cloudProviders = make(map[string]reflect.Type)

func RegisterCloudProvider[T any](name string) {
	cloudProviders[name] = reflect.TypeOf(*new(T))
}

type AwsCloudSpec struct {
	Name                string
	Ephemeral           bool
	Scale               string
	Spot                bool
	InstanceType        string
	BackupInstanceTypes []string
}

type AzureCloudSpec struct {
	Name                string
	Ephemeral           bool
	Scale               string
	Spot                bool
	InstanceType        string
	BackupInstanceTypes []string
}

type GCPCloudSpec struct {
	Name                string
	Ephemeral           bool
	Scale               string
	Spot                bool
	InstanceType        string
	BackupInstanceTypes []string
}

func init() {
	RegisterCloudProvider[AwsCloudSpec]("aws")
	RegisterCloudProvider[AzureCloudSpec]("azure")
	RegisterCloudProvider[GCPCloudSpec]("gcp")
}
