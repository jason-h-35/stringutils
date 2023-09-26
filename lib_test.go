package stringutils

import (
	"strings"
	"testing"
)

func TestUpper(t *testing.T) {
	s := "aoeuaoeu"
	expect := strings.ToUpper(s)
	result := Upper(s)

	if expect != result {
		t.Errorf("expected %v but got %v", expect, result)
	}
}

func TestLower(t *testing.T) {
	s := "ueoaueoa"
	expect := strings.ToLower(s)
	result := Lower(s)

	if expect != result {
		t.Errorf("expected %v but got %v", expect, result)
	}
}

func BenchmarkUpper(t *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < t.N; i++ {
		_ = Upper(s)
	}
}

func BenchmarkLower(t *testing.B) {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i := 0; i < t.N; i++ {
		_ = Lower(s)
	}
}
