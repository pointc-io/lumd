package main

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/tendermint/tmlibs/common"
)

// Requests
type SessionRequestor interface {
	GetRequest() *SessionRequest

	GetPriority() int
}

type SessionService struct {
	common.BaseService

	sessions sync.Map
}

//
type Session struct {
	common.BaseService

	// UID of session
	ID string
	// Request info
	Request *SessionRequest
	// Time info
	Started time.Time
	Ended   time.Time
	// Bandwidth consumed
	BytesUp   int64
	BytesDown int64
	Cost      float64

	mu sync.Mutex
	// Slice of assigned peers
	peers []*PeerSession
	// Context
	ctx    context.Context
	cancel context.CancelFunc
	// IP of connecting client
	raddr net.TCPAddr
	// Channel to handle incoming HTTP requests
	chRequests chan struct{}
}

func NewSession(request *SessionRequest) *Session {
	ctx, cancel := context.WithCancel(context.Background())
	id := NewPeerSessionID()
	session := &Session{
		ID:         id,
		Request:    request,
		Started:    time.Now(),
		Ended:      time.Time{},
		BytesUp:    0,
		BytesDown:  0,
		Cost:       0,
		mu:         sync.Mutex{},
		peers:      make([]*PeerSession, 2),
		ctx:        ctx,
		cancel:     cancel,
		raddr:      net.TCPAddr{},
		chRequests: make(chan struct{}, 16),
	}

	session.BaseService = *common.NewBaseService(nil, "Session", session)

	return session
}

func (self *Session) OnStart() error {
	return nil
}

func (self *Session) OnStop() {
}

func (self *Session) handleRequest() {
}

func (self *Session) GetRequest() *SessionRequest {
	return self.Request
}

func (self *Session) GetPriority() int {
	return 1
}

type SessionRequest struct {
	// Account
	AccountID string `json:"accountId"`
	// Geo constraints
	Geo *Geo `json:"geo"`
	// Maximum life span of session
	MaxAge time.Duration `json:"maxAge"`
	// Max cost in US dollars of session
	MaxCost float32 `json:"maxCost"`
	// Bandwidth limits
	BytesUpLimit   int64 `json:"bytesUpLimit"`
	BytesDownLimit int64 `json:"bytesDownLimit"`
}

type Geo struct {
	Zip       string  `json:"zip"`
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
	State     string  `json:"state"`
	Radius    float32 `json:"radius"`
	Country   string  `json:"country"`
}
