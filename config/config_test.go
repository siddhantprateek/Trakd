package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_ENV_VAR", "test_value")

	value := GetEnv("TEST_ENV_VAR")
	assert.Equal(t, "test_value", value)
	value = GetEnv("NON_EXISTENT_VAR")
	assert.Equal(t, "", value)
	os.Unsetenv("TEST_ENV_VAR")
}
