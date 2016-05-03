// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/leaserunner"
	"github.com/pivotal-golang/lager"
)

type FakeRunnerDB struct {
	GetLeaseStub        func(logger lager.Logger, leaseName string, interval time.Duration) (db.Lease, bool, error)
	getLeaseMutex       sync.RWMutex
	getLeaseArgsForCall []struct {
		logger    lager.Logger
		leaseName string
		interval  time.Duration
	}
	getLeaseReturns struct {
		result1 db.Lease
		result2 bool
		result3 error
	}
}

func (fake *FakeRunnerDB) GetLease(logger lager.Logger, leaseName string, interval time.Duration) (db.Lease, bool, error) {
	fake.getLeaseMutex.Lock()
	fake.getLeaseArgsForCall = append(fake.getLeaseArgsForCall, struct {
		logger    lager.Logger
		leaseName string
		interval  time.Duration
	}{logger, leaseName, interval})
	fake.getLeaseMutex.Unlock()
	if fake.GetLeaseStub != nil {
		return fake.GetLeaseStub(logger, leaseName, interval)
	} else {
		return fake.getLeaseReturns.result1, fake.getLeaseReturns.result2, fake.getLeaseReturns.result3
	}
}

func (fake *FakeRunnerDB) GetLeaseCallCount() int {
	fake.getLeaseMutex.RLock()
	defer fake.getLeaseMutex.RUnlock()
	return len(fake.getLeaseArgsForCall)
}

func (fake *FakeRunnerDB) GetLeaseArgsForCall(i int) (lager.Logger, string, time.Duration) {
	fake.getLeaseMutex.RLock()
	defer fake.getLeaseMutex.RUnlock()
	return fake.getLeaseArgsForCall[i].logger, fake.getLeaseArgsForCall[i].leaseName, fake.getLeaseArgsForCall[i].interval
}

func (fake *FakeRunnerDB) GetLeaseReturns(result1 db.Lease, result2 bool, result3 error) {
	fake.GetLeaseStub = nil
	fake.getLeaseReturns = struct {
		result1 db.Lease
		result2 bool
		result3 error
	}{result1, result2, result3}
}

var _ leaserunner.RunnerDB = new(FakeRunnerDB)
