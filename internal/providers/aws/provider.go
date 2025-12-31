package aws

import (
	"context"
	"fmt"

	"github.com/NoTuxAllowed/dragon-scheduler/internal/interfaces"
)

type Provider struct {
	Config *AwsCloudSpec
}

// Ensure Provider implements CloudProvider
var _ interfaces.CloudProvider = (*Provider)(nil)

func (p *Provider) Provision(ctx context.Context) error {
	fmt.Printf("[AWS] Provisioning instance with type %s (Spot: %v)\n", p.Config.InstanceType, p.Config.Spot)
	return nil
}
func (p *Provider) Check(ctx context.Context) error {
	fmt.Printf("[AWS] Checking instance with type %s (Spot: %v)\n", p.Config.InstanceType, p.Config.Spot)
	return nil
}

func (p *Provider) Terminate(ctx context.Context) error {
	fmt.Printf("[AWS] Terminating instance\n")
	return nil
}

func (p *Provider) StartInstances(ctx context.Context) error {
	fmt.Printf("[AWS] Starting instances\n")
	return nil
}

func (p *Provider) StopInstances(ctx context.Context) error {
	fmt.Printf("[AWS] Stopping instances\n")
	return nil
}

func (p *Provider) RestartInstances(ctx context.Context) error {
	fmt.Printf("[AWS] Restarting instances\n")
	return nil
}
