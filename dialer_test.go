package main

import (
	"fmt"
	"testing"
)

func TestNewSessionID(t *testing.T) {
	fmt.Println(NewPeerSessionID())
}
