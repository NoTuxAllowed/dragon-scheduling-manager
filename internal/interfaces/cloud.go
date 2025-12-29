package interfaces

import "context"

type CloudProvider interface {
	Provision(ctx context.Context) error
	Check(ctx context.Context) error
	Terminate(ctx context.Context) error
	StartInstances(ctx context.Context) error
	StopInstances(ctx context.Context) error
	RestartInstances(ctx context.Context) error
}
