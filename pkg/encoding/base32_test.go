package encoding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase32Encode(t *testing.T) {
	plain := `wes samestine ati iki nelongso
        wong seng tak tresnani mblenjani janji
        opo ora eling naliko semono
        kebak kembang wangi jeroning dodo`
	expectedEncoded := "O5SXGIDTMFWWK43UNFXGKIDBORUSA2LLNEQG4ZLMN5XGO43PBIQCAIBAEAQCAIDXN5XGOIDTMVXGOIDUMFVSA5DSMVZW4YLONEQG2YTMMVXGUYLONEQGUYLONJUQUIBAEAQCAIBAEBXXA3ZAN5ZGCIDFNRUW4ZZANZQWY2LLN4QHGZLNN5XG6CRAEAQCAIBAEAQGWZLCMFVSA23FNVRGC3THEB3WC3THNEQGUZLSN5XGS3THEBSG6ZDP"

	encoded := Base32Encode(plain)
	assert.Equal(t, expectedEncoded, encoded)
}

func TestBase32DecodeSuccess(t *testing.T) {
	validBase32 := "O5SXGIDTMFWWK43UNFXGKIDBORUSA2LLNEQG4ZLMN5XGO43PBIQCAIBAEAQCAIDXN5XGOIDTMVXGOIDUMFVSA5DSMVZW4YLONEQG2YTMMVXGUYLONEQGUYLONJUQUIBAEAQCAIBAEBXXA3ZAN5ZGCIDFNRUW4ZZANZQWY2LLN4QHGZLNN5XG6CRAEAQCAIBAEAQGWZLCMFVSA23FNVRGC3THEB3WC3THNEQGUZLSN5XGS3THEBSG6ZDP"
	expectedResult := `wes samestine ati iki nelongso
        wong seng tak tresnani mblenjani janji
        opo ora eling naliko semono
        kebak kembang wangi jeroning dodo`

	result, err := Base32Decode(validBase32)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestBase32DecodeFailed(t *testing.T) {
	invalidBase32 := "Ksksjdd92832382-8ds-9f8sdjflajldfkj329832038702sjdfksj"

	_, err := Base32Decode(invalidBase32)
	assert.NotNil(t, err)
}
