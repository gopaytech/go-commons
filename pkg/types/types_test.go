package types

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyValueMap(t *testing.T) {
	t.Run("scan accept nil should failed", func(t *testing.T) {
		result := KeyValueMap{}
		err := result.Scan(nil)
		assert.Error(t, err)
	})

	t.Run("scan accept non string should failed", func(t *testing.T) {
		result := KeyValueMap{}
		err := result.Scan(90)
		assert.Error(t, err)
	})

	t.Run("from empty map", func(t *testing.T) {
		origin := KeyValueMap{}
		stringKV, err := origin.Value()
		assert.NoError(t, err)
		assert.Equal(t, stringKV, "")
	})

	t.Run("from non empty map", func(t *testing.T) {
		origin := KeyValueMap{
			"id":   faker.Username(),
			"name": faker.FirstNameMale(),
			"url":  faker.URL(),
		}
		stringKV, err := origin.Value()
		assert.NoError(t, err)
		assert.NotEmpty(t, stringKV)

		reverse := &KeyValueMap{}
		err = reverse.Scan(stringKV)
		assert.NoError(t, err)

		assert.Equal(t, origin, *reverse)
	})

	t.Run("given empty string, expect empty map", func(t *testing.T) {
		reverse := &KeyValueMap{}
		err := reverse.Scan("")
		assert.NoError(t, err)
		assert.Empty(t, reverse)
	})

}
