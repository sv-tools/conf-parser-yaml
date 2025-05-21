package confyaml

import (
	"context"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

// Parser is a conf.ParseFunc to parse the given yaml
func Parser(_ context.Context, r io.Reader) (any, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %w", err)
	}

	var res any
	if err := yaml.Unmarshal(data, &res); err != nil {
		return nil, fmt.Errorf("failed to parse yaml: %w", err)
	}

	return res, nil
}
