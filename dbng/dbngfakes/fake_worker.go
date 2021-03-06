// This file was generated by counterfeiter
package dbngfakes

import (
	"sync"
	"time"

	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"
)

type FakeWorker struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	StateStub        func() dbng.WorkerState
	stateMutex       sync.RWMutex
	stateArgsForCall []struct{}
	stateReturns     struct {
		result1 dbng.WorkerState
	}
	GardenAddrStub        func() *string
	gardenAddrMutex       sync.RWMutex
	gardenAddrArgsForCall []struct{}
	gardenAddrReturns     struct {
		result1 *string
	}
	BaggageclaimURLStub        func() *string
	baggageclaimURLMutex       sync.RWMutex
	baggageclaimURLArgsForCall []struct{}
	baggageclaimURLReturns     struct {
		result1 *string
	}
	HTTPProxyURLStub        func() string
	hTTPProxyURLMutex       sync.RWMutex
	hTTPProxyURLArgsForCall []struct{}
	hTTPProxyURLReturns     struct {
		result1 string
	}
	HTTPSProxyURLStub        func() string
	hTTPSProxyURLMutex       sync.RWMutex
	hTTPSProxyURLArgsForCall []struct{}
	hTTPSProxyURLReturns     struct {
		result1 string
	}
	NoProxyStub        func() string
	noProxyMutex       sync.RWMutex
	noProxyArgsForCall []struct{}
	noProxyReturns     struct {
		result1 string
	}
	ActiveContainersStub        func() int
	activeContainersMutex       sync.RWMutex
	activeContainersArgsForCall []struct{}
	activeContainersReturns     struct {
		result1 int
	}
	ResourceTypesStub        func() []atc.WorkerResourceType
	resourceTypesMutex       sync.RWMutex
	resourceTypesArgsForCall []struct{}
	resourceTypesReturns     struct {
		result1 []atc.WorkerResourceType
	}
	PlatformStub        func() string
	platformMutex       sync.RWMutex
	platformArgsForCall []struct{}
	platformReturns     struct {
		result1 string
	}
	TagsStub        func() []string
	tagsMutex       sync.RWMutex
	tagsArgsForCall []struct{}
	tagsReturns     struct {
		result1 []string
	}
	TeamIDStub        func() int
	teamIDMutex       sync.RWMutex
	teamIDArgsForCall []struct{}
	teamIDReturns     struct {
		result1 int
	}
	TeamNameStub        func() string
	teamNameMutex       sync.RWMutex
	teamNameArgsForCall []struct{}
	teamNameReturns     struct {
		result1 string
	}
	ReloadStub        func() (bool, error)
	reloadMutex       sync.RWMutex
	reloadArgsForCall []struct{}
	reloadReturns     struct {
		result1 bool
		result2 error
	}
	StartTimeStub        func() int64
	startTimeMutex       sync.RWMutex
	startTimeArgsForCall []struct{}
	startTimeReturns     struct {
		result1 int64
	}
	ExpiresAtStub        func() time.Time
	expiresAtMutex       sync.RWMutex
	expiresAtArgsForCall []struct{}
	expiresAtReturns     struct {
		result1 time.Time
	}
	LandStub        func() error
	landMutex       sync.RWMutex
	landArgsForCall []struct{}
	landReturns     struct {
		result1 error
	}
	RetireStub        func() error
	retireMutex       sync.RWMutex
	retireArgsForCall []struct{}
	retireReturns     struct {
		result1 error
	}
	PruneStub        func() error
	pruneMutex       sync.RWMutex
	pruneArgsForCall []struct{}
	pruneReturns     struct {
		result1 error
	}
	DeleteStub        func() error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct{}
	deleteReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorker) Name() string {
	fake.nameMutex.Lock()
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	return fake.nameReturns.result1
}

