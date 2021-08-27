package traefik

import (
	"context"
	"fmt"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/status"
	mesherykube "github.com/layer5io/meshkit/utils/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (mesh *Mesh) installSampleApp(namespace string, del bool, templates []adapter.Template) (string, error) {
	st := status.Installing

	if del {
		st = status.Removing
	}

	for _, template := range templates {
		err := mesh.applyManifest([]byte(template.String()), del, namespace)
		if err != nil {
			return st, ErrSampleApp(err)
		}
	}

	return status.Installed, nil
}

func (mesh *Mesh) applyManifest(contents []byte, isDel bool, namespace string) error {
	kclient := mesh.MesheryKubeclient

	err := kclient.ApplyManifest(contents, mesherykube.ApplyOptions{
		Namespace: namespace,
		Update:    true,
		Delete:    isDel,
	})

	if err != nil {
		return err
	}

	return nil
}

// sidecarInjection enables/disables sidecar injection on a namespace
func (mesh *Mesh) sidecarInjection(namespace string, del bool) error {
	kclient := mesh.KubeClient
	if kclient == nil {
		return ErrNilClient
	}

	// updating the label on the namespace
	ns, err := kclient.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		return err
	}

	if ns.ObjectMeta.Labels == nil {
		ns.ObjectMeta.Labels = map[string]string{}
	}
	ns.ObjectMeta.Labels["traefikmesh.io/monitored-by"] = "traefik"

	if del {
		delete(ns.ObjectMeta.Labels, "traefikmesh.io/monitored-by")
	}

	// updating the annotations on the namespace
	if ns.ObjectMeta.Annotations == nil {
		ns.ObjectMeta.Annotations = map[string]string{}
	}
	ns.ObjectMeta.Annotations["traefikmesh.io/sidecar-injection"] = "enabled"

	if del {
		delete(ns.ObjectMeta.Annotations, "traefikmesh.io/sidecar-injection")
	}

	fmt.Println(ns.ObjectMeta)

	_, err = kclient.CoreV1().Namespaces().Update(context.TODO(), ns, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}
