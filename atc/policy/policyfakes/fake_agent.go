// Code generated by counterfeiter. DO NOT EDIT.
package policyfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/policy"
)

type FakeAgent struct {
	CheckStub        func(policy.PolicyCheckInput) (policy.PolicyCheckOutput, error)
	checkMutex       sync.RWMutex
	checkArgsForCall []struct {
		arg1 policy.PolicyCheckInput
	}
	checkReturns struct {
		result1 policy.PolicyCheckOutput
		result2 error
	}
	checkReturnsOnCall map[int]struct {
		result1 policy.PolicyCheckOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAgent) Check(arg1 policy.PolicyCheckInput) (policy.PolicyCheckOutput, error) {
	fake.checkMutex.Lock()
	ret, specificReturn := fake.checkReturnsOnCall[len(fake.checkArgsForCall)]
	fake.checkArgsForCall = append(fake.checkArgsForCall, struct {
		arg1 policy.PolicyCheckInput
	}{arg1})
	stub := fake.CheckStub
	fakeReturns := fake.checkReturns
	fake.recordInvocation("Check", []interface{}{arg1})
	fake.checkMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAgent) CheckCallCount() int {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	return len(fake.checkArgsForCall)
}

func (fake *FakeAgent) CheckCalls(stub func(policy.PolicyCheckInput) (policy.PolicyCheckOutput, error)) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = stub
}

func (fake *FakeAgent) CheckArgsForCall(i int) policy.PolicyCheckInput {
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	argsForCall := fake.checkArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeAgent) CheckReturns(result1 policy.PolicyCheckOutput, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	fake.checkReturns = struct {
		result1 policy.PolicyCheckOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeAgent) CheckReturnsOnCall(i int, result1 policy.PolicyCheckOutput, result2 error) {
	fake.checkMutex.Lock()
	defer fake.checkMutex.Unlock()
	fake.CheckStub = nil
	if fake.checkReturnsOnCall == nil {
		fake.checkReturnsOnCall = make(map[int]struct {
			result1 policy.PolicyCheckOutput
			result2 error
		})
	}
	fake.checkReturnsOnCall[i] = struct {
		result1 policy.PolicyCheckOutput
		result2 error
	}{result1, result2}
}

func (fake *FakeAgent) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkMutex.RLock()
	defer fake.checkMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAgent) recordInvocation(key string, args []interface{}) {
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

var _ policy.Agent = new(FakeAgent)