func (fake *FakeWorker) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeWorker) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) State() dbng.WorkerState {
	fake.stateMutex.Lock()
	fake.stateArgsForCall = append(fake.stateArgsForCall, struct{}{})
	fake.recordInvocation("State", []interface{}{})
	fake.stateMutex.Unlock()
	if fake.StateStub != nil {
		return fake.StateStub()
	}
	return fake.stateReturns.result1
}

func (fake *FakeWorker) StateCallCount() int {
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	return len(fake.stateArgsForCall)
}

func (fake *FakeWorker) StateReturns(result1 dbng.WorkerState) {
	fake.StateStub = nil
	fake.stateReturns = struct {
		result1 dbng.WorkerState
	}{result1}
}

func (fake *FakeWorker) GardenAddr() *string {
	fake.gardenAddrMutex.Lock()
	fake.gardenAddrArgsForCall = append(fake.gardenAddrArgsForCall, struct{}{})
	fake.recordInvocation("GardenAddr", []interface{}{})
	fake.gardenAddrMutex.Unlock()
	if fake.GardenAddrStub != nil {
		return fake.GardenAddrStub()
	}
	return fake.gardenAddrReturns.result1
}

func (fake *FakeWorker) GardenAddrCallCount() int {
	fake.gardenAddrMutex.RLock()
	defer fake.gardenAddrMutex.RUnlock()
	return len(fake.gardenAddrArgsForCall)
}

func (fake *FakeWorker) GardenAddrReturns(result1 *string) {
	fake.GardenAddrStub = nil
	fake.gardenAddrReturns = struct {
		result1 *string
	}{result1}
}

func (fake *FakeWorker) BaggageclaimURL() *string {
	fake.baggageclaimURLMutex.Lock()
	fake.baggageclaimURLArgsForCall = append(fake.baggageclaimURLArgsForCall, struct{}{})
	fake.recordInvocation("BaggageclaimURL", []interface{}{})
	fake.baggageclaimURLMutex.Unlock()
	if fake.BaggageclaimURLStub != nil {
		return fake.BaggageclaimURLStub()
	}
	return fake.baggageclaimURLReturns.result1
}

func (fake *FakeWorker) BaggageclaimURLCallCount() int {
	fake.baggageclaimURLMutex.RLock()
	defer fake.baggageclaimURLMutex.RUnlock()
	return len(fake.baggageclaimURLArgsForCall)
}

func (fake *FakeWorker) BaggageclaimURLReturns(result1 *string) {
	fake.BaggageclaimURLStub = nil
	fake.baggageclaimURLReturns = struct {
		result1 *string
	}{result1}
}

func (fake *FakeWorker) HTTPProxyURL() string {
	fake.hTTPProxyURLMutex.Lock()
	fake.hTTPProxyURLArgsForCall = append(fake.hTTPProxyURLArgsForCall, struct{}{})
	fake.recordInvocation("HTTPProxyURL", []interface{}{})
	fake.hTTPProxyURLMutex.Unlock()
	if fake.HTTPProxyURLStub != nil {
		return fake.HTTPProxyURLStub()
	}
	return fake.hTTPProxyURLReturns.result1
}

func (fake *FakeWorker) HTTPProxyURLCallCount() int {
	fake.hTTPProxyURLMutex.RLock()
	defer fake.hTTPProxyURLMutex.RUnlock()
	return len(fake.hTTPProxyURLArgsForCall)
}

func (fake *FakeWorker) HTTPProxyURLReturns(result1 string) {
	fake.HTTPProxyURLStub = nil
	fake.hTTPProxyURLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) HTTPSProxyURL() string {
	fake.hTTPSProxyURLMutex.Lock()
	fake.hTTPSProxyURLArgsForCall = append(fake.hTTPSProxyURLArgsForCall, struct{}{})
	fake.recordInvocation("HTTPSProxyURL", []interface{}{})
	fake.hTTPSProxyURLMutex.Unlock()
	if fake.HTTPSProxyURLStub != nil {
		return fake.HTTPSProxyURLStub()
	}
	return fake.hTTPSProxyURLReturns.result1
}

