package traefik

import (
	"fmt"

	"github.com/layer5io/meshkit/errors"
)

var (
	// Error code for failed service mesh installation

	// ErrInstallTraefikCode represents the errors which are generated
	// during traefik mesh install process
	ErrInstallTraefikCode = "1032"

	// ErrCreatingHelmIndexCode represents the errors which are generated
	// during creation of helm index
	ErrCreatingHelmIndexCode = "1033"

	// ErrEntryWithAppVersionNotExistsCode represents the error which is generated
	// when no entry is found with specified name and app version
	ErrEntryWithAppVersionNotExistsCode = "1034"

	// ErrHelmRepositoryNotFoundCode represents the error which is generated when
	// no valid helm repository is found
	ErrHelmRepositoryNotFoundCode = "1035"

	// ErrDecodeYamlCode represents the error which is generated when yaml
	// decode process fails
	ErrDecodeYamlCode = "1036"

	// ErrMeshConfigCode represents the errors which are generated
	// when an invalid mesh config is found
	ErrMeshConfigCode = "1037"

	// ErrApplyHelmChartCode represents the error which are generated
	// during the process of applying helm chart
	ErrApplyHelmChartCode = "1012"

	// ErrConvertingAppVersionToChartVersionCode represents the errors which are generated
	// during the process of converting app version to chart version
	ErrConvertingAppVersionToChartVersionCode = "1038"

	// ErrCreatingTraefikMeshClientCode represents the errors which are generated
	// during creating traefik client process
	ErrCreatingTraefikMeshClientCode = "1039"

	// ErrSampleAppCode represents the errors which are generated
	// duing sample app installation
	ErrSampleAppCode = "1040"

	// ErrCustomOperationCode represents the errors which are generated
	// when an invalid addon operation is requested
	ErrCustomOperationCode = "1041"

	// ErrOpInvalidCode represents the error code which are generated
	// when an invalid operation is requested
	ErrOpInvalidCode = "1042"

	// ErrOpInvalid is an error when an invalid operation is requested
	ErrOpInvalid = errors.New(ErrOpInvalidCode, errors.Alert, []string{"Invalid operation"}, []string{}, []string{}, []string{})

	// ErrNilClientCode represents the error code which is
	// generated when kubernetes client is nil
	ErrNilClientCode = "1013"

	// ErrInvalidOAMComponentTypeCode represents the error code which is
	// generated when an invalid oam component is requested
	ErrInvalidOAMComponentTypeCode = "1014"

	// ErrTraefikCoreComponentFailCode represents the error code which is
	// generated when an Traefik core operations fails
	ErrTraefikCoreComponentFailCode = "1015"
	// ErrProcessOAMCode represents the error code which is
	// generated when an OAM operations fails
	ErrProcessOAMCode = "1016"
	// ErrParseTraefikCoreComponentCode represents the error code which is
	// generated when Traefik core component manifest parsing fails
	ErrParseTraefikCoreComponentCode = "1017"
	// ErrParseOAMComponentCode represents the error code which is
	// generated during the OAM component parsing
	ErrParseOAMComponentCode = "1018"
	// ErrParseOAMConfigCode represents the error code which is
	// generated during the OAM configuration parsing
	ErrParseOAMConfigCode = "1019"

	// ErrNilClient represents the error which is
	// generated when kubernetes client is nil
	ErrNilClient = errors.New(ErrNilClientCode, errors.Alert, []string{"kubernetes client not initialized"}, []string{"Kubernetes client is nil"}, []string{"kubernetes client not initialized"}, []string{"Reconnect the adapter to Meshery server"})

	// ErrParseOAMComponent represents the error which is
	// generated during the OAM component parsing
	ErrParseOAMComponent = errors.New(ErrParseOAMComponentCode, errors.Alert, []string{"error parsing the component"}, []string{"Error occurred while parsing application component in the OAM request made by Meshery Server"}, []string{"Could not unmarshall configuration component received via ProcessOAM gRPC call into a valid Component struct"}, []string{"Check if Meshery Server is creating valid component for ProcessOAM gRPC call. This error should never happen and can be reported as a bug in Meshery Server. Also, check if Meshery Server and adapters are referring to same component struct provided in MeshKit."})

	// ErrParseOAMConfig represents the error which is
	// generated during the OAM configuration parsing
	ErrParseOAMConfig = errors.New(ErrParseOAMConfigCode, errors.Alert, []string{"error parsing the configuration"}, []string{"Error occurred while parsing configuration in the request made by Meshery Server"}, []string{"Could not unmarshall OAM config received via ProcessOAM gRPC call into a valid Config struct"}, []string{"Check if Meshery Server is creating valid config for ProcessOAM gRPC call. This error should never happen and can be reported as a bug in Meshery Server. Also, confirm that Meshery Server and Adapters are referring to same config struct provided in MeshKit"})

	// ErrGetLatestReleaseCode represents the error which is
	// generated when the latest stable version could not
	// be fetched during runtime component registeration
	ErrGetLatestReleaseCode = "1020"
)

