// Package driver provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
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

// Defines values for AttributeDefinitionSchemaType.
const (
	BINARY  AttributeDefinitionSchemaType = "BINARY"
	BOOLEAN AttributeDefinitionSchemaType = "BOOLEAN"
	INTEGER AttributeDefinitionSchemaType = "INTEGER"
	NUMBER  AttributeDefinitionSchemaType = "NUMBER"
	STRING  AttributeDefinitionSchemaType = "STRING"
)

// AttributeDefinitionSchema Schema that describes a driver attribute.
type AttributeDefinitionSchema struct {
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
	Type *AttributeDefinitionSchemaType `json:"type,omitempty"`
}

// AttributeDefinitionSchemaType The type of the attribute.
//   - `STRING` - The attribute is a string.
//   - `INTEGER` - The attribute is an integer.
//   - `NUMBER` - The attribute is a decimal number.
//   - `BOOLEAN` - The attribute is a boolean.
//   - `BINARY` - The attribute is a binary.
type AttributeDefinitionSchemaType string

// DriverActionAttributeSchema Schema that describes driver action attributes.
type DriverActionAttributeSchema struct {
	// Attributes Schema that describes a list of attributes supported by a driver.
	Attributes *DriverAttributesSchema `json:"attributes,omitempty"`

	// Type The type of action.
	//   * `GET_REGISTER` - The action is to get a billing value, for example, instantaneous values.
	//   * `GET_PERIODICAL_PROFILE` - The action is to get a periodical profile, for example, load-profile.
	//   * `GET_IRREGULAR_PROFILE` - The action is to get a non-periodical profile, for, daily profile or monthly billing registers.
	//   * `GET_EVENTS` - The action is to get an event log.
	//   * `GET_CLOCK` - The action is to get the clock.
	//   * `SYNC_CLOCK` - The action is to synchronize the clock. The action synchronizes the clock in the device. If the force attribute is set, it forcefully sets the clock.
	//   * `GET_RELAY_STATE` - The action is to get the relay state.
	//   * `SET_RELAY_STATE` - The action is to set the relay state.
	//   * `GET_DISCONNECTOR_STATE` - The action is to get the disconnector state.
	//   * `SET_DISCONNECTOR_STATE` - The action is to set the disconnector state.
	//   * `GET_TOU` - The action is to get the time-of-use table.
	//   * `SET_TOU` - The action is to set the time-of-use table.
	//   * `GET_LIMITER` - The action is to get the limiter settings.
	//   * `SET_LIMITER` - The action is to set the limiter settings.
	//   * `RESET_BILLING_PERIOD` - The action is to reset the billing period.
	//   * `FW_UPDATE` - The action is to start a firmware update.
	Type *externalRef2.ActionTypeSchema `json:"type,omitempty"`
}

// DriverAttributesSchema Schema that describes a list of attributes supported by a driver.
type DriverAttributesSchema = []AttributeDefinitionSchema

// DriverDetailsSchema Schema that describes communication driver details.
type DriverDetailsSchema = DriverSchema

// DriverSchema Schema that describes communication driver details.
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
	"CDYat0XM2OP/AgAA///8zLsR6RsAAA==",
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

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(path.Dir(pathToFile), "../../k8s/containers/containers.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef1.PathToRawSpec(path.Join(path.Dir(pathToFile), "../../openhes/driver/driverdata/driverdata.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef2.PathToRawSpec(path.Join(path.Dir(pathToFile), "../job/job.yaml")) {
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
