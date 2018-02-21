package main

import (
	"testing"
	"fmt"
)

func TestNewSessionID(t *testing.T) {
	fmt.Println(NewPeerSessionID())
}
