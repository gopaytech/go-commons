package types

import (
	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyValueStrings(t *testing.T) {
	t.Run("from empty map", func(t *testing.T) {
		origin := map[string]string{}
		stringKV, err := FromKeyValueMap(origin)
		assert.NoError(t, err)
		assert.Equal(t, stringKV, "")
	})

	t.Run("from non empty map", func(t *testing.T) {
		origin := map[string]string{
			"id":   faker.Username(),
			"name": faker.FirstNameMale(),
			"url":  faker.URL(),
		}
		stringKV, err := FromKeyValueMap(origin)
		assert.NoError(t, err)
		assert.NotEmpty(t, stringKV)

		reverseMap := ToKeyValueMap(stringKV)
		assert.Equal(t, origin, reverseMap)
	})

	t.Run("given empty string, expect empty map", func(t *testing.T) {
		mapResult := ToKeyValueMap("")
		assert.Empty(t, mapResult)
	})
}
