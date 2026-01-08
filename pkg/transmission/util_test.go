package transmission

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructJSONFields(t *testing.T) {
	type Foo struct {
		Foo string `json:"foo"`
		Bar int    `json:"bar"`
	}

	t.Run("2 struct fields", func(t *testing.T) {
		assert.Equal(t, []string{"foo", "bar"}, structJSONFields[Foo]())
	})

	t.Run("not a struct", func(t *testing.T) {
		assert.Panics(t, func() { structJSONFields[int]() })
	})

	t.Run("embedded struct", func(t *testing.T) {
		type Baz struct {
			Foo
			Baz bool `json:"baz"`
			Bas int  `json:"bas"`
		}

		assert.Equal(t, []string{"foo", "bar", "baz", "bas"}, structJSONFields[Baz]())
	})

	t.Run("unexported fields", func(t *testing.T) {
		type Unexported struct {
			Foo string `json:"foo"`
			bar int
		}
		assert.Equal(t, []string{"foo"}, structJSONFields[Unexported]())
	})

	t.Run("without JSON tag", func(t *testing.T) {
		type NoTag struct {
			Foo string
		}
		assert.Equal(t, []string{"Foo"}, structJSONFields[NoTag]())
	})

	t.Run("nested structs", func(t *testing.T) {
		type Nested struct {
			Abc Foo `json:"abc"`
		}
		assert.Equal(t, []string{"abc"}, structJSONFields[Nested]())
	})
}