func (fake *FakeWorker) HTTPSProxyURLCallCount() int {
	fake.hTTPSProxyURLMutex.RLock()
	defer fake.hTTPSProxyURLMutex.RUnlock()
	return len(fake.hTTPSProxyURLArgsForCall)
}

func (fake *FakeWorker) HTTPSProxyURLReturns(result1 string) {
	fake.HTTPSProxyURLStub = nil
	fake.hTTPSProxyURLReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) NoProxy() string {
	fake.noProxyMutex.Lock()
	fake.noProxyArgsForCall = append(fake.noProxyArgsForCall, struct{}{})
	fake.recordInvocation("NoProxy", []interface{}{})
	fake.noProxyMutex.Unlock()
	if fake.NoProxyStub != nil {
		return fake.NoProxyStub()
	}
	return fake.noProxyReturns.result1
}

func (fake *FakeWorker) NoProxyCallCount() int {
	fake.noProxyMutex.RLock()
	defer fake.noProxyMutex.RUnlock()
	return len(fake.noProxyArgsForCall)
}

func (fake *FakeWorker) NoProxyReturns(result1 string) {
	fake.NoProxyStub = nil
	fake.noProxyReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) ActiveContainers() int {
	fake.activeContainersMutex.Lock()
	fake.activeContainersArgsForCall = append(fake.activeContainersArgsForCall, struct{}{})
	fake.recordInvocation("ActiveContainers", []interface{}{})
	fake.activeContainersMutex.Unlock()
	if fake.ActiveContainersStub != nil {
		return fake.ActiveContainersStub()
	}
	return fake.activeContainersReturns.result1
}

func (fake *FakeWorker) ActiveContainersCallCount() int {
	fake.activeContainersMutex.RLock()
	defer fake.activeContainersMutex.RUnlock()
	return len(fake.activeContainersArgsForCall)
}

func (fake *FakeWorker) ActiveContainersReturns(result1 int) {
	fake.ActiveContainersStub = nil
	fake.activeContainersReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorker) ResourceTypes() []atc.WorkerResourceType {
	fake.resourceTypesMutex.Lock()
	fake.resourceTypesArgsForCall = append(fake.resourceTypesArgsForCall, struct{}{})
	fake.recordInvocation("ResourceTypes", []interface{}{})
	fake.resourceTypesMutex.Unlock()
	if fake.ResourceTypesStub != nil {
		return fake.ResourceTypesStub()
	}
	return fake.resourceTypesReturns.result1
}

func (fake *FakeWorker) ResourceTypesCallCount() int {
	fake.resourceTypesMutex.RLock()
	defer fake.resourceTypesMutex.RUnlock()
	return len(fake.resourceTypesArgsForCall)
}

func (fake *FakeWorker) ResourceTypesReturns(result1 []atc.WorkerResourceType) {
	fake.ResourceTypesStub = nil
	fake.resourceTypesReturns = struct {
		result1 []atc.WorkerResourceType
	}{result1}
}

func (fake *FakeWorker) Platform() string {
	fake.platformMutex.Lock()
	fake.platformArgsForCall = append(fake.platformArgsForCall, struct{}{})
	fake.recordInvocation("Platform", []interface{}{})
	fake.platformMutex.Unlock()
	if fake.PlatformStub != nil {
		return fake.PlatformStub()
	}
	return fake.platformReturns.result1
}

func (fake *FakeWorker) PlatformCallCount() int {
	fake.platformMutex.RLock()
	defer fake.platformMutex.RUnlock()
	return len(fake.platformArgsForCall)
}

func (fake *FakeWorker) PlatformReturns(result1 string) {
	fake.PlatformStub = nil
	fake.platformReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) Tags() []string {
	fake.tagsMutex.Lock()
	fake.tagsArgsForCall = append(fake.tagsArgsForCall, struct{}{})
	fake.recordInvocation("Tags", []interface{}{})
	fake.tagsMutex.Unlock()
	if fake.TagsStub != nil {
		return fake.TagsStub()
	}
	return fake.tagsReturns.result1
}

