package config

import (
	"strings"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/meshes"
	smp "github.com/layer5io/service-mesh-performance/spec"
)

var (
	TraefikOperation          = strings.ToLower(smp.ServiceMesh_TRAEFIK_MESH.Enum().String())
	TraefikBookStoreOperation = "traefik_bookstore_app"
	ServiceName               = "service_name"
)

func getOperations(dev adapter.Operations) adapter.Operations {
	versions, _ := getLatestReleaseNames(3)

	dev[TraefikMeshOperation] = &adapter.Operation{
		Type:        int32(meshes.OpCategory_INSTALL),
		Description: "Traefik Mesh",
		Versions:    versions,
		Templates: []adapter.Template{
			"templates/traefik-mesh.yaml",
		},
		AdditionalProperties: map[string]string{},
	}

	return dev
}
