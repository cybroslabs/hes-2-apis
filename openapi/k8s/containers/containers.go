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

// ContainerResourcesCpuIntSchema CPU resource value in numeric form. Represents number of cores, or millicores if the value is suffixed with 'm'.
type ContainerResourcesCpuIntSchema = float32

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

// AsContainerResourcesCpuIntSchema returns the union data inside the ContainerResourcesSchema_Cpu as a ContainerResourcesCpuIntSchema
func (t ContainerResourcesSchema_Cpu) AsContainerResourcesCpuIntSchema() (ContainerResourcesCpuIntSchema, error) {
	var body ContainerResourcesCpuIntSchema
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContainerResourcesCpuIntSchema overwrites any union data inside the ContainerResourcesSchema_Cpu as the provided ContainerResourcesCpuIntSchema
func (t *ContainerResourcesSchema_Cpu) FromContainerResourcesCpuIntSchema(v ContainerResourcesCpuIntSchema) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContainerResourcesCpuIntSchema performs a merge with any union data inside the ContainerResourcesSchema_Cpu, using the provided ContainerResourcesCpuIntSchema
func (t *ContainerResourcesSchema_Cpu) MergeContainerResourcesCpuIntSchema(v ContainerResourcesCpuIntSchema) error {
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

	"H4sIAAAAAAAC/7RU62vbMBD/V8StEBiOHbdsUH/ZI4yurKGl6T6FDBTlXN+wHpPkpiH4fx+yU6dJk/UB",
	"++ZEd7/7PXRagdDSaIXKO8hW4ESBkjefQ608J4X2Gp2urEA3NNW58uOmJFTM0QlLxpNWkMHw6iez61J2",
	"x8sKGSmmKomWBMu1lTG7RmPRhWHhYIaW6ZwJbdFFTFsmqSyp+ckoZ77ocBxzVZ7TPc7ZgnzBerIXQwR4",
	"z6UpEbI0gjCAe8ggLzX3EIEkRbKSkA3iNAK/NAgZtFOhjvbLG3v7OnnOW1K3/1kdpIOBhAgM9x5tIPNr",
	"kvZPp5NB/3T6Xn46gk5fy2e/vhFKbZeH9LWnOxJjdnOYJmeVIh80SuSusihR+Yi5ShSMO9YbUS8YwyTO",
	"aLb06LZFHX/4OKKDqiY/Rmc3V1N6obZ/pcbVnMl96rYJrUCYamN229HxrCMwVhu0ntB1xSvQCi9zyCYr",
	"OLKYQwbvks1GJet1Sp65bHX0hvbNKtbTesP3tUhbt6KuI7D4pyKLc8gmjcgOetoloWe/UfgDSeDBF+J6",
	"2/3mcnAmHiAepeEaR0uS5Ddmp49CSc+aSAJXdI9qDqUXHN6Bg2N4AWC6B2u6exc2yK/zvot/a+7bMHaS",
	"W1N6mlkoJJXrRmFAEz58ouRUQtYcfRbLmdWu5DMXCx3sVFwGjGHzP7vgM8c88nBS2dBVeG9cliSLxSLe",
	"bk5IzfG+jyouvCyD1J0F1VJqxb5cnTOp51i2t+LSoPr+bQwRlCRQuWY91yS+ji/6aTx4MlsbVK0nsba3",
	"ybrRJQ/1dQSefPP0PDf0Dq1r6Q3idWtA54Ygg5N4EJ+071bhIFNVWdZ/AwAA//9G3QI5RwcAAA==",
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
