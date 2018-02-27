package main

import (
	"container/list"
	"github.com/tendermint/tmlibs/common"
	"net"
	"sync"
)

// Maintains a list of SuperProxies
// Creates a sortable weight from
type SuperProxyService struct {
	common.BaseService

	active list.List
}

func (s *SuperProxyService) OnStart() error {
	return nil
}

func (s *SuperProxyService) OnStop() {

}

//
type SuperProxy struct {
	common.BaseService

	mu       sync.Mutex
	addr     net.TCPAddr
	latency  int
	sessions map[string]*PeerSession
}

func (s *SuperProxy) OnStart() error {
	return nil
}

func (s *SuperProxy) OnStop() {

}
