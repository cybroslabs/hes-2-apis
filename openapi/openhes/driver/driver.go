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

// DriverDetailsSchema defines model for DriverDetailsSchema.
type DriverDetailsSchema struct {
	// Spec Schema that registers new or updates already registered communication drivers.
	Spec *DriverSpecSchema `json:"spec,omitempty"`

	// Status Schema that describes the status of a driver.
	Status *DriverStatusSchema `json:"status,omitempty"`
}

// DriverSchema Schema that describes communication driver details.
type DriverSchema struct {
	// Spec Schema that registers new or updates already registered communication drivers.
	Spec *DriverSpecSchema `json:"spec,omitempty"`
}

// DriverSpecSchema Schema that registers new or updates already registered communication drivers.
type DriverSpecSchema struct {
	Image string `json:"image"`

	// Resources Resource values for a container.
	Resources externalRef0.ContainerResourcesSetSchema `json:"resources"`
}

// DriverStatusSchema Schema that describes the status of a driver.
type DriverStatusSchema struct {
	DriverType string `json:"driverType"`
}

// GenericObject defines model for GenericObject.
type GenericObject = interface{}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"GD7+NwAA//8+ksxcVh0AAA==",
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
