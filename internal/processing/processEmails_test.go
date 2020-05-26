package processing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cleanEmail(t *testing.T) {
	_, err := cleanEmail("")
	assert.NotNil(t, err)

	_, err = cleanEmail("test@gmail.com@gmail.com")
	assert.NotNil(t, err)

	_, err = cleanEmail("@gmail.com")
	assert.NotNil(t, err)

	v, err := cleanEmail("test.email+spam@gmail.com")
	assert.Nil(t, err)
	assert.Equal(t, "testemail", v)

	v, _ = cleanEmail("test.email@gmail.com")
	assert.Equal(t, "testemail", v)

	v, _ = cleanEmail("testemail@gmail.com")
	assert.Equal(t, "testemail", v)
}

func BenchmarkCleanEmail(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = cleanEmail("test.email+spam@gmail.com")
	}
	// Note: to optimize performance - replace strings.* calls with bytes.WriteString calls
}
