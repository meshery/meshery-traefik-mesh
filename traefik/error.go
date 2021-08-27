package traefik

import (
	"fmt"

	"github.com/layer5io/meshkit/errors"
)

var (
	// Error code for failed service mesh installation

	// ErrInstallTraefikCode represents the errors which are generated
	// during traefik mesh install process
	ErrInstallTraefikCode = "replace_traefik_test_code"

	// ErrCreatingHelmIndexCode represents the errors which are generated
	// during creation of helm index
	ErrCreatingHelmIndexCode = "replace_traefik_test_code"

	// ErrEntryWithAppVersionNotExistsCode represents the error which is generated
	// when no entry is found with specified name and app version
	ErrEntryWithAppVersionNotExistsCode = "replace_traefik_test_code"

	// ErrHelmRepositoryNotFoundCode represents the error which is generated when
	// no valid helm repository is found
	ErrHelmRepositoryNotFoundCode = "replace_traefik_test_code"

	// ErrDecodeYamlCode represents the error which is generated when yaml
	// decode process fails
	ErrDecodeYamlCode = "replace_traefik_test_code"

	// ErrMeshConfigCode represents the errors which are generated
	// when an invalid mesh config is found
	ErrMeshConfigCode = "replace_traefik_test_code"

	// ErrApplyHelmChartCode represents the error which are generated
	// during the process of applying helm chart
	ErrApplyHelmChartCode = "1012"

	// ErrConvertingAppVersionToChartVersionCode represents the errors which are generated
	// during the process of converting app version to chart version
	ErrConvertingAppVersionToChartVersionCode = "replace_traefik_test_code"

	// ErrCreatingTraefikMeshClientCode represents the errors which are generated
	// during creating traefik client process
	ErrCreatingTraefikMeshClientCode = "replace_traefik_test_code"

	// ErrSampleAppCode represents the errors which are generated
	// duing sample app installation
	ErrSampleAppCode = "replace_traefik_test_code"

	// ErrCustomOperationCode represents the errors which are generated
	// when an invalid addon operation is requested
	ErrCustomOperationCode = "replace_traefik_test_code"

	// ErrOpInvalidCode represents the error code which are generated
	// when an invalid operation is requested
	ErrOpInvalidCode = "replace_traefik_test_code"

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

	// ErrParseOAMComponent represents the error which is
	// generated during the OAM component parsing
	ErrParseOAMComponent = errors.New(ErrParseOAMComponentCode, errors.Alert, []string{"error parsing the component"}, []string{"Error occured while prasing application component in the OAM request made"}, []string{"Invalid OAM component passed in OAM request"}, []string{"Check if your request has vaild OAM components"})

	// ErrParseOAMConfig represents the error which is
	// generated during the OAM configuration parsing
	ErrParseOAMConfig = errors.New(ErrParseOAMConfigCode, errors.Alert, []string{"error parsing the configuration"}, []string{"Error occured while prasing component config in the OAM request made"}, []string{"Invalid OAM config passed in OAM request"}, []string{"Check if your request has vaild OAM config"})
)

// ErrInstallTraefik is the error for install mesh
func ErrInstallTraefik(err error) error {
	return errors.New(ErrInstallTraefikCode, errors.Alert, []string{"Error with traefik operation"}, []string{err.Error()}, []string{}, []string{})
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
	return errors.New(ErrApplyHelmChartCode, errors.Alert, []string{"Error occured while applying Helm Chart"}, []string{err.Error()}, []string{}, []string{})
}

// ErrConvertingAppVersionToChartVersion is the error for converting app version to chart version
func ErrConvertingAppVersionToChartVersion(err error) error {
	return errors.New(ErrConvertingAppVersionToChartVersionCode, errors.Alert, []string{"Error occured while converting app version to chart version"}, []string{err.Error()}, []string{}, []string{})
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
