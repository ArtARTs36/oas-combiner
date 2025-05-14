package internal

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Combine(spec Spec) (Spec, error) {
	newSpec := spec

	for _, include := range spec.Include {
		file, err := os.ReadFile(include.Ref)
		if err != nil {
			return newSpec, fmt.Errorf("read file %s: %w", include.Ref, err)
		}

		var includeSpec Spec
		if err = yaml.Unmarshal(file, &includeSpec); err != nil {
			return newSpec, fmt.Errorf("unmarshal yaml from %s: %w", include.Ref, err)
		}

		for path, operations := range includeSpec.Paths {
			if _, exists := newSpec.Paths[path]; exists {
				return Spec{}, fmt.Errorf("path %s from %s already contains in spec", path, include.Ref)
			}

			newSpec.Paths[path] = operations
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

	newSpec.Include = nil

	return newSpec, nil
}
