package iface

import (
	"testing"
)

func TestMockBoolean(t *testing.T) {
	var iface Iface = &mockIface{
		calls: map[string][][]interface{} {
			"Boolean": {
				{},
				{},
			},
		},
		returns: map[string][][]interface{} {
			"Boolean": {
				{true},
				{false},
			},
		},
		stubs: map[string]func() []interface{} {
		},
	}

	if got, want := iface.Boolean(), true; got != want {
		t.Errorf("bool: got %v, want %v", got, want)
	}

	if got, want := iface.Boolean(), false; got != want {
		t.Errorf("bool: got %v, want %v", got, want)
	}
}
