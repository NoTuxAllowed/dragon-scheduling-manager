package aws

import "github.com/NoTuxAllowed/dragon-scheduler/internal/config"

type AwsCloudSpec struct {
	Name                string
	Ephemeral           bool
	Scale               string
	Spot                bool
	InstanceType        string
	BackupInstanceTypes []string
}

func init() {
	config.RegisterCloudProvider("aws", func(spec *AwsCloudSpec) (any, error) {
		return &Provider{Config: spec}, nil
	})
}
