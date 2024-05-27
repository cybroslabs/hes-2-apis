// Package driver provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package driver

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	externalRef0 "github.com/cybroslabs/hes-2-apis/openapi/k8s/containers"
	externalRef1 "github.com/cybroslabs/hes-2-apis/openapi/openhes/driver/driverdata"
	externalRef2 "github.com/cybroslabs/hes-2-apis/openapi/openhes/job"
	"github.com/getkin/kin-openapi/openapi3"
)

// Defines values for DriverAttributeSchemaType.
const (
	BINARY  DriverAttributeSchemaType = "BINARY"
	BOOLEAN DriverAttributeSchemaType = "BOOLEAN"
	INTEGER DriverAttributeSchemaType = "INTEGER"
	NUMBER  DriverAttributeSchemaType = "NUMBER"
	STRING  DriverAttributeSchemaType = "STRING"
)

// DriverActionAttributeSchema Schema that holds the specification of the driver action attribute.
type DriverActionAttributeSchema struct {
	// Attributes The list of attributes supported by the driver.
	Attributes *DriverAttributesSchema `json:"attributes,omitempty"`

	// Type The type of action.
	//   * `GET_REGISTER` - The action is to get billing value, for example instantaneous values.
	//   * `GET_PERIODICAL_PROFILE` - The action is to get periodical profile, for example load-profile.
	//   * `GET_IRREGULAR_PROFILE` - The action is to get non-periodical profile, for example daily profile or monthly billing registers.
	//   * `GET_EVENTS` - The action is to get event log.
	//   * `GET_CLOCK` - The action is to get the clock.
	//   * `SYNC_CLOCK` - The action is to synchronize the clock. The action synchronizes the clock in the device, it can forcefully set it if force attribute is set.
	//   * `GET_RELAY_STATE` - The action is to get the relay state.
	//   * `SET_RELAY_STATE` - The action is to set the relay state.
	//   * `GET_DISCONNECTOR_STATE` - The action is to get the disconnector state.
	//   * `SET_DISCONNECTOR_STATE` - The action is to set the disconnector state.
	//   * `GET_TOU` - The action is to get the time-of-use table.
	//   * `SET_TOU` - The action is to set the time-of-use table.
	//   * `GET_LIMITER` - The action is to get the limiter settings.
	//   * `SET_LIMITER` - The action is to set the limiter settings.
	//   * `RESET_BILLING_PERIOD` - The action is to reset billing period.
	//   * `FW_UPDATE` - The action is to start firmware update.
	Type *externalRef2.ActionTypeSchema `json:"type,omitempty"`
}

// DriverAttributeSchema Schema that holds the specification of the driver attribute.
type DriverAttributeSchema struct {
	DefaultValue *GenericObject `json:"defaultValue,omitempty"`

	// Description The description of the attribute.
	Description *string `json:"description,omitempty"`

	// Mandatory Indicates whether the attribute is mandatory or not.
	Mandatory *bool `json:"mandatory,omitempty"`

	// Name The name of the attribute.
	Name *string `json:"name,omitempty"`

	// Type The type of the attribute.
	//   * `STRING` - The attribute is a string.
	//   * `INTEGER` - The attribute is an integer.
	//   * `NUMBER` - The attribute is a decimal number.
	//   * `BOOLEAN` - The attribute is a boolean.
	//   * `BINARY` - The attribute is a binary.
	Type *DriverAttributeSchemaType `json:"type,omitempty"`
}

// DriverAttributeSchemaType The type of the attribute.
//   - `STRING` - The attribute is a string.
//   - `INTEGER` - The attribute is an integer.
//   - `NUMBER` - The attribute is a decimal number.
//   - `BOOLEAN` - The attribute is a boolean.
//   - `BINARY` - The attribute is a binary.
type DriverAttributeSchemaType string

// DriverAttributesSchema The list of attributes supported by the driver.
type DriverAttributesSchema = []DriverAttributeSchema

// DriverDetailsSchema Schema that holds the communication driver details.
type DriverDetailsSchema = DriverSchema

// DriverSchema Schema that holds the communication driver details.
type DriverSchema struct {
	// Spec Schema that registers new or updates already registered communication drivers.
	Spec *DriverSpecSchema `json:"spec,omitempty"`
}

// DriverSpecSchema Schema that registers new or updates already registered communication drivers.
type DriverSpecSchema struct {
	DriverType string `json:"driverType"`
	Image      string `json:"image"`

	// Resources Resource values for a container.
	Resources externalRef0.ContainerResourcesSetSchema `json:"resources"`
}

// GenericObject defines model for GenericObject.
type GenericObject = interface{}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	pathPrefix := path.Dir(pathToFile)

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "../../k8s/containers/containers.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef1.PathToRawSpec(path.Join(pathPrefix, "../../openhes/driver/driverdata/driverdata.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef2.PathToRawSpec(path.Join(pathPrefix, "../job/job.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
