package azure

import (
	"context"
	"fmt"

	"github.com/NoTuxAllowed/dragon-scheduler/internal/interfaces"
)

type Provider struct {
	Config *AzureCloudSpec
}

// Ensure Provider implements CloudProvider
var _ interfaces.CloudProvider = (*Provider)(nil)

func (p *Provider) Provision(ctx context.Context) error {
	fmt.Printf("[Azure] Provisioning instance with type %s (Spot: %v)\n", p.Config.InstanceType, p.Config.Spot)
	return nil
}
func (p *Provider) Check(ctx context.Context) error {
	fmt.Printf("[Azure] Checking instance with type %s (Spot: %v)\n", p.Config.InstanceType, p.Config.Spot)
	return nil
}

func (p *Provider) Terminate(ctx context.Context) error {
	fmt.Printf("[Azure] Terminating instance\n")
	return nil
}

func (p *Provider) StartInstances(ctx context.Context) error {
	fmt.Printf("[Azure] Starting instances\n")
	return nil
}

func (p *Provider) StopInstances(ctx context.Context) error {
	fmt.Printf("[Azure] Stopping instances\n")
	return nil
}

func (p *Provider) RestartInstances(ctx context.Context) error {
	fmt.Printf("[Azure] Restarting instances\n")
	return nil
}
