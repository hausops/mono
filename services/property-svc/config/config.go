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
	Store Store `yaml:"store"`
}

func (c Config) Validate() error {
	switch c.Mode {
	case ModeProd, ModeDev:
		break
	default:
		return fmt.Errorf("unknown mode: %v", c.Mode)
	}

	switch c.Store {
	case StoreLocal, StoreMongo:
		break
	default:
		return fmt.Errorf("unknown store: %v", c.Store)
	}

	return nil
}

// LoadFromFile loads a YAML configuration from filename and decode it to c.
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

// Load decodes YAML data from r to c, sets defaults for missing fields,
// and performs validation.
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
	c.Store = StoreMongo
}

type Mode string

const (
	ModeProd Mode = "prod"
	ModeDev  Mode = "dev"
)

type Store string

const (
	StoreLocal Store = "local"
	StoreMongo Store = "mongo"
)
