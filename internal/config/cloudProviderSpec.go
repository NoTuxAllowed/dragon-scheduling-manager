package config

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
	RegisterCloudSpec[AwsCloudSpec]("aws")
	RegisterCloudSpec[AzureCloudSpec]("azure")
	RegisterCloudSpec[GCPCloudSpec]("gcp")
}
