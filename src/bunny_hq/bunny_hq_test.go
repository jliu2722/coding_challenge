package main

import "testing"

func TestBunnyHQ(t *testing.T) {
	if BunnyHQ("R5,L5,R5,R3") != 12 {
		t.Error("Expected 12 blocks for input R5,L5,R5,R3")
	}
	if BunnyHQ("R2,R2,R2") != 2 {
		t.Error("Expected 12 blocks for input R5,L5,R5,R3")
	}
	if BunnyHQ("R2,L3") != 5 {
		t.Error("Expected 12 blocks for input R5,L5,R5,R3")
	}
}