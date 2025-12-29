package azure

import "github.com/NoTuxAllowed/dragon-scheduler/internal/config"

type AzureCloudSpec struct {
	Name                string
	Ephemeral           bool
	Scale               string
	Spot                bool
	InstanceType        string
	BackupInstanceTypes []string
}

func init() {
	config.RegisterCloudProvider("azure", func(spec *AzureCloudSpec) (any, error) {
		return &Provider{Config: spec}, nil
	})
}
