package config

import (
	"errors"
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Mode      mode
	Datastore datastore
}

func (c *Config) UnmarshalYAML(node *yaml.Node) error {
	var conf struct {
		Mode      mode                 `yaml:"mode"`
		Datastore map[string]yaml.Node `yaml:"datastore"`
	}

	if err := node.Decode(&conf); err != nil {
		return fmt.Errorf("decode config from YAML: %w", err)
	}

	// Set mode
	c.Mode = conf.Mode

	// Set datastore
	switch kind := conf.Datastore["kind"].Value; kind {
	case "mock":
		c.Datastore = MockDatastore{}
	case "mongo":
		var mongo MongoDatastore
		spec := conf.Datastore["spec"]
		if err := spec.Decode(&mongo); err != nil {
			return fmt.Errorf("decode datastore spec from YAML: %w", err)
		}
		c.Datastore = mongo
	default:
		return fmt.Errorf("unknown datastore kind: %s", kind)
	}

	if err := c.Validate(); err != nil {
		return fmt.Errorf("validate config: %w", err)
	}

	return nil
}

func (c *Config) Validate() error {
	switch c.Mode {
	case ModeProd, ModeDev:
	default:
		return fmt.Errorf("unknown mode: %v", c.Mode)
	}

	switch t := c.Datastore.(type) {
	case MockDatastore:
	case MongoDatastore:
		if t.URI == "" {
			return errors.New("mongo datastore missing URI field")
		}
	default:
		return fmt.Errorf("unknown datastore: %v", c.Datastore)
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

type datastore interface {
	isDatastore()
}

type MockDatastore struct{}

func (d MockDatastore) isDatastore() {}

type MongoDatastore struct {
	// URI is the mongo connection URI.
	URI string `yaml:"uri"`
}

func (d MongoDatastore) isDatastore() {}
