package iface

import (
	"testing"
)

func TestMockBoolean(t *testing.T) {
	iface := newMockIface().
		ExpectBoolean().AndReturn(true).
		ExpectBoolean().AndReturn(false).
		StubBooleanReturn(true)

	if got, want := iface.Boolean(), true; got != want {
		t.Errorf("bool: got %v, want %v", got, want)
	}

	if got, want := iface.Boolean(), false; got != want {
		t.Errorf("bool: got %v, want %v", got, want)
	}

	if got, want := iface.Boolean(), true; got != want {
		t.Errorf("bool: got %v, want %v", got, want)
	}

	if got, want := iface.Boolean(), true; got != want {
		t.Errorf("bool: got %v, want %v", got, want)
	}
}
