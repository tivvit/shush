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
	c.Backend.Badger = nil
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
	if c.Shortener.DefaultHashAlgo != "fnv32" {
		t.Error("Default hash algo is not fnv32")
	}
	if c.Shortener.DefaultShortener != "generator" {
		t.Error("Default shortener is not generator")
	}
	if c.Shortener.GenMaxRetries != 10 {
		t.Error("Default max retries is not 10")
	}
	if c.Shortener.DefaultLen != 5 {
		t.Error("Default len is not 5")
	}
	if c.Shortener.Maxlen != 50 {
		t.Error("Max len is not 50")
	}
	if !c.Shortener.AllowedHashAlgo["md5"] {
		t.Error("md5 is not allowed by default")
	}
	if !c.Shortener.AllowedShorteners["hash"] {
		t.Error("hash is not allowed by default")
	}
	if !c.Shortener.AllowedShorteners["generator"] {
		t.Error("generator is not allowed by default")
	}
}
