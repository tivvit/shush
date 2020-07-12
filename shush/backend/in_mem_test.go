package backend

import (
	"testing"
	"time"
)

func TestInMem(t *testing.T) {
	im := NewInMem()
	err := im.Set("a", "a", time.Duration(1*time.Second))
	if err != nil {
		t.Error(err)
	}
	err = im.Set("b", "b", time.Duration(2*time.Second))
	if err != nil {
		t.Error(err)
	}
	err = im.Set("c", "c", time.Duration(0))
	if err != nil {
		t.Error(err)
	}
	v, err := im.Get("a")
	if err != nil {
		t.Error(err)
	}
	if v != "a" {
		t.Error("invalid value for a")
	}
	v, err = im.Get("b")
	if err != nil {
		t.Error(err)
	}
	if v != "b" {
		t.Error("invalid value for b")
	}
	v, err = im.Get("c")
	if err != nil {
		t.Error(err)
	}
	if v != "c" {
		t.Error("invalid value for c")
	}
	v, err = im.Get("d")
	if err == nil {
		t.Error("Invalid value for d")
	}
	a, err := im.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(a) != 3 {
		t.Error("invalid number of values")
	}
	time.Sleep(1 * time.Second)
	a, err = im.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(a) != 2 {
		t.Error("invalid number of values")
	}
	time.Sleep(1 * time.Second)
	a, err = im.GetAll()
	if err != nil {
		t.Error(err)
	}
	if len(a) != 1 {
		t.Error("invalid number of values")
	}
	v, err = im.Get("c")
	if err != nil {
		t.Error(err)
	}
	if v != "c" {
		t.Error("invalid value for c")
	}
	err = im.Close()
	if err != nil {
		t.Error(err)
	}
	_, err = im.Get("a")
	if err == nil {
		t.Error("should be closed")
	}
	err = im.Set("a", "a", time.Duration(time.Second))
	if err == nil {
		t.Error("should be closed")
	}
}
