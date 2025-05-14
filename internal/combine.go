package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Combine(spec Spec) (OpenAPISpec, error) {
	newSpec := spec.OpenAPISpec

	if newSpec.Paths == nil {
		newSpec.Paths = map[string]Operations{}
	}

	for _, include := range spec.Include {
		file, err := os.ReadFile(include.Ref)
		if err != nil {
			return newSpec, fmt.Errorf("read file %s: %w", include.Ref, err)
		}

		var includeSpec OpenAPISpec
		if err = yaml.Unmarshal(file, &includeSpec); err != nil {
			return newSpec, fmt.Errorf("unmarshal yaml from %s: %w", include.Ref, err)
		}

		for path, operations := range includeSpec.Paths {
			if _, exists := newSpec.Paths[path]; exists {
				return OpenAPISpec{}, fmt.Errorf("path %s from %s already contains in spec", path, include.Ref)
			}

			newSpec.Paths[path] = operations

			for method := range newSpec.Paths[path] {
				for code, defResp := range spec.DefaultResponses {
					if _, exists := newSpec.Paths[path][method].Responses[code]; exists {
						continue
					}

					newSpec.Paths[path][method].Responses[code] = defResp
				}
			}
		}

		for name, components := range includeSpec.Components {
			if _, exists := newSpec.Components[name]; !exists {
				newSpec.Components[name] = components
			} else {
				for k, component := range components {
					newSpec.Components[name][k] = component
				}
			}
		}
	}

	return newSpec, nil
}
