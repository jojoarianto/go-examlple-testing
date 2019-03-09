// tradiontal testing
package main 

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

// TestAdditional
func TestAdditional(t *testing.T) {
    if Additional(3, 7) != 10 {
        t.Error("Expected 3 + 7 to equal 10")
    }
}

// TestAdditional with testify
func TestAdditionalWithTestify(t *testing.T) {
    assert.Equal(t, Additional(10,5), 15)
}
