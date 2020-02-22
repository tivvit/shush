package config

import (
	"github.com/tivvit/shush/shush/config/backend"
	"testing"
)

func TestConf(t *testing.T) {
	c, err := NewConf("")
	if err != nil {
		t.Errorf("Default conf is invalid: %s", err)
	}
	err = c.validate()
	if err != nil {
		t.Error("Validity not checked in constructor")
	}
	c.Backend.InMem = nil
	err = c.validate()
	if err == nil {
		t.Error("0 backends allowed")
	}
	c.Backend.InMem = &backend.InMem{}
	c.Backend.JsonFile = &backend.JsonFile{}
	err = c.validate()
	if err == nil {
		t.Error(">1 backend allowed")
	}
}