package config

import (
	"os"
	"testing"
)

type testConfig struct {
	TestField string `yaml:"test"`
}

func TestNewConfigLoader_Default(t *testing.T) {
	config := NewConfigLoader()

	if config.env != "development" {
		t.Errorf("Expected env to equal 'development', got %q", config.env)
	}
}

func TestNewConfigLoader_WithSetEnv(t *testing.T) {
	os.Setenv("GO_ENV", "production")

	config := NewConfigLoader()

	if config.env != "production" {
		t.Errorf("Expected env to equal 'production', got %q", config.env)
	}
}

func TestLoader_OK(t *testing.T) {
	os.Setenv("GO_ENV", "test")

	filenames := make(map[string]string)
	filenames["test"] = "test.yaml"

	mock := &testConfig{}

	config := NewConfigLoader()

	if err := config.Load(&mock, filenames); err != nil {
		t.Error("Expected err to equal null")
	}

	if mock.TestField != "example value" {
		t.Errorf("Expected %q, got %q", "example value", mock.TestField)
	}
}

func TestLoad_NoConfigFile(t *testing.T) {
	os.Setenv("GO_ENV", "test")

	filenames := make(map[string]string)
	filenames["test"] = "no_such_file.yaml"

	mock := &testConfig{}

	config := NewConfigLoader()

	if err := config.Load(&mock, filenames); err != ErrReadFile {
		t.Errorf("Expected err to equal %q, got %q", ErrReadFile, err)
	}
}

func TestLoad_MissingKey(t *testing.T) {
	os.Setenv("GO_ENV", "test")

	filenames := make(map[string]string)
	filenames["invalid_key"] = "test.yaml"

	mock := &testConfig{}

	config := NewConfigLoader()

	if err := config.Load(&mock, filenames); err != ErrMissingKey {
		t.Errorf("Expected err to equal %q, got %q", ErrMissingKey, err)
	}
}

func TestLoad_InvalidConfigFile(t *testing.T) {
	os.Setenv("GO_ENV", "test")

	filenames := make(map[string]string)
	filenames["test"] = "invalid.yaml"

	mock := &testConfig{}

	config := NewConfigLoader()

	if err := config.Load(&mock, filenames); err != ErrCouldNotParse {
		t.Errorf("Expected err to equal %q, got %q", ErrCouldNotParse, err)
	}
}
