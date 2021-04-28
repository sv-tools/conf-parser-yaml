package confyaml

import (
	"context"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// Parser is a conf.ParseFunc to parse the given yaml
func Parser(_ context.Context, r io.Reader) (interface{}, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var res interface{}
	if err := yaml.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	return res, nil
}
