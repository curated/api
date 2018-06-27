package config_test

import (
	"os"
	"testing"

	"github.com/curated/elastic/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	c := config.New("../")

	assert.Equal(t, "test", c.Env)
	assert.Equal(t, "../", c.Root)
	assert.NotEmpty(t, c.Elastic.URL)
	assert.NotEmpty(t, c.Elastic.Username)
	assert.NotEmpty(t, c.Elastic.Password)
}

func TestConfigOverride(t *testing.T) {
	orig := os.Getenv("CONFIG")
	os.Setenv("CONFIG", "config/config.sample.json")
	c := config.New("../")

	assert.Equal(t, "sample", c.Env)
	assert.Equal(t, "../", c.Root)
	assert.Empty(t, c.Elastic.URL)
	assert.Empty(t, c.Elastic.Username)
	assert.Empty(t, c.Elastic.Password)

	os.Setenv("CONFIG", orig)
}
