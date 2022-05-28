package traefik

import (
	"sync"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/status"
	mesherykube "github.com/layer5io/meshkit/utils/kubernetes"
)

func (mesh *Mesh) installSampleApp(namespace string, del bool, templates []adapter.Template, kubeconfigs []string) (string, error) {
	st := status.Installing

	if del {
		st = status.Removing
	}

	for _, template := range templates {
		err := mesh.applyManifest([]byte(template.String()), del, namespace, kubeconfigs)
		if err != nil {
			return st, ErrSampleApp(err)
		}
	}

	return status.Installed, nil
}

func (mesh *Mesh) applyManifest(contents []byte, isDel bool, namespace string, kubeconfigs []string) error {

	var wg sync.WaitGroup
	var errs []error
	for _, k8sconfig := range kubeconfigs {
		wg.Add(1)
		go func(k8sconfig string) {
			defer wg.Done()
			kClient, err := mesherykube.New([]byte(k8sconfig))
			if err != nil {
				errs = append(errs, err)
				return
			}
			err = kClient.ApplyManifest(contents, mesherykube.ApplyOptions{
				Namespace: namespace,
				Update:    true,
				Delete:    isDel,
			})
			if err != nil {
				errs = append(errs, err)
			}
		}(k8sconfig)
	}
	wg.Wait()
	if len(errs) != 0 {
		return mergeErrors(errs)
	}
	return nil
}
