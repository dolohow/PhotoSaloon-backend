package common

import (
	"testing"
)

func TestJSONMsg_WithOutInterpolation(t *testing.T) {
	msg := JSONMsg("Test with string")
	expectedMsg := "{\"msg\":\"Test with string\"}"

	if msg != expectedMsg {
		t.Errorf("Expected `%s`, got `%s`", expectedMsg, msg)
	}
}

func TestJSONMsg_WithInterpolation(t *testing.T) {
	var msg string
	var expectedMsg string

	msg = JSONMsg("Test with string %s", "interpolation")
	expectedMsg = "{\"msg\":\"Test with string interpolation\"}"

	if msg != expectedMsg {
		t.Errorf("Expected `%s`, got `%s`", expectedMsg, msg)
	}

	msg = JSONMsg("Test with multiple %s %s", "string", "interpolation")
	expectedMsg = "{\"msg\":\"Test with multiple string interpolation\"}"

	if msg != expectedMsg {
		t.Errorf("Expected `%s`, got `%s`", expectedMsg, msg)
	}
}
