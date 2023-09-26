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
