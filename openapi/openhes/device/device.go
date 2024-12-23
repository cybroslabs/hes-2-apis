// Package device provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package device

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	externalRef0 "github.com/cybroslabs/hes-2-apis/openapi/openhes/job"
	"github.com/getkin/kin-openapi/openapi3"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

// CommunicationUnitID The ID of the communication unit. The ID must be unique across all OpenHES components.
type CommunicationUnitID = openapi_types.UUID

// CommunicationUnitListID The list of communication unit IDs.
type CommunicationUnitListID = []CommunicationUnitID

// CommunicationUnitSchema Schema that describes communication unit details.
type CommunicationUnitSchema struct {
	ConnectionInfo externalRef0.ConnectionTypeSchema `json:"connectionInfo"`

	// ExternalID The public ID of the communication unit.
	ExternalID *string `json:"externalID"`

	// Id The ID of the communication unit. The ID must be unique across all OpenHES components.
	Id CommunicationUnitID `json:"id"`

	// Name The name of the communication unit.
	Name string `json:"name"`
}

// DeviceGroupID The ID of the device group. The ID must be unique across all OpenHES components.
type DeviceGroupID = openapi_types.UUID

// DeviceGroupSchema Schema that describes device group details.
type DeviceGroupSchema struct {
	// ExternalID The public ID of the device group.
	ExternalID *string `json:"externalID"`

	// Id The ID of the device group. The ID must be unique across all OpenHES components.
	Id DeviceGroupID `json:"id"`

	// Name The name of the device group.
	Name string `json:"name"`
}

// DeviceID The ID of the device. The ID must be unique across all OpenHES components.
type DeviceID = openapi_types.UUID

// DeviceSchema Schema that describes device details.
type DeviceSchema struct {
	// Attributes Schema that describes a set of attributes and their values. The key is the property name and the value is the property value. The value can be of type string, integer, number, boolean, binary, or null.
	Attributes externalRef0.AttributesSchema `json:"attributes"`

	// CommunicationUnitID The list of communication unit IDs.
	CommunicationUnitID CommunicationUnitListID `json:"communicationUnitID"`

	// ExternalID The public ID of the device.
	ExternalID *string `json:"externalID"`

	// Id The ID of the device. The ID must be unique across all OpenHES components.
	Id DeviceID `json:"id"`

	// Name The name of the device.
	Name string `json:"name"`

	// Timezone The timezone of the device.
	Timezone *string `json:"timezone"`
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xXXW/rNg/+K4Le92IDHLtnRXfhq3XtsAXIQQO0AwacFYUsM7HO9HUkuU1W5L8PkhXn",
	"w85n23NRNKBI6iH5kKJfMVVCKwnSWZy/YksrECT8vFFC1JJR4piSf0rmhrdeXIKlhmkvxDl+qAANb5Ga",
	"IFcBousmqJbMpShqiNo6VIAXfqsBEWqUtYhwju40yD9+u0crIOnfEid4oowgDue4rlmJE+zmGnCOrTNM",
	"TvEi6QIcMbsTJGfWeZhdiGh4G29kDkQI/f8GJjjH/8tWmLKYmawvLYsWHTGGzHvB3Qf7LrhGjlxFHGpO",
	"CrB9MEtwhPEIVRulwTgGAS9VUgL1qkM5UYciUBok0ezJ/6/APn1VRfi7ab08zDVEvIsEw8yBkYTvSq2u",
	"C87ofhoE0LLmnBQccO5MDT0VZeWZ2ZdEQD84f3IIl2ByBHLqKpx/6qBaJNjAt5oZKHH+BQcqhuuS7bQ/",
	"traq+ArUeWC38Mwo/G5UrQ+3TxmU0dRrf2TjrIE6jZXrAPfx8WTGbET+XlzZzP3RLOliOY8fu+lwLBM+",
	"ngNnlX9P4YlzhhW1A3vuELpuPawGEO1/iU6aE/FtOGecxWK8Ly1PZ+QRXEywYwL+VXKH0+XpGZHtHYPd",
	"8iTrVOjrhKPr71lVlsz7Jny8yTY5v5vg/MtrJw+thEkHUzDrIlmLYlNSKMWBSLx43E7DIjmqMQiyELaL",
	"VdCIyNLnmBn0THgNtmnmf2COmA3Jj50zb0od1RvljkqQNh4aBUqkHwi+jnMNqAk8QTHcBDVBJiiGlqCC",
	"SWLmCVIG+RCbmm9WJcGzwVQNlpvMMpR0VZJ1lQETWhnn66CJpyOeMlfVRUqVyOi8MMpyUtisAjv4aUA0",
	"a3s+izXP2isCvY5bS8aVkjBi8sTRFbLpTRFnEtAPQpUgfkSrF7xvmkWm9E+J4KzROKNPtVL8qZkVXd8B",
	"HPIqaHh75Dz3gqdg6H0KMmOiFuFuwWT7e7srtrp6zUmyjH6F9ZQ27t0l81esJMSWffuG+kD1ULevxHt4",
	"3CbX+3i9B8MI/6xmZOn28fjMbduewPfPd39dH2C4H91EluPYyD00V8btY/mSZz9fXV1e7edagkviyHe6",
	"qlJ2xzXDMSJlacBaPwv9LV5313O7q89m5O1tNiM4Al1LTbJRk/Nbbr05TuDMw804G44PsObDkqu/DzlC",
	"/h3Vby2hd9FWUPfXyxuy+FFOlXSEhghBEMaD/4n6ZfVa+sdzuVjl+CbI0YgUFjkg/qQ23qpyTts8y15e",
	"XtJN44zJEmYDkGnlBMed/cXvxUqi6/EwvDLcookyy68InGDOKEgb1scI4tf70eBTetG529PPqtpQSJWZ",
	"ZtHQZkv9sIw6Dkdc+gzGNvAu0mgaSY5zfJlepJc+v8RVFud+eVn8FwAA//9Ho+z8thIAAA==",
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

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(path.Dir(pathToFile), "../job/job.yaml")) {
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
