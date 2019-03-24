package param

import (
	"fmt"
	"testing"
)

func TestParameterValue_GetValue(t *testing.T) {
	tests := []struct {
		name   string
		secret bool
	}{
		{name: "secret", secret: true},
		{name: "clear", secret: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pv := NewParameterValue("val", test.secret)

			assertEqual(t, "val", pv.GetValue())
		})
	}
}

func TestParameterValue_IsSecret(t *testing.T) {
	tests := []struct {
		name   string
		secret bool
	}{
		{name: "secret", secret: true},
		{name: "clear", secret: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pv := NewParameterValue("val", test.secret)

			assertEqual(t, test.secret, pv.IsSecret())
		})
	}
}

func TestParameterValue_String(t *testing.T) {
	tests := []struct {
		name     string
		expected string
		secret   bool
	}{
		{name: "secret", expected: "*****", secret: true},
		{name: "clear", expected: "val", secret: false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			pv := NewParameterValue("val", test.secret)

			assertEqual(t, test.expected, fmt.Sprintf("%s", pv))
		})
	}
}