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

	externalRef0 "github.com/cybroslabs/hes-2-apis/openapi/openhes/driver/driverdata"
	externalRef1 "github.com/cybroslabs/hes-2-apis/openapi/openhes/job"
	"github.com/getkin/kin-openapi/openapi3"
)

// Defines values for AttributeDefinitionSchemaType.
const (
	BINARY      AttributeDefinitionSchemaType = "BINARY"
	BOOLEAN     AttributeDefinitionSchemaType = "BOOLEAN"
	INTEGER     AttributeDefinitionSchemaType = "INTEGER"
	NUMBER      AttributeDefinitionSchemaType = "NUMBER"
	STRING      AttributeDefinitionSchemaType = "STRING"
	TIMESTAMP   AttributeDefinitionSchemaType = "TIMESTAMP"
	TIMESTAMPTZ AttributeDefinitionSchemaType = "TIMESTAMPTZ"
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
	//   * `TIMESTAMP` - The attribute is a timestamp; it shal not timezone-aware and it's always encoded as a UNIX timestamp number (~UTC).
	//   * `TIMESTAMPTZ` - The attribute is a timestamp; it is timezone-aware and it's always encoded as an ISO 8601 string.
	Type *AttributeDefinitionSchemaType `json:"type,omitempty"`
}

// AttributeDefinitionSchemaType The type of the attribute.
//   - `STRING` - The attribute is a string.
//   - `INTEGER` - The attribute is an integer.
//   - `NUMBER` - The attribute is a decimal number.
//   - `BOOLEAN` - The attribute is a boolean.
//   - `BINARY` - The attribute is a binary.
//   - `TIMESTAMP` - The attribute is a timestamp; it shal not timezone-aware and it's always encoded as a UNIX timestamp number (~UTC).
//   - `TIMESTAMPTZ` - The attribute is a timestamp; it is timezone-aware and it's always encoded as an ISO 8601 string.
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
	Type *externalRef1.ActionTypeSchema `json:"type,omitempty"`
}

// DriverAppProtocolRefsSchema The list of application protocol references valid for the datalink template.
type DriverAppProtocolRefsSchema = []string

// DriverAppProtocolSchema Schema that describes the application protocol.
type DriverAppProtocolSchema struct {
	// Attributes Schema that describes a list of attributes supported by a driver.
	Attributes *DriverAttributesSchema `json:"attributes,omitempty"`

	// Protocol The type of the attribute.
	//   * `APPPROTO_IEC_62056_21` - The IEC 62056-21 (IEC-61107, VDEW) protocol.
	//   * `APPPROTO_DLMS_SN` - The DLMS short-name protocol.
	//   * `APPPROTO_DLMS_LN` - The DLMS logical-name protocol.
	//   * `APPPROTO_SCTM` - The SCTM protocol.
	//   * `APPPROTO_LIS200` - The LIS200 protocol.
	//   * `APPPROTO_ANSI_C12` - The ANSI C12 protocol.
	//   * `APPPROTO_MQTT` - The MQTT protocol.
	Protocol *externalRef1.ApplicationProtocolSchema `json:"protocol,omitempty"`
}

// DriverAppProtocolsSchema The list of application protocols valid for the datalink template.
type DriverAppProtocolsSchema = []DriverAppProtocolSchema

// DriverAttributesSchema Schema that describes a list of attributes supported by a driver.
type DriverAttributesSchema = []AttributeDefinitionSchema

// DriverCommunicationTemplateSchema Schema that describes the communication template.
type DriverCommunicationTemplateSchema struct {
	// AppProtocolTemplates The list of application protocols valid for the datalink template.
	AppProtocolTemplates *DriverAppProtocolsSchema `json:"appProtocolTemplates,omitempty"`

	// DatalinkTemplates The list of datalink templates valid for the communication template.
	DatalinkTemplates *DriverDatalinkTemplateSchemas `json:"datalinkTemplates,omitempty"`

	// Type The type of the communication template.
	Type *string `json:"type,omitempty"`
}

// DriverDatalinkTemplateSchema Schema that describes the datalink template.
type DriverDatalinkTemplateSchema struct {
	// AppProtocolRefs The list of application protocol references valid for the datalink template.
	AppProtocolRefs *DriverAppProtocolRefsSchema `json:"appProtocolRefs,omitempty"`

	// Attributes Schema that describes a list of attributes supported by a driver.
	Attributes *DriverAttributesSchema `json:"attributes,omitempty"`

	// LinkProtocol The link protocol.
	LinkProtocol *string `json:"linkProtocol,omitempty"`
}

// DriverDatalinkTemplateSchemas The list of datalink templates valid for the communication template.
type DriverDatalinkTemplateSchemas = []DriverDatalinkTemplateSchema

// DriverSchema Schema that describes communication driver details.
type DriverSchema struct {
	// Spec Schema containing communication driver spec.
	Spec *DriverSpecSchema `json:"spec,omitempty"`
}

// DriverSpecSchema Schema containing communication driver spec.
type DriverSpecSchema struct {
	// DriverType The driver type identifier.
	DriverType string `json:"driverType"`

	// Version The version of the driver.
	Version string `json:"version"`
}

// GenericObject defines model for GenericObject.
type GenericObject = interface{}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"/D8AAP//0Z7k61AbAAA=",
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

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(path.Dir(pathToFile), "../../openhes/driver/driverdata/driverdata.yaml")) {
		if _, ok := res[rawPath]; ok {
			// it is not possible to compare functions in golang, so always overwrite the old value
		}
		res[rawPath] = rawFunc
	}
	for rawPath, rawFunc := range externalRef1.PathToRawSpec(path.Join(path.Dir(pathToFile), "../job/job.yaml")) {
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
