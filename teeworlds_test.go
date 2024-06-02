package main

import (
	"testing"
)

func TestIsCtrlMsg(t *testing.T) {
	data := []byte{0x04, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0x04}

	got := isCtrlMsg(data)
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
