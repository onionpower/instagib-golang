package structs

import (
	"errors"
	"fmt"
	"testing"
)

func TestSetValueEqualsExpected(t *testing.T) {
	m := NewHMap()
	m.Set("1", 1)

	v, err := m.Get("1")

	if v != 1 {
		t.Error(fmt.Sprintf("value expected %v, actual %v", 1, v))
	}
	if err != nil {
		t.Error(fmt.Sprintf("expected err is nil, but actual %v", err))
	}
}

func TestDeleteNonexistentKeyCorrespErr(t *testing.T) {
	m := NewHMap()
	delErr1 := m.Delete("a")
	if !errors.Is(delErr1, KeyNotFoundErr) {
		t.Error(fmt.Sprintf("expected err is KeyNotFoundErr, but got %v", delErr1))
	}

	m.Set("b", 1)
	delErr2 := m.Delete("a")
	if !errors.Is(delErr1, KeyNotFoundErr) {
		t.Error(fmt.Sprintf("expected err is KeyNotFoundErr, but got %v", delErr2))
	}
}

func TestDeletedValueIsOutOfMap(t *testing.T) {
	m := NewHMap()
	m.Set("1", 1)

	delErr := m.Delete("1")
	v, getErr := m.Get("1")

	if delErr != nil {
		t.Error(fmt.Sprintf("expected err is nil, but actual %v", delErr))
	}
	if v != 0 {
		t.Error(fmt.Sprintf("expected value of nonexistent key is 0, but actual is %v", v))
	}
	if !errors.Is(getErr, KeyNotFoundErr) {
		t.Error(fmt.Sprintf("expected err is KeyNotFoundErr, but got %v", getErr))
	}
}
