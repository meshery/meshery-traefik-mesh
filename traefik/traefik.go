// Package traefik - Common operations for the adapter
package traefik

import (
	"context"
	"fmt"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-adapter-library/common"
	"github.com/layer5io/meshery-adapter-library/meshes"
	"github.com/layer5io/meshery-adapter-library/status"
	internalconfig "github.com/layer5io/meshery-traefik-mesh/internal/config"
	"github.com/layer5io/meshery-traefik-mesh/traefik/oam"
	meshkitCfg "github.com/layer5io/meshkit/config"
	"github.com/layer5io/meshkit/errors"
	"github.com/layer5io/meshkit/logger"
	"github.com/layer5io/meshkit/models"
	"github.com/layer5io/meshkit/models/oam/core/v1alpha1"
	"github.com/layer5io/meshkit/utils/events"
	"gopkg.in/yaml.v2"
)

const (
	// SMIManifest is the manifest.yaml file for smi conformance tool
	SMIManifest = "https://raw.githubusercontent.com/layer5io/learn-layer5/master/smi-conformance/manifest.yml"
)

// Mesh represents the traefik-mesh adapter and embeds adapter.Adapter
type Mesh struct {
	adapter.Adapter // Type Embedded
}

// New initializes treafik-mesh handler.
func New(c meshkitCfg.Handler, l logger.Handler, kc meshkitCfg.Handler, e *events.EventStreamer) adapter.Handler {
	return &Mesh{
		Adapter: adapter.Adapter{
			Config:            c,
			Log:               l,
			KubeconfigHandler: kc,
			EventStreamer:     e,
		},
	}
}

