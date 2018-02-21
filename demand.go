package main

import (
	"sync"

	"github.com/tendermint/tmlibs/common"
)

// Keeps track locations that will be needed.
// Tries it's best to ensure there is at least one
// IPSession that fulfills it.
type DemandService struct {
	common.BaseService

	requests sync.Map
}

//
func NewDemandService() *DemandService {
	return &DemandService{
		requests: sync.Map{},
	}
}

// Represents an anticipated future session. This can dramatically speed up
// establishing a session.
type DemandRequest struct {
	ID        string
	AccountID string
	UID       string
	Request   *SessionRequest
}

func (self *DemandRequest) GetRequest() *SessionRequest {
	return self.Request
}

func (self *DemandRequest) GetPriority() int {
	return 2
}
