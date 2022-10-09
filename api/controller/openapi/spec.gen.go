// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7RWz27TThB+lWp+v6Mbu/xT5RPij1DFoUiop6pCW3uabLF3l9lJ1SryIRRx48idE+IJ",
	"OPVt/CJodp2mSQxENDklXu98833fzsx6AoWtnTVo2EM+AadI1chI4enIIx2U8k8byMEpHkECRtUIOYw9",
	"0jtdQgKEH8aasIScaYwJ+GKEtZKw/wnPIIf/0nmWNL716UEJTdPMdod8zwkVo2R9Izz+hhA2knUemgQi",
	"T7xUtasQ8r0E+MoJT20Yh0iQwOXu0O52q2Nt+MkjCTxy5T8nlQcJUFV1eAb58QQcWYfEGoMgXa7lwl0L",
	"jyXopEnWZnFyq9SenmPBM17xdb5MCWulqwWroFT6vR54xfZptzgobA23uJ5Jm6HgxqO/G9tOv7bTH+30",
	"czv91k4/tdMvq3FL+gJI0hH5Hfse5uPZsmas/TrnJGAduiJSVytUIuQqB9mnzZmVJKw5KK3RezXE3VOr",
	"SKr+AslrK32RDbLBnuSyDo1yGnJ4OMgGGSShZQLV9Jb9EFl+RJlibY0ULrxCjqqFn3fW+Kj5QZbJT2EN",
	"owlxyrlKFyEyPfeSf7Jmw8UEQVuJviDtOPI/fC3kH8dUi68ODCMZVe28RbpA2nlJZCnYyGroZxZKvYKz",
	"vkfYvKW7QYGen9nyamOqlmdG08RDXjBxb6Mm9nkYaZT3N7JJulpJJ92EbSJahYyr/r4I652/d6f3cb+O",
	"+Za0m+4yQPpqbhs18qfa3zT/rR73hlpGcTFaNWR+I93Pk8032/Jd2dtsW6qe8K0Q9kQnFmEqW6gKEhhT",
	"BTmMmF2epmFxZD3n+9l+lspoDtdlBF6GaK+/t9c37cef7fXN4mcONCfNrwAAAP//Lc2dPioJAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
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
	var res = make(map[string]func() ([]byte, error))
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
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
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