//CreateKubeconfigs creates and writes passed kubeconfig onto the filesystem
func (mesh *Mesh) CreateKubeconfigs(kubeconfigs []string) error {
	var errs = make([]error, 0)
	for _, kubeconfig := range kubeconfigs {
		kconfig := models.Kubeconfig{}
		err := yaml.Unmarshal([]byte(kubeconfig), &kconfig)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		// To have control over what exactly to take in on kubeconfig
		mesh.KubeconfigHandler.SetKey("kind", kconfig.Kind)
		mesh.KubeconfigHandler.SetKey("apiVersion", kconfig.APIVersion)
		mesh.KubeconfigHandler.SetKey("current-context", kconfig.CurrentContext)
		err = mesh.KubeconfigHandler.SetObject("preferences", kconfig.Preferences)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = mesh.KubeconfigHandler.SetObject("clusters", kconfig.Clusters)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = mesh.KubeconfigHandler.SetObject("users", kconfig.Users)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		err = mesh.KubeconfigHandler.SetObject("contexts", kconfig.Contexts)
		if err != nil {
			errs = append(errs, err)
			continue
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return mergeErrors(errs)
}

// ApplyOperation applies the operation on traefik mesh
func (mesh *Mesh) ApplyOperation(ctx context.Context, opReq adapter.OperationRequest) error {
	err := mesh.CreateKubeconfigs(opReq.K8sConfigs)
	if err != nil {
		return err
	}
	kubeconfigs := opReq.K8sConfigs
	operations := make(adapter.Operations)
	err = mesh.Config.GetObject(adapter.OperationsKey, &operations)
	if err != nil {
		return err
	}

	e := &meshes.EventsResponse{
		OperationId:   opReq.OperationID,
		Summary:       status.Deploying,
		Details:       "Operation is not supported",
		Component:     internalconfig.ServerConfig["type"],
		ComponentName: internalconfig.ServerConfig["name"],
	}

	switch opReq.OperationName {
	case internalconfig.TraefikMeshOperation:
		go func(hh *Mesh, ee *meshes.EventsResponse) {
			version := string(operations[opReq.OperationName].Versions[0])
			stat, err := hh.installTraefikMesh(opReq.IsDeleteOperation, version, opReq.Namespace, kubeconfigs)
			if err != nil {
				summary := fmt.Sprintf("Error while %s Traefik service mesh", stat)
				hh.streamErr(summary, ee, err)
				return
			}
			ee.Summary = fmt.Sprintf("Traefik service mesh %s successfully", stat)
			ee.Details = fmt.Sprintf("The Traefik service mesh is now %s.", stat)
			hh.StreamInfo(ee)
		}(mesh, e)
	case common.BookInfoOperation, common.HTTPBinOperation, common.ImageHubOperation, common.EmojiVotoOperation:
		go func(hh *Mesh, ee *meshes.EventsResponse) {
			appName := operations[opReq.OperationName].AdditionalProperties[common.ServiceName]
			stat, err := hh.installSampleApp(opReq.Namespace, opReq.IsDeleteOperation, operations[opReq.OperationName].Templates, kubeconfigs)
			if err != nil {
				summary := fmt.Sprintf("Error while %s %s application", stat, appName)
				hh.streamErr(summary, ee, err)
				return
			}
			ee.Summary = fmt.Sprintf("%s application %s successfully", appName, stat)
			ee.Details = fmt.Sprintf("The %s application is now %s.", appName, stat)
			hh.StreamInfo(ee)
		}(mesh, e)
	case common.CustomOperation:
		go func(hh *Mesh, ee *meshes.EventsResponse) {
			stat, err := hh.applyCustomOperation(opReq.Namespace, opReq.CustomBody, opReq.IsDeleteOperation, kubeconfigs)
			if err != nil {
				summary := fmt.Sprintf("Error while %s custom operation", stat)
				hh.streamErr(summary, ee, err)
				return
			}
			ee.Summary = fmt.Sprintf("Manifest %s successfully", status.Deployed)
			ee.Details = ""
			hh.StreamInfo(ee)
		}(mesh, e)
	case common.SmiConformanceOperation:
		go func(hh *Mesh, ee *meshes.EventsResponse) {
			name := operations[opReq.OperationName].Description
			_, err := hh.RunSMITest(adapter.SMITestOptions{
				Ctx:         context.TODO(),
				OperationID: ee.OperationId,
				Manifest:    SMIManifest,
				Namespace:   "meshery",
				Labels:      make(map[string]string),
				Annotations: make(map[string]string),
				Kubeconfigs: kubeconfigs,
			})
			if err != nil {
				summary := fmt.Sprintf("Error while %s %s test", status.Running, name)
				hh.streamErr(summary, ee, err)
				return
			}
			ee.Summary = fmt.Sprintf("%s test %s successfully", name, status.Completed)
			ee.Details = ""
			hh.StreamInfo(ee)
		}(mesh, e)
	default:
		mesh.streamErr("Invalid operation", e, ErrOpInvalid)
	}

	return nil
}

// ProcessOAM will handles the grpc invocation for handling OAM objects
func (mesh *Mesh) ProcessOAM(ctx context.Context, oamReq adapter.OAMRequest) (string, error) {
	err := mesh.CreateKubeconfigs(oamReq.K8sConfigs)
	if err != nil {
		return "", err
	}
	kubeconfigs := oamReq.K8sConfigs
	var comps []v1alpha1.Component
	for _, acomp := range oamReq.OamComps {
		comp, err := oam.ParseApplicationComponent(acomp)
		if err != nil {
			mesh.Log.Error(ErrParseOAMComponent)
			continue
		}

		comps = append(comps, comp)
	}

	config, err := oam.ParseApplicationConfiguration(oamReq.OamConfig)
	if err != nil {
		mesh.Log.Error(ErrParseOAMConfig)
	}

	// If operation is delete then first HandleConfiguration and then handle the deployment
	if oamReq.DeleteOp {
		// Process configuration
		msg2, err := mesh.HandleApplicationConfiguration(config, oamReq.DeleteOp, kubeconfigs)
		if err != nil {
			return msg2, ErrProcessOAM(err)
		}

		// Process components
		msg1, err := mesh.HandleComponents(comps, oamReq.DeleteOp, kubeconfigs)
		if err != nil {
			return msg1 + "\n" + msg2, ErrProcessOAM(err)
		}

		return msg1 + "\n" + msg2, nil
	}

	// Process components
	msg1, err := mesh.HandleComponents(comps, oamReq.DeleteOp, kubeconfigs)
	if err != nil {
		return msg1, ErrProcessOAM(err)
	}

	// Process configuration
	msg2, err := mesh.HandleApplicationConfiguration(config, oamReq.DeleteOp, kubeconfigs)
	if err != nil {
		return msg1 + "\n" + msg2, ErrProcessOAM(err)
	}

	return msg1 + "\n" + msg2, nil
}

func (mesh *Mesh) streamErr(summary string, e *meshes.EventsResponse, err error) {
	e.Summary = summary
	e.Details = err.Error()
	e.ErrorCode = errors.GetCode(err)
	e.ProbableCause = errors.GetCause(err)
	e.SuggestedRemediation = errors.GetRemedy(err)
	mesh.StreamErr(e, err)
}
