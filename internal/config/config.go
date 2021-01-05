package config

import (
	"path"

	"github.com/layer5io/meshery-adapter-library/common"
	"github.com/layer5io/meshery-adapter-library/config"
	configprovider "github.com/layer5io/meshery-adapter-library/config/provider"
	"github.com/layer5io/meshery-adapter-library/status"
	"github.com/layer5io/meshkit/utils"
)

const (
	// TraefikMeshOperation is the default name for the install
	// and uninstall commands on the traefik mesh
	TraefikMeshOperation = "traefik-mesh"
)

var (
	configRootPath = path.Join(utils.GetHome(), ".meshery")

	// Config is the collection of ServerConfig, MeshConfig and ProviderConfig
	Config = configprovider.Options{
		ServerConfig:   ServerConfig,
		MeshSpec:       MeshSpec,
		ProviderConfig: ProviderConfig,
		Operations:     Operations,
	}

	// ServerConfig is the configuration for the gRPC server
	ServerConfig = map[string]string{
		"name":    "traefik-mesh-adapter",
		"port":    "10006",
		"version": "v1.0.0",
	}

	// MeshSpec is the spec for the service mesh associated with this adapter
	MeshSpec = map[string]string{
		"name":     "traefik mesh",
		"status":   status.None,
		"traceurl": status.None,
		"version":  status.None,
	}

	// ProviderConfig is the config for the configuration provider
	ProviderConfig = map[string]string{
		configprovider.FilePath: configRootPath,
		configprovider.FileType: "yaml",
		configprovider.FileName: "traefik-mesh",
	}

	// KubeConfig - Controlling the kubeconfig lifecycle with viper
	KubeConfig = map[string]string{
		configprovider.FilePath: configRootPath,
		configprovider.FileType: "yaml",
		configprovider.FileName: "kubeconfig",
	}

	// Operations represents the set of valid operations that are available
	// to the adapter
	Operations = getOperations(common.Operations)
)

// New creates a new config instance
func New(provider string) (config.Handler, error) {
	// Config provider
	switch provider {
	case configprovider.ViperKey:
		return configprovider.NewViper(Config)
	case configprovider.InMemKey:
		return configprovider.NewInMem(Config)
	}

	return nil, ErrEmptyConfig
}

// NewKubeconfigBuilder returns a config handler based on the provider
//
// Valid prividers are "viper" and "in-mem"
func NewKubeconfigBuilder(provider string) (config.Handler, error) {
	opts := configprovider.Options{}
	opts.ProviderConfig = KubeConfig

	// Config provider
	switch provider {
	case configprovider.ViperKey:
		return configprovider.NewViper(opts)
	case configprovider.InMemKey:
		return configprovider.NewInMem(opts)
	}
	return nil, ErrEmptyConfig
}

// RootPath returns the config root path for the adapter
func RootPath() string {
	return configRootPath
}
