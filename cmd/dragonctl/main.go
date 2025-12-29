package main

import (
	"context"
	"fmt"
	"os"

	"github.com/NoTuxAllowed/dragon-scheduler/internal/config"
	"github.com/NoTuxAllowed/dragon-scheduler/internal/interfaces"
	_ "github.com/NoTuxAllowed/dragon-scheduler/internal/providers/aws"
	_ "github.com/NoTuxAllowed/dragon-scheduler/internal/providers/azure"
	_ "github.com/NoTuxAllowed/dragon-scheduler/internal/providers/gcp"
)

func main() {
	fmt.Println("this is dragonctl")
	
	data, err := os.ReadFile("/home/bigpod/test.yaml")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	manifest, err := config.LoadManifest(data)
	if err != nil {
		fmt.Printf("Error loading manifest: %v\n", err)
		return
	}

	spec, ok := manifest.Spec.(config.ClusterSpecV1)
	if !ok {
		fmt.Printf("Unknown or unsupported spec type: %T\n", manifest.Spec)
		return
	}

	factory, ok := config.GetProviderFactory(spec.CloudProvider)
	if !ok {
		fmt.Printf("Unknown provider: %s\n", spec.CloudProvider)
		return
	}

	provider, err := factory(spec.Spec)
	if err != nil {
		fmt.Printf("Failed to create provider: %v\n", err)
		return
	}

	cloud, ok := provider.(interfaces.CloudProvider)
	if !ok {
		fmt.Printf("Provider %s does not implement CloudProvider\n", spec.CloudProvider)
		return
	}

	ctx := context.Background()
	if err := cloud.Provision(ctx); err != nil {
		fmt.Printf("Provision failed: %v\n", err)
		return
	}
	if err := cloud.Check(ctx); err != nil {
		fmt.Printf("Check failed: %v\n", err)
	}
	if err := cloud.Terminate(ctx); err != nil {
		fmt.Printf("Terminate failed: %v\n", err)
	}
}
