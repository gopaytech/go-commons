package strings

import (
	"fmt"
	"io/ioutil"
	"testing"

	fake "github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

func TestSaveToFileSuccess(t *testing.T) {
	source := `I had to meet you here today
There's just so many things to say
Please don't stop me till I'm through
This is something I hate to do`

	fileName := "/tmp/kissAndSayGoodby"
	sourceByte := []byte(source)
	err := ioutil.WriteFile(fileName, sourceByte, 0644)
	assert.Nil(t, err)

	stringByte, err := ioutil.ReadFile(fileName)
	assert.Nil(t, err)
	assert.Equal(t, source, string(stringByte))
}

func TestKVJoin(t *testing.T) {
	fake.Seed(0)
	key := fake.FirstName()
	value := fake.LastName()

	kv := KVJoin(key, value)
	assert.Equal(t, fmt.Sprintf("%s=%s", key, value), kv)
}

func TestKVSplit(t *testing.T) {
	fake.Seed(0)
	key := fake.FirstName()
	value := fake.LastName()
	kv := fmt.Sprintf("%s=%s", key, value)

	resultKey, resultValue := KVSplit(kv)
	assert.Equal(t, key, resultKey)
	assert.Equal(t, value, resultValue)
}

func TestKVSplitFailed(t *testing.T) {
	fake.Seed(0)
	key := fake.FirstName()
	value := fake.LastName()
	kv := fmt.Sprintf("%s%s", key, value)

	resultKey, resultValue := KVSplit(kv)
	assert.Equal(t, resultKey, kv)
	assert.Equal(t, resultValue, "")
}
