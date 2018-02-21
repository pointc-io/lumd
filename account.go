package main

import "sync"
import "github.com/tendermint/tmlibs/common"

type AccountService struct {
	common.BaseService
}

type Account struct {
	common.BaseService

	ID                    string
	Name                  string
	Balance               float32
	DefaultMaxSessionCost float32

	PrecisionCost Cost
	CountryCost   Cost
	DCCost        Cost

	Webhooks *Webhooks

	mu       sync.Mutex
	sessions sync.Map
}

type User struct {
	ID        string
	AccountID string
	Username  string
	Email     string
	Password  string

	DefaultMaxSessionCost float32
}

type CostType int

const (
	Monthly CostType = iota
	PerGB
)

type Cost struct {
	Type  int
	Value float32
}

type Webhooks struct {
	Demand string
}
