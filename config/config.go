package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/golang/glog"
)

const (
	configKey   = "CONFIG"
	defaultFile = "config/config.prod.json"
)

// Config values
type Config struct {
	Elastic struct {
		URL      string
		Username string
		Password string
	}
}

// New creates config from ENV or default file
func New() *Config {
	f := os.Getenv(configKey)
	if len(f) == 0 {
		f = defaultFile
	}

	b, err := ioutil.ReadFile(f)
	if err != nil {
		glog.Fatalf("Failed loading '%s' with error: %v", f, err)
	}

	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		glog.Fatalf("Failed parsing '%s' with error: %v", f, err)
	}

	return &c
}
