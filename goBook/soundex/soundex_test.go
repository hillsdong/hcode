package main

import (
	"testing"
)

func TestSoundex(t *testing.T) {
	t.Error(soundex("adobe"))
}
