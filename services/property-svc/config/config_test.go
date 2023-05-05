package config_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hausops/mono/services/property-svc/config"
)

func TestLoad(t *testing.T) {
	t.Run("set defaults", func(t *testing.T) {
		var c config.Config

		err := config.Load(bytes.NewReader([]byte("")), &c)
		if err != nil {
			t.Fatalf(`Load(""); unexpected error: %v`, err)
		}

		want := config.Config{
			Mode:  config.ModeProd,
			Proxy: config.ProxyDapr,
		}
		if diff := cmp.Diff(want, c); diff != "" {
			t.Errorf(`Load(""); (-want +got)\n%s`, diff)
		}
	})

	t.Run("valid config", func(t *testing.T) {
		var c config.Config

		configData, err := os.ReadFile("testdata/TestLoad/valid_config.yaml")
		if err != nil {
			t.Fatalf("read testdata: %v", err)
		}

		err = config.Load(bytes.NewReader(configData), &c)
		if err != nil {
			t.Fatalf("Load(...); unexpected error: %v", err)
		}

		want := config.Config{
			Mode:  config.ModeDev,
			Proxy: config.ProxyNone,
		}
		if diff := cmp.Diff(want, c); diff != "" {
			t.Errorf("Load(...); (-want +got)\n%s", diff)
		}
	})

	t.Run("invalid config", func(t *testing.T) {
		var c config.Config

		configData, err := os.ReadFile("testdata/TestLoad/invalid_config.yaml")
		if err != nil {
			t.Fatalf("read testdata: %v", err)
		}

		err = config.Load(bytes.NewReader(configData), &c)
		if err == nil {
			t.Fatalf("Load(...); expected error, but got nil")
		}
	})
}
