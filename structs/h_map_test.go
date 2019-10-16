package structs

import (
	"errors"
	"fmt"
	"strconv"
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

func TestRebalanceToGreater(t *testing.T) {
	m := NewHMap()
	for i := 0; i < 400; i++ {
		m.Set(string(i), 5)
	}
	if !m.satisfiesLf() {
		t.Error("map internal array doesn't grows according when the load factor is out of bounds")
	}
}

func TestRebalanceToLower(t *testing.T) {
	m := NewHMap()
	for i := 0; i < 400; i++ {
		m.Set(strconv.Itoa(i), 5)
	}
	fmt.Println(len(m.bs))
	for i := 1; i < 400; i++ {
		err := m.Delete(strconv.Itoa(i))
		fmt.Println(err)
	}
	fmt.Println(len(m.bs))
	if !m.satisfiesLf() {
		t.Error("map internal array doesn't grows according when the load factor is out of bounds")
	}
}
