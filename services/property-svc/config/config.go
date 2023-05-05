package config

import (
	"errors"
	"fmt"
	"io"
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

func LoadFromFile(filename string, c *Config) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("open config file: %w", err)
	}
	defer f.Close()

	if err := Load(f, c); err != nil {
		return fmt.Errorf("load config %s: %w", filename, err)
	}

	return nil
}

func Load(r io.Reader, c *Config) error {
	setDefaults(c)

	decoder := yaml.NewDecoder(r)
	decoder.SetStrict(true)
	if err := decoder.Decode(c); err != nil && !errors.Is(err, io.EOF) {
		return fmt.Errorf("decode config: %w", err)
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
