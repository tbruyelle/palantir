package palantir

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRegistrationHasExpired(t *testing.T) {
	assert := assert.New(t)
	// Create a registration with 5-days-ago date
	reg := Registration{Date: time.Now().Unix() - 3600*24*5}

	reg.TryDuration = 2
	assert.True(reg.HasExpired())

	reg.TryDuration = 20
	assert.False(reg.HasExpired())

	reg.TryDuration = 5
	assert.False(reg.HasExpired())
}
