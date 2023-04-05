package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mode Mode `yaml:"mode"`
}

func Load(filename string, c *Config) error {
	setDefaults(c)

	d, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(d, c); err != nil {
		return fmt.Errorf("cannot unmarshal config %s: %w", filename, err)
	}
	return nil
}

func setDefaults(c *Config) {
	c.Mode = ModeProd
}

type Mode string

const (
	ModeProd Mode = "prod"
	ModeDev  Mode = "dev"
)
