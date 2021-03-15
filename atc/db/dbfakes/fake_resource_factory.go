// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/db"
)

type FakeResourceFactory struct {
	AllResourcesStub        func() ([]db.Resource, error)
	allResourcesMutex       sync.RWMutex
	allResourcesArgsForCall []struct {
	}
	allResourcesReturns struct {
		result1 []db.Resource
		result2 error
	}
	allResourcesReturnsOnCall map[int]struct {
		result1 []db.Resource
		result2 error
	}
	ResourceStub        func(int) (db.Resource, bool, error)
	resourceMutex       sync.RWMutex
	resourceArgsForCall []struct {
		arg1 int
	}
	resourceReturns struct {
		result1 db.Resource
		result2 bool
		result3 error
	}
	resourceReturnsOnCall map[int]struct {
		result1 db.Resource
		result2 bool
		result3 error
	}
	VisibleResourcesStub        func([]string) ([]db.Resource, error)
	visibleResourcesMutex       sync.RWMutex
	visibleResourcesArgsForCall []struct {
		arg1 []string
	}
	visibleResourcesReturns struct {
		result1 []db.Resource
		result2 error
	}
	visibleResourcesReturnsOnCall map[int]struct {
		result1 []db.Resource
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceFactory) AllResources() ([]db.Resource, error) {
	fake.allResourcesMutex.Lock()
	ret, specificReturn := fake.allResourcesReturnsOnCall[len(fake.allResourcesArgsForCall)]
	fake.allResourcesArgsForCall = append(fake.allResourcesArgsForCall, struct {
	}{})
	stub := fake.AllResourcesStub
	fakeReturns := fake.allResourcesReturns
	fake.recordInvocation("AllResources", []interface{}{})
	fake.allResourcesMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResourceFactory) AllResourcesCallCount() int {
	fake.allResourcesMutex.RLock()
	defer fake.allResourcesMutex.RUnlock()
	return len(fake.allResourcesArgsForCall)
}

func (fake *FakeResourceFactory) AllResourcesCalls(stub func() ([]db.Resource, error)) {
	fake.allResourcesMutex.Lock()
	defer fake.allResourcesMutex.Unlock()
	fake.AllResourcesStub = stub
}

func (fake *FakeResourceFactory) AllResourcesReturns(result1 []db.Resource, result2 error) {
	fake.allResourcesMutex.Lock()
	defer fake.allResourcesMutex.Unlock()
	fake.AllResourcesStub = nil
	fake.allResourcesReturns = struct {
		result1 []db.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) AllResourcesReturnsOnCall(i int, result1 []db.Resource, result2 error) {
	fake.allResourcesMutex.Lock()
	defer fake.allResourcesMutex.Unlock()
	fake.AllResourcesStub = nil
	if fake.allResourcesReturnsOnCall == nil {
		fake.allResourcesReturnsOnCall = make(map[int]struct {
			result1 []db.Resource
			result2 error
		})
	}
	fake.allResourcesReturnsOnCall[i] = struct {
		result1 []db.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) Resource(arg1 int) (db.Resource, bool, error) {
	fake.resourceMutex.Lock()
	ret, specificReturn := fake.resourceReturnsOnCall[len(fake.resourceArgsForCall)]
	fake.resourceArgsForCall = append(fake.resourceArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.ResourceStub
	fakeReturns := fake.resourceReturns
	fake.recordInvocation("Resource", []interface{}{arg1})
	fake.resourceMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeResourceFactory) ResourceCallCount() int {
	fake.resourceMutex.RLock()
	defer fake.resourceMutex.RUnlock()
	return len(fake.resourceArgsForCall)
}

func (fake *FakeResourceFactory) ResourceCalls(stub func(int) (db.Resource, bool, error)) {
	fake.resourceMutex.Lock()
	defer fake.resourceMutex.Unlock()
	fake.ResourceStub = stub
}

func (fake *FakeResourceFactory) ResourceArgsForCall(i int) int {
	fake.resourceMutex.RLock()
	defer fake.resourceMutex.RUnlock()
	argsForCall := fake.resourceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeResourceFactory) ResourceReturns(result1 db.Resource, result2 bool, result3 error) {
	fake.resourceMutex.Lock()
	defer fake.resourceMutex.Unlock()
	fake.ResourceStub = nil
	fake.resourceReturns = struct {
		result1 db.Resource
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceFactory) ResourceReturnsOnCall(i int, result1 db.Resource, result2 bool, result3 error) {
	fake.resourceMutex.Lock()
	defer fake.resourceMutex.Unlock()
	fake.ResourceStub = nil
	if fake.resourceReturnsOnCall == nil {
		fake.resourceReturnsOnCall = make(map[int]struct {
			result1 db.Resource
			result2 bool
			result3 error
		})
	}
	fake.resourceReturnsOnCall[i] = struct {
		result1 db.Resource
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceFactory) VisibleResources(arg1 []string) ([]db.Resource, error) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.visibleResourcesMutex.Lock()
	ret, specificReturn := fake.visibleResourcesReturnsOnCall[len(fake.visibleResourcesArgsForCall)]
	fake.visibleResourcesArgsForCall = append(fake.visibleResourcesArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.VisibleResourcesStub
	fakeReturns := fake.visibleResourcesReturns
	fake.recordInvocation("VisibleResources", []interface{}{arg1Copy})
	fake.visibleResourcesMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeResourceFactory) VisibleResourcesCallCount() int {
	fake.visibleResourcesMutex.RLock()
	defer fake.visibleResourcesMutex.RUnlock()
	return len(fake.visibleResourcesArgsForCall)
}

func (fake *FakeResourceFactory) VisibleResourcesCalls(stub func([]string) ([]db.Resource, error)) {
	fake.visibleResourcesMutex.Lock()
	defer fake.visibleResourcesMutex.Unlock()
	fake.VisibleResourcesStub = stub
}

func (fake *FakeResourceFactory) VisibleResourcesArgsForCall(i int) []string {
	fake.visibleResourcesMutex.RLock()
	defer fake.visibleResourcesMutex.RUnlock()
	argsForCall := fake.visibleResourcesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeResourceFactory) VisibleResourcesReturns(result1 []db.Resource, result2 error) {
	fake.visibleResourcesMutex.Lock()
	defer fake.visibleResourcesMutex.Unlock()
	fake.VisibleResourcesStub = nil
	fake.visibleResourcesReturns = struct {
		result1 []db.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) VisibleResourcesReturnsOnCall(i int, result1 []db.Resource, result2 error) {
	fake.visibleResourcesMutex.Lock()
	defer fake.visibleResourcesMutex.Unlock()
	fake.VisibleResourcesStub = nil
	if fake.visibleResourcesReturnsOnCall == nil {
		fake.visibleResourcesReturnsOnCall = make(map[int]struct {
			result1 []db.Resource
			result2 error
		})
	}
	fake.visibleResourcesReturnsOnCall[i] = struct {
		result1 []db.Resource
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.allResourcesMutex.RLock()
	defer fake.allResourcesMutex.RUnlock()
	fake.resourceMutex.RLock()
	defer fake.resourceMutex.RUnlock()
	fake.visibleResourcesMutex.RLock()
	defer fake.visibleResourcesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceFactory) recordInvocation(key string, args []interface{}) {
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

var _ db.ResourceFactory = new(FakeResourceFactory)
