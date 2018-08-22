package main

import (
	"context"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/defaults"
	"github.com/containerd/containerd/namespaces"
)

const (
	fluxDefaultNamespace = "flux"
	defaultRuntime       = "io.containerd.runc.v1"
)

// GetContext returns a new Context under a default Containerd namespace.
func GetContext() context.Context {
	return namespaces.WithNamespace(context.Background(), fluxDefaultNamespace)
}

// NewClient returns a new containerd Client.
func NewClient() (*containerd.Client, error) {
	return containerd.New(
		defaults.DefaultAddress,
		containerd.WithDefaultRuntime(defaultRuntime),
	)
}
