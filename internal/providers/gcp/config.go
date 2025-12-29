package gcp

import "github.com/NoTuxAllowed/dragon-scheduler/internal/config"

type GCPCloudSpec struct {
	Name                string
	Ephemeral           bool
	Scale               string
	Spot                bool
	InstanceType        string
	BackupInstanceTypes []string
}

func init() {
	config.RegisterCloudProvider("gcp", func(spec *GCPCloudSpec) (any, error) {
		return &Provider{Config: spec}, nil
	})
}
