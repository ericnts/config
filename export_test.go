package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	assert.Equal(t, Options.Name, "config")
	assert.Equal(t, Options.TcpPort, uint16(8081))
	assert.Equal(t, Options.LogResponse, true)
}
