package config

import (
	"os"
	"testing"
)

func TestEnvConfig_GetInt(t *testing.T) {
	os.Setenv("TEST_INT", "123")
	defer os.Unsetenv("TEST_INT")

	e := &EnvConfig{}
	val := e.GetInt("TEST_INT")
	if val != 123 {
		t.Errorf("expected 123, got %d", val)
	}
}

func TestEnvConfig_GetInt64(t *testing.T) {
	os.Setenv("TEST_INT64", "1234567890123")
	defer os.Unsetenv("TEST_INT64")

	e := &EnvConfig{}
	val := e.GetInt64("TEST_INT64")
	if val != 1234567890123 {
		t.Errorf("expected 1234567890123, got %d", val)
	}
}

func TestEnvConfig_GetUint64(t *testing.T) {
	os.Setenv("TEST_UINT64", "1234567890123")
	defer os.Unsetenv("TEST_UINT64")

	e := &EnvConfig{}
	val := e.GetUint64("TEST_UINT64")
	if val != 1234567890123 {
		t.Errorf("expected 1234567890123, got %d", val)
	}
}

func TestEnvConfig_GetFloat64(t *testing.T) {
	os.Setenv("TEST_FLOAT64", "123.456")
	defer os.Unsetenv("TEST_FLOAT64")

	e := &EnvConfig{}
	val := e.GetFloat64("TEST_FLOAT64")
	if val != 123.456 {
		t.Errorf("expected 123.456, got %f", val)
	}
}

func TestEnvConfig_GetString(t *testing.T) {
	os.Setenv("TEST_STRING", "test_value")
	defer os.Unsetenv("TEST_STRING")

	e := &EnvConfig{}
	val := e.GetString("TEST_STRING")
	if val != "test_value" {
		t.Errorf("expected 'test_value', got '%s'", val)
	}
}

func TestEnvConfig_GetStringSlice(t *testing.T) {
	os.Setenv("TEST_STRING_SLICE", "val1,val2,val3")
	defer os.Unsetenv("TEST_STRING_SLICE")

	e := &EnvConfig{}
	val := e.GetStringSlice("TEST_STRING_SLICE")
	expected := []string{"val1", "val2", "val3"}
	for i, v := range val {
		if v != expected[i] {
			t.Errorf("expected '%s', got '%s'", expected[i], v)
		}
	}
}

func TestEnvConfig_GetStringMap(t *testing.T) {
	os.Setenv("TEST_STRING_MAP", "key1=val1,key2=val2")
	defer os.Unsetenv("TEST_STRING_MAP")

	e := &EnvConfig{}
	val := e.GetStringMap("TEST_STRING_MAP")
	expected := map[string]interface{}{"key1": "val1", "key2": "val2"}
	for k, v := range val {
		if v != expected[k] {
			t.Errorf("expected '%s', got '%s'", expected[k], v)
		}
	}
}