func (fake *FakeWorker) TagsCallCount() int {
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	return len(fake.tagsArgsForCall)
}

func (fake *FakeWorker) TagsReturns(result1 []string) {
	fake.TagsStub = nil
	fake.tagsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeWorker) TeamID() int {
	fake.teamIDMutex.Lock()
	fake.teamIDArgsForCall = append(fake.teamIDArgsForCall, struct{}{})
	fake.recordInvocation("TeamID", []interface{}{})
	fake.teamIDMutex.Unlock()
	if fake.TeamIDStub != nil {
		return fake.TeamIDStub()
	}
	return fake.teamIDReturns.result1
}

func (fake *FakeWorker) TeamIDCallCount() int {
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	return len(fake.teamIDArgsForCall)
}

func (fake *FakeWorker) TeamIDReturns(result1 int) {
	fake.TeamIDStub = nil
	fake.teamIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeWorker) TeamName() string {
	fake.teamNameMutex.Lock()
	fake.teamNameArgsForCall = append(fake.teamNameArgsForCall, struct{}{})
	fake.recordInvocation("TeamName", []interface{}{})
	fake.teamNameMutex.Unlock()
	if fake.TeamNameStub != nil {
		return fake.TeamNameStub()
	}
	return fake.teamNameReturns.result1
}

func (fake *FakeWorker) TeamNameCallCount() int {
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	return len(fake.teamNameArgsForCall)
}

func (fake *FakeWorker) TeamNameReturns(result1 string) {
	fake.TeamNameStub = nil
	fake.teamNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeWorker) Reload() (bool, error) {
	fake.reloadMutex.Lock()
	fake.reloadArgsForCall = append(fake.reloadArgsForCall, struct{}{})
	fake.recordInvocation("Reload", []interface{}{})
	fake.reloadMutex.Unlock()
	if fake.ReloadStub != nil {
		return fake.ReloadStub()
	}
	return fake.reloadReturns.result1, fake.reloadReturns.result2
}

func (fake *FakeWorker) ReloadCallCount() int {
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	return len(fake.reloadArgsForCall)
}

func (fake *FakeWorker) ReloadReturns(result1 bool, result2 error) {
	fake.ReloadStub = nil
	fake.reloadReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeWorker) StartTime() int64 {
	fake.startTimeMutex.Lock()
	fake.startTimeArgsForCall = append(fake.startTimeArgsForCall, struct{}{})
	fake.recordInvocation("StartTime", []interface{}{})
	fake.startTimeMutex.Unlock()
	if fake.StartTimeStub != nil {
		return fake.StartTimeStub()
	}
	return fake.startTimeReturns.result1
}

func (fake *FakeWorker) StartTimeCallCount() int {
	fake.startTimeMutex.RLock()
	defer fake.startTimeMutex.RUnlock()
	return len(fake.startTimeArgsForCall)
}

func (fake *FakeWorker) StartTimeReturns(result1 int64) {
	fake.StartTimeStub = nil
	fake.startTimeReturns = struct {
		result1 int64
	}{result1}
}

func (fake *FakeWorker) ExpiresAt() time.Time {
	fake.expiresAtMutex.Lock()
	fake.expiresAtArgsForCall = append(fake.expiresAtArgsForCall, struct{}{})
	fake.recordInvocation("ExpiresAt", []interface{}{})
	fake.expiresAtMutex.Unlock()
	if fake.ExpiresAtStub != nil {
		return fake.ExpiresAtStub()
	}
	return fake.expiresAtReturns.result1
}

func (fake *FakeWorker) ExpiresAtCallCount() int {
	fake.expiresAtMutex.RLock()
	defer fake.expiresAtMutex.RUnlock()
	return len(fake.expiresAtArgsForCall)
}

func (fake *FakeWorker) ExpiresAtReturns(result1 time.Time) {
	fake.ExpiresAtStub = nil
	fake.expiresAtReturns = struct {
		result1 time.Time
	}{result1}
}

