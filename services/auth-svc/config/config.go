package config

import (
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Mode mode
}

func (c *Config) Validate() error {
	switch c.Mode {
	case ModeProd, ModeDev:
	default:
		return fmt.Errorf("unknown mode: %v", c.Mode)
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
	decoder := yaml.NewDecoder(r)
	decoder.KnownFields(true)
	if err := decoder.Decode(c); err != nil && !errors.Is(err, io.EOF) {
		return fmt.Errorf("decode config: %w", err)
	}
	return nil
}

type mode string

const (
	ModeProd mode = "prod"
	ModeDev  mode = "dev"
)
