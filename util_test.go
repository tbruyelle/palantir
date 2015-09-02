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

	assert.True(reg.HasExpired(2))
	assert.False(reg.HasExpired(20))
	assert.False(reg.HasExpired(5))
}
