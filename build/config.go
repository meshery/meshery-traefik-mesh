package build

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshkit/utils"
	"github.com/layer5io/meshkit/utils/manifests"
	walker "github.com/layer5io/meshkit/utils/walker"
	smp "github.com/layer5io/service-mesh-performance/spec"
)

var DefaultVersion string
var DefaultURL string
var DefaultGenerationMethod string
var WorkloadPath string
var AllVersions []string
var CRDNames []string

// NewConfig creates the configuration for creating components
func NewConfig(version string) manifests.Config {
	return manifests.Config{
		Name:        smp.ServiceMesh_Type_name[int32(smp.ServiceMesh_TRAEFIK_MESH)],
		MeshVersion: version,
		CrdFilter: manifests.NewCueCrdFilter(manifests.ExtractorPaths{
			NamePath:    "spec.names.kind",
			IdPath:      "spec.names.kind",
			VersionPath: "spec.versions[0].name",
			GroupPath:   "spec.group",
			SpecPath:    "spec.versions[0].schema.openAPIV3Schema.properties.spec"}, false),
		ExtractCrds: func(manifest string) []string {
			crds := strings.Split(manifest, "---")
			// trim the spaces
			for _, crd := range crds {
				crd = strings.TrimSpace(crd)
			}
			return crds
		},
	}
}

func init() {
	wd, _ := os.Getwd()
	WorkloadPath = filepath.Join(wd, "templates", "oam", "workloads")
	AllVersions, _ = utils.GetLatestReleaseTagsSorted("traefik", "mesh")
	if len(AllVersions) == 0 {
		return
	}
	DefaultVersion = AllVersions[len(AllVersions)-1]
	DefaultGenerationMethod = adapter.Manifests

	//Get all the crd names
	w := walker.NewGithub()
	err := w.Owner("traefik").
		Repo("mesh-helm-chart").
		Branch("master").
		Root("mesh/crds/**").
		RegisterFileInterceptor(func(gca walker.GithubContentAPI) error {
			if gca.Content != "" {
				CRDNames = append(CRDNames, gca.Name)
			}
			return nil
		}).Walk()
	if err != nil {
		fmt.Println("Could not find CRD names. Will fail component creation...", err.Error())
	}
	DefaultURL = "https://raw.githubusercontent.com/traefik/mesh-helm-chart/" + "master" + "/mesh/crds/"
}
