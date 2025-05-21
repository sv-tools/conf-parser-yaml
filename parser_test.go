package confyaml_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/sv-tools/conf"

	confyaml "github.com/sv-tools/conf-parser-yaml"
)

const data = `
foo: 42
bar: test
`

const wrongData = `
foo: 42
- bar: test
`

func TestParser(t *testing.T) {
	c := conf.New().WithReaders(conf.NewStreamParser(strings.NewReader(data)).WithParser(confyaml.Parser))
	require.NoError(t, c.Load(t.Context()))

	require.Equal(t, 42, c.GetInt("foo"))
	require.Equal(t, "test", c.Get("bar"))
}

var errFake = errors.New("fake error")

type testReader struct{}

func (t *testReader) Read(_ []byte) (int, error) {
	return 0, errFake
}

func TestParserErrors(t *testing.T) {
	c := conf.New().WithReaders(conf.NewStreamParser(&testReader{}).WithParser(confyaml.Parser))
	require.ErrorIs(t, c.Load(t.Context()), errFake)

	c = conf.New().WithReaders(conf.NewStreamParser(strings.NewReader(wrongData)).WithParser(confyaml.Parser))
	require.ErrorContains(t, c.Load(t.Context()), "yaml: line 1: did not find expected key")
}

func ExampleParser() {
	c := conf.New().WithReaders(conf.NewStreamParser(strings.NewReader(data)).WithParser(confyaml.Parser))
	if err := c.Load(context.Background()); err != nil {
		panic(err)
	}

	fmt.Println(c.GetInt("foo"))
	// Output: 42
}
