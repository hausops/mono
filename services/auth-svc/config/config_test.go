package config_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hausops/mono/services/auth-svc/config"
)

func TestLoad(t *testing.T) {
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
			Mode: config.ModeDev,
			Datastore: config.RedisDatastore{
				URI: "redis://test-redis-user:test-redis-password@localhost:6379/0",
			},
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
			t.Errorf("Load(...); want error, but got nil")
		}
	})
}

func TestConfig_Validate(t *testing.T) {
	tcs := []struct {
		name    string
		config  config.Config
		wantErr bool
	}{
		{
			name: "valid configuration",
			config: config.Config{
				Mode:      config.ModeProd,
				Datastore: config.LocalDatastore{},
			},
			wantErr: false,
		},
		{
			name: "invalid mode",
			config: config.Config{
				Mode:      "invalid",
				Datastore: config.LocalDatastore{},
			},
			wantErr: true,
		},
		{
			name: "invalid datastore",
			config: config.Config{
				Mode:      config.ModeProd,
				Datastore: config.RedisDatastore{},
			},
			wantErr: true,
		},
		{
			name: "invalid datastore (nil)",
			config: config.Config{
				Mode:      config.ModeProd,
				Datastore: nil,
			},
			wantErr: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.config.Validate()
			if tc.wantErr && err == nil {
				t.Fatalf("Validate(); want error, but got nil")
			} else if !tc.wantErr && err != nil {
				t.Errorf("Validate() error = %v, want no error", err)
			}
		})
	}
}
