package encryption

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var aes = NewAESEncryptionService([]byte("1111111111111111"))
var plainJSON = []byte(`{"type":"login","user_name":"test","ping_interval":45000,"ppks":[{"ppk_num":0,"pwd":"0","license_key":[0,0,0,0,0,0]}]}`)

func TestAesEncryptionService_Encode(t *testing.T) {
	t.Run("plain JSON string", func(t *testing.T) {
		encoded, err := aes.Encode(plainJSON)

		assert.NoError(t, err)
		assert.NotEmpty(t, encoded)
	})

	t.Run("bytes string", func(t *testing.T) {
		encoded, err := aes.Encode([]byte(`hello!`))

		assert.NoError(t, err)
		t.Log("encoded 2", string(encoded))
	})

	t.Run("same bytes string again", func(t *testing.T) {
		encoded, err := aes.Encode([]byte(`hello!`))

		assert.NoError(t, err)
		t.Log("encoded 3", string(encoded))
	})
}

func TestAesEncryptionService_Decode(t *testing.T) {
	t.Run("string encoded with Encode func", func(t *testing.T) {
		decoded, err := aes.Decode([]byte(``))

		assert.NoError(t, err)
		assert.Equal(t, string(plainJSON), string(decoded))
	})

	t.Run("string encoded with www.browserling.com", func(t *testing.T) {
		encoded, err := aes.Decode([]byte(`U2FsdGVkX1+6/f81ec3nfCjVRv8f133JtC0592XC6xg=`))

		assert.NoError(t, err)
		assert.Equal(t, "test", string(encoded))
	})
}