// ErrInstallTraefik is the error for install mesh
func ErrInstallTraefik(err error) error {
	return errors.New(ErrInstallTraefikCode, errors.Alert, []string{"Error with traefik operation"}, []string{err.Error()}, []string{}, []string{})
}

// ErrDecodeYaml is the error when the yaml unmarshal fails
func ErrDecodeYaml(err error) error {
	return errors.New(ErrDecodeYamlCode, errors.Alert, []string{"Error occurred while decoding YAML"}, []string{err.Error()}, []string{}, []string{})
}

// ErrCreatingHelmIndex is the error for creating helm index
func ErrCreatingHelmIndex(err error) error {
	return errors.New(ErrCreatingHelmIndexCode, errors.Alert, []string{"Error while creating Helm Index"}, []string{err.Error()}, []string{}, []string{})
}

// ErrEntryWithAppVersionNotExists is the error when an entry with the given app version is not found
func ErrEntryWithAppVersionNotExists(entry, appVersion string) error {
	return errors.New(ErrEntryWithAppVersionNotExistsCode, errors.Alert, []string{"Entry for the app version does not exist"}, []string{fmt.Sprintf("entry %s with app version %s does not exists", entry, appVersion)}, []string{}, []string{})
}

// ErrHelmRepositoryNotFound is the error when no valid remote helm repository is found
func ErrHelmRepositoryNotFound(repo string, err error) error {
	return errors.New(ErrHelmRepositoryNotFoundCode, errors.Alert, []string{"Helm repo not found"}, []string{fmt.Sprintf("either the repo %s does not exists or is corrupt: %v", repo, err)}, []string{}, []string{})
}

// ErrMeshConfig is the error for mesh config
func ErrMeshConfig(err error) error {
	return errors.New(ErrMeshConfigCode, errors.Alert, []string{"Error configuration mesh"}, []string{err.Error()}, []string{}, []string{})
}

// ErrApplyHelmChart is the error for applying helm chart
func ErrApplyHelmChart(err error) error {
	return errors.New(ErrApplyHelmChartCode, errors.Alert, []string{"Error occurred while applying Helm Chart"}, []string{err.Error()}, []string{}, []string{})
}

// ErrConvertingAppVersionToChartVersion is the error for converting app version to chart version
func ErrConvertingAppVersionToChartVersion(err error) error {
	return errors.New(ErrConvertingAppVersionToChartVersionCode, errors.Alert, []string{"Error occurred while converting app version to chart version"}, []string{err.Error()}, []string{}, []string{})
}

// ErrCreatingTraefikMeshClient is the error for streaming event
func ErrCreatingTraefikMeshClient(err error) error {
	return errors.New(ErrCreatingTraefikMeshClientCode, errors.Alert, []string{"Unable to create a new NSM client"}, []string{err.Error()}, []string{}, []string{})
}

// ErrSampleApp is the error for streaming event
func ErrSampleApp(err error) error {
	return errors.New(ErrSampleAppCode, errors.Alert, []string{"Error with sample app operation"}, []string{err.Error()}, []string{}, []string{})
}

// ErrCustomOperation is the error for streaming event
func ErrCustomOperation(err error) error {
	return errors.New(ErrCustomOperationCode, errors.Alert, []string{"Error with custom operation"}, []string{err.Error()}, []string{}, []string{})
}

// ErrParseTraefikCoreComponent is the error when traefik core component manifest parsing fails
func ErrParseTraefikCoreComponent(err error) error {
	return errors.New(ErrParseTraefikCoreComponentCode, errors.Alert, []string{"Traefik core component manifest parsing failing"}, []string{err.Error()}, []string{}, []string{})
}

// ErrInvalidOAMComponentType is the error when the OAM component name is not valid
func ErrInvalidOAMComponentType(compName string) error {
	return errors.New(ErrInvalidOAMComponentTypeCode, errors.Alert, []string{"invalid OAM component name: ", compName}, []string{}, []string{}, []string{})
}

// ErrTraefikCoreComponentFail is the error when core Traefik component processing fails
func ErrTraefikCoreComponentFail(err error) error {
	return errors.New(ErrTraefikCoreComponentFailCode, errors.Alert, []string{"error in Traefik core component"}, []string{err.Error()}, []string{}, []string{})
}

// ErrProcessOAM is a generic error which is thrown when an OAM operations fails
func ErrProcessOAM(err error) error {
	return errors.New(ErrProcessOAMCode, errors.Alert, []string{"error performing OAM operations"}, []string{err.Error()}, []string{}, []string{})
}

// ErrGetLatestRelease is the error for get latest versions
func ErrGetLatestRelease(err error) error {
	return errors.New(ErrGetLatestReleaseCode, errors.Alert, []string{"Could not get latest version"}, []string{err.Error()}, []string{"Latest version could not be found at the specified url"}, []string{})
}
