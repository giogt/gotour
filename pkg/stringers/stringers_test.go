package stringers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IPAddr_String(t *testing.T) {
	loopback := IPAddr{127, 0, 0, 1}
	googleDNS := IPAddr{8, 8, 8, 8}

	assert.Equal(t, "127.0.0.1", loopback.String())
	assert.Equal(t, "8.8.8.8", googleDNS.String())

	// fmt functions use the Stringer interface [method String()]
	assert.Equal(t, "127.0.0.1", fmt.Sprint(loopback))
	assert.Equal(t, "8.8.8.8", fmt.Sprint(googleDNS))
}
