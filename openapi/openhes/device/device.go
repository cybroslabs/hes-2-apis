// Package device provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.2.0 DO NOT EDIT.
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
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RXW2/jNhP9KwN+36Ms7XaRPuip27hoDQRYA0mftkFBUWOLLS9aXpIYgf57QUqy5Ejx",
	"Jdl9skHNcM6cOcMhnwnTstYKlbMkfyaWVShp/PvZOcML79DexsWwRsuSO64VFWujazSOY7SlavdlQ/Kv",
	"z8TtaiQ5sc5wtSVNsl/hyuEWzXhJeVkcrhRaC6SKNPcJUV4IWggkuTMem4SUaJnhdYhPctKCAldRB+2X",
	"Ai1QsOhAb4Du0QNVJbgKuYEHKjzaFO4qhH9xB9yGD1C3uexAUYm9eWs8MYmr7Q6tAaMKCgwhQw7QJp5A",
	"l24CbZIJdKklUHBFzS4BbSCkmP6lSNLnr4t/kDmSkKfFVi+6xX0q6VCSscmCy1obF+pQU1eRnGy5q3yR",
	"Mi0ztiuMtoIWNqvQLn5a0JrbTNeoaM3jb4U224cgTZOQay2lV5zRQPWfirvVMmx+yH9gYLWMeVcIbOwC",
	"XnHXcrRagvTWBYa84t88AmVGWwtUCPhSo/rjt1sYJNiSsdFGUkdy4j0vB3L2mpoCvOH2VZCC26iIKURY",
	"LbuI3KGMQv6/wQ3Jyf+yAVPW9UQ2R0uzR0eNobtZcEP7nKPgGZglOspFB7U+aDymlUIWTFdqo09n0Fvf",
	"7WrscDUJwSeHRlHxGoW1LwRnx8sdwR327EzlePlGlkNrzoOLTXsCl+TqBtU2NMfHCaomIQa/eW6wJPlX",
	"EiUXwyUv6b1/2aix3GNO15VWeFnB4/ES3GCINlfq7rScL1D0by16Lkp84AzPyL9d+DvuEbaX9IlLL6Op",
	"5Gr//+VB/oK30SZJD/Y0YQNXWmE3Q86X8B2r9yq+xG9cpss8b9FwKnrX+2k+4+8XKMBGPxD8pA4iz635",
	"e6vV7XK6SgPPF6R0d70+kUqlrXtlsqyBlqVBa8OcDLsF23Grj+U9EXQ/D2c6RRt3rFF6Nn++uvp0dZzR",
	"rm8cq99bh7BF0rLRYZ8ryTLi/N1oX58eyG1SsA3WP3IUj0BdppAxwGMT7uLZdJD595pKh9yfPY+mWN42",
	"iV6Xw7lK+PEaeFP5jxR+uMGfKs7kpdKEyT17h73o5tHdKt9yQRqfKd9Nfpcr7723nymHybguU1mG7Xh3",
	"E2VaOcriOYySchFPwY3+ZXiShBdKHy4n13EdbmhhwSENX7wJXpVztc2z7PHxMT10zrgq8WmBKq2cFGTy",
	"SAwl1Qo+r1cgdYnCwkabXugkIYIzVDbS2YH49fZm8TH9MIkdXkpWe8Mw1WabdY426+3DNOBO4BlBH9DY",
	"Ft6HtHPt3mMkJ5/SNnp4yVmSB+00/wUAAP//wsuuXqUPAAA=",
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

	for rawPath, rawFunc := range externalRef0.PathToRawSpec(path.Join(pathPrefix, "../job/job.yaml")) {
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
