package confyaml

import (
	"context"
	"io"

	"gopkg.in/yaml.v3"
)

// Parser is a conf.ParseFunc to parse the given yaml
func Parser(ctx context.Context, r io.Reader) (interface{}, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var res interface{}
	if err := yaml.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, ctx.Err()
}
