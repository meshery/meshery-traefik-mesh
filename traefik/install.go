package traefik

import (
	"fmt"
	"sync"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/status"
	mesherykube "github.com/layer5io/meshkit/utils/kubernetes"
)

func (mesh *Mesh) installTraefikMesh(del bool, version, namespace string, kubeconfigs []string) (string, error) {
	mesh.Log.Debug(fmt.Sprintf("Requested install of version: %s", version))
	mesh.Log.Debug(fmt.Sprintf("Requested action is delete: %v", del))
	mesh.Log.Debug(fmt.Sprintf("Requested action is in namespace: %s", namespace))

	st := status.Installing
	if del {
		st = status.Removing
	}

	err := mesh.Config.GetObject(adapter.MeshSpecKey, mesh)
	if err != nil {
		return st, ErrMeshConfig(err)
	}

	err = mesh.applyHelmChart(del, version, namespace, kubeconfigs)
	if err != nil {
		return st, ErrApplyHelmChart(err)
	}

	st = status.Installed
	if del {
		st = status.Removed
	}

	return st, nil
}

func (mesh *Mesh) applyHelmChart(del bool, version, namespace string, kubeconfigs []string) error {
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
			repo := "https://helm.traefik.io/mesh"
			chart := "traefik-mesh"
			var act mesherykube.HelmChartAction
			if del {
				act = mesherykube.UNINSTALL
			} else {
				act = mesherykube.INSTALL
			}
			err = kClient.ApplyHelmChart(mesherykube.ApplyHelmChartConfig{
				ChartLocation: mesherykube.HelmChartLocation{
					Repository: repo,
					Chart:      chart,
					AppVersion: version,
				},
				Namespace:       namespace,
				Action:          act,
				CreateNamespace: true,
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
