package config_test

import (
	"os"
	"strings"
	"testing"

	"github.com/curated/elastic/config"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	c := config.New()
	assert.NotEmpty(t, c.Elastic.URL)
	assert.NotEmpty(t, c.Elastic.Username)
	assert.NotEmpty(t, c.Elastic.Password)
}

func TestConfigOverride(t *testing.T) {
	os.Setenv("CONFIG", "config/config.json")
	c := config.New()
	assert.Empty(t, c.Elastic.URL)
	assert.Empty(t, c.Elastic.Username)
	assert.Empty(t, c.Elastic.Password)
}

func TestGetPath(t *testing.T) {
	main := config.GetPath("main.go")
	ct := config.GetPath("config/config_test.go")
	assert.True(t, strings.Contains(main, "/src/github.com/curated/elastic/main.go"))
	assert.True(t, strings.Contains(ct, "/src/github.com/curated/elastic/config/config_test.go"))
}
