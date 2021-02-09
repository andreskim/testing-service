package config

import (
	"os"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

// Testing sequence of configuration
func Test(t *testing.T) {
	s := "unittest"
	v := "TESTING_SERVICE_WORKDIR"
	os.Setenv(v, s)
	ReadConfig()
	o := GetString("Workdir")
	assert.EqualValues(t, o, s, "Misconfiguration detected")
}
