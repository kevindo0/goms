package utils

import (
    "testing"
)

func TestRandomIDGenerator(t *testing.T) {
    id := RandomIDGenerator()()
    if id != "15" {
        t.Errorf("RandomIDGenerator id: %s", id)
    }
}