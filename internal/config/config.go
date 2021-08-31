package config

import (
	"path"
	"strings"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/common"
	"github.com/layer5io/meshery-adapter-library/config"
	"github.com/layer5io/meshery-adapter-library/status"
	configprovider "github.com/layer5io/meshkit/config/provider"
	"github.com/layer5io/meshkit/utils"
	smp "github.com/layer5io/service-mesh-performance/spec"
)

const (
	// OAM metadata constants
	OAMAdapterNameMetadataKey       = "adapter.meshery.io/name"
	OAMComponentCategoryMetadataKey = "ui.meshery.io/category"
)

var (
	// TraefikMeshOperation is the default name for the install
	// and uninstall commands on the traefik mesh
	TraefikMeshOperation = strings.ToLower(smp.ServiceMesh_TRAEFIK_MESH.Enum().String())

	configRootPath = path.Join(utils.GetHome(), ".meshery")

	// ServerConfig is the configuration for the gRPC server
	ServerConfig = map[string]string{
		"name":     smp.ServiceMesh_TRAEFIK_MESH.Enum().String(),
		"port":     "10006",
		"type":     "adapter",
		"traceurl": status.None,
	}

	// MeshSpec is the spec for the service mesh associated with this adapter
	MeshSpec = map[string]string{
		"name":    smp.ServiceMesh_TRAEFIK_MESH.Enum().String(),
		"status":  status.None,
		"version": status.None,
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
func New(provider string) (h config.Handler, err error) {
	opts := configprovider.Options{
		FilePath: configRootPath,
		FileName: "traefik",
		FileType: "yaml",
	}
	// Config provider
	switch provider {
	case configprovider.ViperKey:
		h, err = configprovider.NewViper(opts)
		if err != nil {
			return nil, err
		}
	case configprovider.InMemKey:
		h, err = configprovider.NewInMem(opts)
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrEmptyConfig
	}

	// Setup Server config
	if err := h.SetObject(adapter.ServerKey, ServerConfig); err != nil {
		return nil, err
	}

	// setup Mesh config
	if err := h.SetObject(adapter.MeshSpecKey, MeshSpec); err != nil {
		return nil, err
	}

	// setup Operation Config
	if err := h.SetObject(adapter.OperationsKey, Operations); err != nil {
		return nil, err
	}

	return h, nil
}

// NewKubeconfigBuilder returns a config handler based on the provider
//
// Valid prividers are "viper" and "in-mem"
func NewKubeconfigBuilder(provider string) (config.Handler, error) {
	opts := configprovider.Options{
		FilePath: configRootPath,
		FileType: "yaml",
		FileName: "kubeconfig",
	}

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
