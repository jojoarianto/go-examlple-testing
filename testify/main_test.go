package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAdditional with tradional way
func TestAdditional(t *testing.T) {
	if Additional(3, 7) != 10 {
		t.Error("Expected 3 + 7 to equal 10")
	}
}

// TestSubtraction with tradional way
func TestSubtraction(t *testing.T) {
	if Subtraction(21, 6) != 15 {
		t.Error("Expected 21 - 6 to equal 15")
	}
}

// TestAdditional with testify
func TestAdditionalWithTestify(t *testing.T) {
	assert.Equal(t, Additional(10, 5), 15)
}

// TestAdditional with testify
func TestSubtractionWithTestify(t *testing.T) {
	assert.Equal(t, Subtraction(20, 4), 16)
}

// TestAngkaToRomawi with testify
func TestAngkaToRomawi(t *testing.T) {
	var tests = []struct {
		input    int
		expected string
	}{
		{1, "I"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{10, "X"},
		{11, "XI"},
		{12, "XII"},
		{13, "XIII"},
		{14, "XIV"},
		{15, "XV"},
		{19, "XIX"},
		{20, "XX"},
		{30, "XXX"},
		{1637, "MDCXXXVII"},
	}

	assert := assert.New(t)

	for _, test := range tests {
		assert.Equal(AngkaToRomawi(test.input), test.expected)
	}
}
