package traefik

import (
	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/status"
	"github.com/layer5io/meshkit/utils"
	mesherykube "github.com/layer5io/meshkit/utils/kubernetes"
)

func (mesh *Mesh) installSampleApp(namespace string, del bool, templates []adapter.Template) (string, error) {
	st := status.Installing

	if del {
		st = status.Removing
	}

	for _, template := range templates {
		contents, err := utils.ReadFileSource(string(template))
		if err != nil {
			return st, ErrSampleApp(err)
		}

		err = mesh.applyManifest([]byte(contents), del, namespace)
		if err != nil {
			return st, ErrSampleApp(err)
		}
	}

	return status.Installed, nil
}

func (mesh *Mesh) applyManifest(contents []byte, isDel bool, namespace string) error {
	kclient, err := mesherykube.New(mesh.KubeClient, mesh.RestConfig)
	if err != nil {
		return err
	}

	err = kclient.ApplyManifest(contents, mesherykube.ApplyOptions{
		Namespace: namespace,
		Update:    true,
		Delete:    isDel,
	})

	if err != nil {
		return err
	}

	return nil
}
