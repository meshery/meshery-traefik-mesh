package traefik

import (
	"context"
	"fmt"

	"github.com/layer5io/meshery-adapter-library/adapter"
	adapterconfig "github.com/layer5io/meshery-adapter-library/config"
	"github.com/layer5io/meshery-adapter-library/status"
	internalConfig "github.com/layer5io/meshery-maesh/internal/config"
	"github.com/layer5io/meshkit/logger"
)

// Mesh represents the traefik-mesh adapter and embeds adapter.Adapter
type Mesh struct {
	adapter.Adapter // Type Embedded
}

// New initializes treafik-mesh handler.
func New(c adapterconfig.Handler, l logger.Handler, kc adapterconfig.Handler) adapter.Handler {
	return &Mesh{
		Adapter: adapter.Adapter{
			Config:            c,
			Log:               l,
			KubeconfigHandler: kc,
		},
	}
}

// ApplyOperation applies the operation on traefik mesh
func (mesh *Mesh) ApplyOperation(ctx context.Context, opReq adapter.OperationRequest) error {
	operations := make(adapter.Operations)
	err := mesh.Config.GetObject(adapter.OperationsKey, &operations)
	if err != nil {
		return err
	}

	e := &adapter.Event{
		Operationid: opReq.OperationID,
		Summary:     status.Deploying,
		Details:     "Operation is not supported",
	}

	switch opReq.OperationName {
	case internalConfig.TraefikMeshOperation:
		go func(hh *Mesh, ee *adapter.Event) {
			version := string(operations[opReq.OperationName].Versions[0])
			stat, err := hh.installTraefikMesh(opReq.IsDeleteOperation, version, opReq.Namespace)
			if err != nil {
				e.Summary = fmt.Sprintf("Error while %s Traefik service mesh", stat)
				e.Details = err.Error()
				hh.StreamErr(e, err)
				return
			}
			ee.Summary = fmt.Sprintf("Traefik service mesh %s successfully", stat)
			ee.Details = fmt.Sprintf("The Traefik service mesh is now %s.", stat)
			hh.StreamInfo(e)
		}(mesh, e)
	default:
		mesh.StreamErr(e, ErrOpInvalid)
	}

	return nil
}
