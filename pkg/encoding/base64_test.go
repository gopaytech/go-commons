package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64DecodeSuccess(t *testing.T) {
	validBase64 := "d2VzIHNhbWVzdGluZSBhdGkgaWtpIG5lbG9uZ3NvCiAgICAgICAgd29uZyBzZW5nIHRhayB0cmVzbmFuaSBtYmxlbmphbmkgamFuamkKICAgICAgICBvcG8gb3JhIGVsaW5nIG5hbGlrbyBzZW1vbm8KICAgICAgICBrZWJhayBrZW1iYW5nIHdhbmdpIGplcm9uaW5nIGRvZG8="
	expectedResult := `wes samestine ati iki nelongso
        wong seng tak tresnani mblenjani janji
        opo ora eling naliko semono
        kebak kembang wangi jeroning dodo`

	result, err := Base64Decode(validBase64)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestBase64DecodeByteSuccess(t *testing.T) {
	validBase64 := "d2VzIHNhbWVzdGluZSBhdGkgaWtpIG5lbG9uZ3NvCiAgICAgICAgd29uZyBzZW5nIHRhayB0cmVzbmFuaSBtYmxlbmphbmkgamFuamkKICAgICAgICBvcG8gb3JhIGVsaW5nIG5hbGlrbyBzZW1vbm8KICAgICAgICBrZWJhayBrZW1iYW5nIHdhbmdpIGplcm9uaW5nIGRvZG8="
	expectedResult := `wes samestine ati iki nelongso
        wong seng tak tresnani mblenjani janji
        opo ora eling naliko semono
        kebak kembang wangi jeroning dodo`

	result, err := Base64DecodeBytes(validBase64)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, string(result))
}

func TestBase64DecodeFailed(t *testing.T) {
	invalidBase64 := "Ksksjdd92832382-8ds-9f8sdjflajldfkj329832038702sjdfksj"

	_, err := Base64Decode(invalidBase64)
	assert.NotNil(t, err)
}

func TestBase64Encode(t *testing.T) {
	plain := `wes samestine ati iki nelongso
        wong seng tak tresnani mblenjani janji
        opo ora eling naliko semono
        kebak kembang wangi jeroning dodo`
	expectedEncoded := "d2VzIHNhbWVzdGluZSBhdGkgaWtpIG5lbG9uZ3NvCiAgICAgICAgd29uZyBzZW5nIHRhayB0cmVzbmFuaSBtYmxlbmphbmkgamFuamkKICAgICAgICBvcG8gb3JhIGVsaW5nIG5hbGlrbyBzZW1vbm8KICAgICAgICBrZWJhayBrZW1iYW5nIHdhbmdpIGplcm9uaW5nIGRvZG8="

	encoded := Base64Encode(plain)
	assert.Equal(t, expectedEncoded, encoded)
}
