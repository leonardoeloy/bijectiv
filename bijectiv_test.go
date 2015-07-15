// Bijectiv - https://github.com/leonardoeloy/bijectiv
// Copyright 2015 Leonardo Eloy

package bijectiv

import (
	"testing"
)

func TestNew(t *testing.T) {
	b := New()

	if b == nil {
		t.Fatal("Received nil instance")
	}

	if b.Alphabet != default_alphabet {
		t.Errorf("Expected default alphabet [%s], but got [%s]", default_alphabet, b.Alphabet)
		t.Fail()
	}

	if b.Base != len(default_alphabet) {
		t.Errorf("Expected base to be [%d], but got [%d]", len(default_alphabet), b.Base)
		t.FailNow()
	}
}

func TestNewWithAlphabet(t *testing.T) {
	alpha := "ABCD"
	b := NewAlphabet(alpha)

	if b == nil {
		t.Fatal("Received nil instance")
	}

	if b.Alphabet != alpha {
		t.Errorf("Expected default alphabet [%s], but got [%s]", alpha, b.Alphabet)
		t.Fail()
	}

	if b.Base != len(alpha) {
		t.Errorf("Expected base to be [%d], but got [%d]", len(alpha), b.Base)
		t.Fail()
	}
}

func TestEncode(t *testing.T) {
	b := New()
	value := 100
	n := b.Encode(value)
	expected := "bM"

	if n != expected {
		t.Errorf("Expected encoding of [%d] to be [%s], got [%s]", value, expected, n)
		t.Fail()
	}
}

func TestDencode(t *testing.T) {
	b := New()
	value := "bM"
	n := b.Decode(value)
	expected := 100

	if n != expected {
		t.Errorf("Expected decoding of [%s] to be [%d], got [%d]", value, expected, n)
		t.Fail()
	}
}

func TestReturnFirstLetterIfZero(t *testing.T) {
	b := New()
	n := b.Encode(0)

	if n != b.Alphabet[:0] {
		t.Errorf("Expected encoding of [0] to be [%s], got [%s]",  b.Alphabet[:0], n)
		t.Fail()
	}
}