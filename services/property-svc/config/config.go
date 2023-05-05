package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mode  Mode  `yaml:"mode"`
	Proxy Proxy `yaml:"proxy"`
}

func (c Config) Validate() error {
	switch c.Mode {
	case ModeProd, ModeDev:
		break
	default:
		return fmt.Errorf("unknown mode: %v", c.Mode)
	}

	switch c.Proxy {
	case ProxyNone, ProxyDapr:
		break
	default:
		return fmt.Errorf("unknown proxy: %v", c.Proxy)
	}

	return nil
}

func LoadByFilename(filename string, c *Config) error {
	b, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("read config from file %s: %w", filename, err)
	}

	if err := Load(b, c); err != nil {
		return fmt.Errorf("load config from file %s: %w", filename, err)
	}

	return nil
}

func Load(b []byte, c *Config) error {
	setDefaults(c)

	if err := yaml.UnmarshalStrict(b, c); err != nil {
		return fmt.Errorf("cannot unmarshal config: %w", err)
	}

	if err := c.Validate(); err != nil {
		return fmt.Errorf("validate config: %w", err)
	}

	return nil
}

func setDefaults(c *Config) {
	c.Mode = ModeProd
	c.Proxy = ProxyDapr
}

type Mode string

const (
	ModeProd Mode = "prod"
	ModeDev  Mode = "dev"
)

type Proxy string

const (
	ProxyNone Proxy = "none"
	ProxyDapr Proxy = "dapr"
)
