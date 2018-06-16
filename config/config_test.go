package config_test

import (
	"os"
	"testing"

	"github.com/curated/elastic/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	os.Setenv("CONFIG", "config/config.test.json")
	c := config.New("../")

	assert.NotEmpty(t, c.Elastic.URL)
	assert.NotEmpty(t, c.Elastic.Username)
	assert.NotEmpty(t, c.Elastic.Password)
}

func TestConfigOverride(t *testing.T) {
	os.Setenv("CONFIG", "config/config.json")
	c := config.New("../")

	assert.Empty(t, c.Elastic.URL)
	assert.Empty(t, c.Elastic.Username)
	assert.Empty(t, c.Elastic.Password)
}
