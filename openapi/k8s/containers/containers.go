// Package containers provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package containers

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/runtime"
)

// ContainerResourcesCpuNumSchema CPU resource value in numeric form. Represents number of cores, or millicores if the value is suffixed with 'm'.
type ContainerResourcesCpuNumSchema = float32

// ContainerResourcesCpuStrSchema CPU resource value in string form. Represents number of cores, or millicores if the value is suffixed with 'm'.
type ContainerResourcesCpuStrSchema = string

// ContainerResourcesMemorySchema Memory resource value. The value is suffixed with a unit of measurement, such as 'Mi' for mebibytes.
type ContainerResourcesMemorySchema = string

// ContainerResourcesSchema CPU and memory resource values.
type ContainerResourcesSchema struct {
	Cpu ContainerResourcesSchema_Cpu `json:"cpu"`

	// Memory Memory resource value. The value is suffixed with a unit of measurement, such as 'Mi' for mebibytes.
	Memory ContainerResourcesMemorySchema `json:"memory"`
}

// ContainerResourcesSchema_Cpu defines model for ContainerResourcesSchema.Cpu.
type ContainerResourcesSchema_Cpu struct {
	union json.RawMessage
}

// ContainerResourcesSetSchema Resource values for a container.
type ContainerResourcesSetSchema struct {
	// Limits CPU and memory resource values.
	Limits ContainerResourcesSchema `json:"limits"`

	// Requests CPU and memory resource values.
	Requests *ContainerResourcesSchema `json:"requests,omitempty"`
}

// AsContainerResourcesCpuStrSchema returns the union data inside the ContainerResourcesSchema_Cpu as a ContainerResourcesCpuStrSchema
func (t ContainerResourcesSchema_Cpu) AsContainerResourcesCpuStrSchema() (ContainerResourcesCpuStrSchema, error) {
	var body ContainerResourcesCpuStrSchema
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContainerResourcesCpuStrSchema overwrites any union data inside the ContainerResourcesSchema_Cpu as the provided ContainerResourcesCpuStrSchema
func (t *ContainerResourcesSchema_Cpu) FromContainerResourcesCpuStrSchema(v ContainerResourcesCpuStrSchema) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContainerResourcesCpuStrSchema performs a merge with any union data inside the ContainerResourcesSchema_Cpu, using the provided ContainerResourcesCpuStrSchema
func (t *ContainerResourcesSchema_Cpu) MergeContainerResourcesCpuStrSchema(v ContainerResourcesCpuStrSchema) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

// AsContainerResourcesCpuNumSchema returns the union data inside the ContainerResourcesSchema_Cpu as a ContainerResourcesCpuNumSchema
func (t ContainerResourcesSchema_Cpu) AsContainerResourcesCpuNumSchema() (ContainerResourcesCpuNumSchema, error) {
	var body ContainerResourcesCpuNumSchema
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContainerResourcesCpuNumSchema overwrites any union data inside the ContainerResourcesSchema_Cpu as the provided ContainerResourcesCpuNumSchema
func (t *ContainerResourcesSchema_Cpu) FromContainerResourcesCpuNumSchema(v ContainerResourcesCpuNumSchema) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContainerResourcesCpuNumSchema performs a merge with any union data inside the ContainerResourcesSchema_Cpu, using the provided ContainerResourcesCpuNumSchema
func (t *ContainerResourcesSchema_Cpu) MergeContainerResourcesCpuNumSchema(v ContainerResourcesCpuNumSchema) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JSONMerge(t.union, b)
	t.union = merged
	return err
}

func (t ContainerResourcesSchema_Cpu) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ContainerResourcesSchema_Cpu) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RUWWvbQBD+K8s0YCiyjoQWopcepqSlcRPi9Mm4sJZH8RTt0T3iGKP/XlZy5NixmwP6",
	"JntnvvmOnV1BoYRWEqWzkK/AFnMUvPkcKOk4STRXaJU3BdqB9j+8GDUloWKGtjCkHSkJOQwufzKzLmW3",
	"vPLISDLpBRoqWKmMiNkVaoM2DAsHUzRMlaxQBm3ElGGCqoqan4xK5uYdjmXWlyXd4YwtyM1ZT/RiiADv",
	"uNAVQp5FEAZwBzmUleIOIhAkSXgBeRqnaRaBW2qEHNq5UEf7BY6ceZlA6wzJm/+sD7I0FRCB5s6hCWR+",
	"jbP+6WSc9k8nb8WHI+j0tXz26xuiUGZ5SF97uiMxZteHaXLmJbmgUSC33qBA6SJmfTFn3LLekHrBGCZw",
	"StOlQ7st6vjd+yEdVDX+Pjy7vpzQM7X9KzUuZ0zsU7dNaAWF9huz246OZx2BNkqjcYS2K16BknhRQj5e",
	"wZHBEnJ4k2x2KlkvVPLEZaujV7RvlrGe1Bu+L0XauhV1HYHBP54MziAfNyI76EmXhJr+xsIdSALdoTCu",
	"tt1vLgdnxT3EgzRs42hFgtzG7OxBKNlZE0ngivZBzaH0gsM7cHAMzwDM9mBNdu/CBvll3nfxb819HcZO",
	"cmtKjzMLhSRL1SgMaIULnyg4VZA3Rx+L5dQoW/GpjQsV7JRcBIxB8z8751PLHPJw4k3omjunbZ4ki8Ui",
	"3m5OSM7wro8ynjtRBak7C6qEUJJ9uvzGhJph1d6KC43y65cRRFBRgdI267km8Xl03s/i9NFspVG2nsTK",
	"3CTrRpvc19cROHLN0/PU0Fs0tqWXxuvWgM41QQ4ncRqftO/W3EIufVXVfwMAAP//oc6Ge0kHAAA=",
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