func (fake *FakeWorker) Land() error {
	fake.landMutex.Lock()
	fake.landArgsForCall = append(fake.landArgsForCall, struct{}{})
	fake.recordInvocation("Land", []interface{}{})
	fake.landMutex.Unlock()
	if fake.LandStub != nil {
		return fake.LandStub()
	}
	return fake.landReturns.result1
}

func (fake *FakeWorker) LandCallCount() int {
	fake.landMutex.RLock()
	defer fake.landMutex.RUnlock()
	return len(fake.landArgsForCall)
}

func (fake *FakeWorker) LandReturns(result1 error) {
	fake.LandStub = nil
	fake.landReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorker) Retire() error {
	fake.retireMutex.Lock()
	fake.retireArgsForCall = append(fake.retireArgsForCall, struct{}{})
	fake.recordInvocation("Retire", []interface{}{})
	fake.retireMutex.Unlock()
	if fake.RetireStub != nil {
		return fake.RetireStub()
	}
	return fake.retireReturns.result1
}

func (fake *FakeWorker) RetireCallCount() int {
	fake.retireMutex.RLock()
	defer fake.retireMutex.RUnlock()
	return len(fake.retireArgsForCall)
}

func (fake *FakeWorker) RetireReturns(result1 error) {
	fake.RetireStub = nil
	fake.retireReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorker) Prune() error {
	fake.pruneMutex.Lock()
	fake.pruneArgsForCall = append(fake.pruneArgsForCall, struct{}{})
	fake.recordInvocation("Prune", []interface{}{})
	fake.pruneMutex.Unlock()
	if fake.PruneStub != nil {
		return fake.PruneStub()
	}
	return fake.pruneReturns.result1
}

func (fake *FakeWorker) PruneCallCount() int {
	fake.pruneMutex.RLock()
	defer fake.pruneMutex.RUnlock()
	return len(fake.pruneArgsForCall)
}

func (fake *FakeWorker) PruneReturns(result1 error) {
	fake.PruneStub = nil
	fake.pruneReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorker) Delete() error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct{}{})
	fake.recordInvocation("Delete", []interface{}{})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub()
	}
	return fake.deleteReturns.result1
}

func (fake *FakeWorker) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeWorker) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.stateMutex.RLock()
	defer fake.stateMutex.RUnlock()
	fake.gardenAddrMutex.RLock()
	defer fake.gardenAddrMutex.RUnlock()
	fake.baggageclaimURLMutex.RLock()
	defer fake.baggageclaimURLMutex.RUnlock()
	fake.hTTPProxyURLMutex.RLock()
	defer fake.hTTPProxyURLMutex.RUnlock()
	fake.hTTPSProxyURLMutex.RLock()
	defer fake.hTTPSProxyURLMutex.RUnlock()
	fake.noProxyMutex.RLock()
	defer fake.noProxyMutex.RUnlock()
	fake.activeContainersMutex.RLock()
	defer fake.activeContainersMutex.RUnlock()
	fake.resourceTypesMutex.RLock()
	defer fake.resourceTypesMutex.RUnlock()
	fake.platformMutex.RLock()
	defer fake.platformMutex.RUnlock()
	fake.tagsMutex.RLock()
	defer fake.tagsMutex.RUnlock()
	fake.teamIDMutex.RLock()
	defer fake.teamIDMutex.RUnlock()
	fake.teamNameMutex.RLock()
	defer fake.teamNameMutex.RUnlock()
	fake.reloadMutex.RLock()
	defer fake.reloadMutex.RUnlock()
	fake.startTimeMutex.RLock()
	defer fake.startTimeMutex.RUnlock()
	fake.expiresAtMutex.RLock()
	defer fake.expiresAtMutex.RUnlock()
	fake.landMutex.RLock()
	defer fake.landMutex.RUnlock()
	fake.retireMutex.RLock()
	defer fake.retireMutex.RUnlock()
	fake.pruneMutex.RLock()
	defer fake.pruneMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeWorker) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ dbng.Worker = new(FakeWorker)
