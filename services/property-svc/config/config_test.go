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
			Store: config.StoreMongo,
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
			Store: config.StoreMongo,
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
				Mode:  config.ModeProd,
				Store: config.StoreLocal,
			},
			wantErr: false,
		},
		{
			name: "invalid mode",
			config: config.Config{
				Mode:  "invalid",
				Store: config.StoreLocal,
			},
			wantErr: true,
		},
		{
			name: "invalid proxy",
			config: config.Config{
				Mode:  config.ModeProd,
				Store: "invalid",
			},
			wantErr: true,
		},
		{
			name: "invalid mode and proxy",
			config: config.Config{
				Mode:  "invalid",
				Store: "invalid",
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
