package traefik

import (
	"github.com/layer5io/meshery-adapter-library/status"
)

func (mesh *Mesh) applyCustomOperation(namespace string, manifest string, isDel bool) (string, error) {
	st := status.Starting

	err := mesh.applyManifest([]byte(manifest), isDel, namespace)
	if err != nil {
		return st, ErrCustomOperation(err)
	}

	return status.Completed, nil
}
