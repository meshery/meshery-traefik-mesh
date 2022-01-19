package build

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/layer5io/meshery-adapter-library/adapter"
	"github.com/layer5io/meshery-traefik-mesh/internal/config"
	"github.com/layer5io/meshkit/utils/manifests"
	smp "github.com/layer5io/service-mesh-performance/spec"
	"gopkg.in/yaml.v2"
)

var DefaultVersion string
var DefaultURL string
var DefaultGenerationMethod = adapter.HelmCHARTS
var WorkloadPath string

//NewConfig creates the configuration for creating components
func NewConfig(version string) manifests.Config {
	return manifests.Config{
		Name:        smp.ServiceMesh_Type_name[int32(smp.ServiceMesh_TRAEFIK_MESH)],
		MeshVersion: version,
		Filter: manifests.CrdFilter{
			RootFilter:    []string{"$[?(@.kind==\"CustomResourceDefinition\")]"},
			NameFilter:    []string{"$..[\"spec\"][\"names\"][\"kind\"]"},
			VersionFilter: []string{"$[0]..spec.versions[0]"},
			GroupFilter:   []string{"$[0]..spec"},
			SpecFilter:    []string{"$[0]..openAPIV3Schema.properties.spec"},
			ItrFilter:     []string{"$[?(@.spec.names.kind"},
			ItrSpecFilter: []string{"$[?(@.spec.names.kind"},
			VField:        "name",
			GField:        "group",
		},
	}
}

// returns latest valid appversion and chartversion
func getLatestValidAppVersionAndChartVersion() (string, string, error) {
	res, err := http.Get("https://helm.traefik.io/traefik/index.yaml")
	if err != nil {
		return "", "", config.ErrGetLatestReleases(err)
	}
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", "", config.ErrGetLatestReleases(err)
	}
	var h helmIndex
	err = yaml.Unmarshal(content, &h)
	if err != nil {
		return "", "", config.ErrGetLatestReleases(err)
	}

	return h.Entries["traefik"][0].AppVersion, h.Entries["traefik"][0].Version, nil
}

// Below structs are helper structs to unmarshall and extract certain fields from helm chart's index.yaml
type helmIndex struct {
	Entries map[string][]data `yaml:"entries"`
}

type data struct {
	AppVersion string `yaml:"appVersion"`
	Version    string `yaml:"version"`
}

func init() {
	wd, _ := os.Getwd()
	WorkloadPath = filepath.Join(wd, "templates", "oam", "workloads")
	var chartv string
	DefaultVersion, chartv, _ = getLatestValidAppVersionAndChartVersion()
	DefaultURL = "https://helm.traefik.io/traefik/traefik-" + chartv + ".tgz"

}
