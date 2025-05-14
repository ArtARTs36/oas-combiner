package internal

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

type ShortSlice[T string | int | float32 | float64] struct {
	Values []T
}

func (s ShortSlice[T]) UnmarshalYAML(n *yaml.Node) error {
	if n.Kind != yaml.SequenceNode {
		return fmt.Errorf("yaml must contain a sequence node, has %v", n.Kind)
	}

	if s.Values == nil {
		s.Values = make([]T, 0)
	}

	for _, item := range n.Content {
		var val T

		if err := item.Decode(&val); err != nil {
			return err
		}

		s.Values = append(s.Values, val)
	}

	return nil
}

func (s ShortSlice[T]) MarshalYAML() (interface{}, error) {
	content := make([]*yaml.Node, 0, len(s.Values)*2)
	for _, val := range s.Values {
		content = append(content,
			&yaml.Node{Kind: yaml.ScalarNode, Value: fmt.Sprintf("%s", val)},
		)
	}

	return &yaml.Node{
		Kind:    yaml.MappingNode,
		Content: content,
	}, nil
}
