// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/worker"
)

type FakeInputSource struct {
	DestinationPathStub        func() string
	destinationPathMutex       sync.RWMutex
	destinationPathArgsForCall []struct {
	}
	destinationPathReturns struct {
		result1 string
	}
	destinationPathReturnsOnCall map[int]struct {
		result1 string
	}
	SourceStub        func() worker.ArtifactSource
	sourceMutex       sync.RWMutex
	sourceArgsForCall []struct {
	}
	sourceReturns struct {
		result1 worker.ArtifactSource
	}
	sourceReturnsOnCall map[int]struct {
		result1 worker.ArtifactSource
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInputSource) DestinationPath() string {
	fake.destinationPathMutex.Lock()
	ret, specificReturn := fake.destinationPathReturnsOnCall[len(fake.destinationPathArgsForCall)]
	fake.destinationPathArgsForCall = append(fake.destinationPathArgsForCall, struct {
	}{})
	stub := fake.DestinationPathStub
	fakeReturns := fake.destinationPathReturns
	fake.recordInvocation("DestinationPath", []interface{}{})
	fake.destinationPathMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeInputSource) DestinationPathCallCount() int {
	fake.destinationPathMutex.RLock()
	defer fake.destinationPathMutex.RUnlock()
	return len(fake.destinationPathArgsForCall)
}

func (fake *FakeInputSource) DestinationPathCalls(stub func() string) {
	fake.destinationPathMutex.Lock()
	defer fake.destinationPathMutex.Unlock()
	fake.DestinationPathStub = stub
}

func (fake *FakeInputSource) DestinationPathReturns(result1 string) {
	fake.destinationPathMutex.Lock()
	defer fake.destinationPathMutex.Unlock()
	fake.DestinationPathStub = nil
	fake.destinationPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInputSource) DestinationPathReturnsOnCall(i int, result1 string) {
	fake.destinationPathMutex.Lock()
	defer fake.destinationPathMutex.Unlock()
	fake.DestinationPathStub = nil
	if fake.destinationPathReturnsOnCall == nil {
		fake.destinationPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.destinationPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInputSource) Source() worker.ArtifactSource {
	fake.sourceMutex.Lock()
	ret, specificReturn := fake.sourceReturnsOnCall[len(fake.sourceArgsForCall)]
	fake.sourceArgsForCall = append(fake.sourceArgsForCall, struct {
	}{})
	stub := fake.SourceStub
	fakeReturns := fake.sourceReturns
	fake.recordInvocation("Source", []interface{}{})
	fake.sourceMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeInputSource) SourceCallCount() int {
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return len(fake.sourceArgsForCall)
}

func (fake *FakeInputSource) SourceCalls(stub func() worker.ArtifactSource) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = stub
}

func (fake *FakeInputSource) SourceReturns(result1 worker.ArtifactSource) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = nil
	fake.sourceReturns = struct {
		result1 worker.ArtifactSource
	}{result1}
}

func (fake *FakeInputSource) SourceReturnsOnCall(i int, result1 worker.ArtifactSource) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = nil
	if fake.sourceReturnsOnCall == nil {
		fake.sourceReturnsOnCall = make(map[int]struct {
			result1 worker.ArtifactSource
		})
	}
	fake.sourceReturnsOnCall[i] = struct {
		result1 worker.ArtifactSource
	}{result1}
}

func (fake *FakeInputSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.destinationPathMutex.RLock()
	defer fake.destinationPathMutex.RUnlock()
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInputSource) recordInvocation(key string, args []interface{}) {
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

var _ worker.InputSource = new(FakeInputSource)
