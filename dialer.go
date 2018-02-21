package main

import (
	"fmt"
	"strings"

	"github.com/satori/go.uuid"
)

// In charge of creating new IP sessions by
// dialing a SuperProxy with a new Global SessionID.
type SuperProxyDialer struct {
}


