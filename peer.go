package main

import (
	"sync"
	"time"
	"context"

	"github.com/tendermint/tmlibs/common"
	"fmt"
	"strings"
	"github.com/satori/go.uuid"
)

type PeerSessionStatus int

const (
	PeerStarted PeerSessionStatus = iota
	PeerLost
)g

// An session with a SuperProxy that maintains the same
// exit node (IP address)
type PeerSession struct {
	common.BaseService

	Proxy   *SuperProxy
	ID      string
	Auth    string
	Created time.Time
	Ended   time.Time

	// Bandwidth
	BytesUp          int64
	BytesDown        int64
	SessionBytesUp   int64
	SessionBytesDown int64

	AssignedAt time.Time
	AssignedTo *Session
	used       bool

	mu          sync.Mutex
	ctx         context.Context
	cancel      context.CancelFunc
	lastRequest time.Time
	chRequests  chan struct{}
	ticker      *time.Ticker
}

func NewPeerSessionID() string {
	uid, err := uuid.NewV4()
	if err != nil {
		return ""
	}

	return fmt.Sprintf("glob_%v", strings.Replace(uid.String(), "-", "", 4))
}

func NewPeerSession(zone string, country string) *PeerSession {
	p := &PeerSession{
		Proxy:            nil,
		ID:               NewPeerSessionID(),
		Created:          time.Now(),
		Ended:            time.Time{},
		BytesUp:          0,
		BytesDown:        0,
		SessionBytesUp:   0,
		SessionBytesDown: 0,
		AssignedAt:       time.Time{},
		AssignedTo:       nil,
		used:             false,
		mu:               sync.Mutex{},
		ctx:              nil,
		cancel:           nil,
		chRequests:       nil,
		ticker:           time.NewTicker(time.Second * 30),
	}

	p.Auth = BuildAuth(zone, country, p.ID)
	p.BaseService = *common.NewBaseService(nil, "PeerSession", p)

	return nil
}

func (p *PeerSession) OnStart() error {
	return nil
}

func BuildAuth(zone, country, sessionID string) string {
	return fmt.Sprintf(
		"lum-customer-%v-zone-%v-country-%v-session-%v",
		strings.ToLower(config.LumCustomer),
		strings.ToLower(zone),
		strings.ToLower(country),
		strings.ToLower(sessionID),
	)
}

type IPUsage struct {
	BilledUp    int64
	BilledDown  int64
	BilledTotal int64

	UnbilledUp    int64
	UnbilledDown  int64
	UnbilledTotal int64
}
