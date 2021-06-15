// Code generated by counterfeiter. DO NOT EDIT.
package cloudinitfakes

import (
	"sync"

	"github.com/vmware-tanzu/cluster-api-provider-byoh/agent/cloudinit"
)

type FakeICmdRunner struct {
	RunCmdStub        func(string) error
	runCmdMutex       sync.RWMutex
	runCmdArgsForCall []struct {
		arg1 string
	}
	runCmdReturns struct {
		result1 error
	}
	runCmdReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeICmdRunner) RunCmd(arg1 string) error {
	fake.runCmdMutex.Lock()
	ret, specificReturn := fake.runCmdReturnsOnCall[len(fake.runCmdArgsForCall)]
	fake.runCmdArgsForCall = append(fake.runCmdArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RunCmdStub
	fakeReturns := fake.runCmdReturns
	fake.recordInvocation("RunCmd", []interface{}{arg1})
	fake.runCmdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeICmdRunner) RunCmdCallCount() int {
	fake.runCmdMutex.RLock()
	defer fake.runCmdMutex.RUnlock()
	return len(fake.runCmdArgsForCall)
}

func (fake *FakeICmdRunner) RunCmdCalls(stub func(string) error) {
	fake.runCmdMutex.Lock()
	defer fake.runCmdMutex.Unlock()
	fake.RunCmdStub = stub
}

func (fake *FakeICmdRunner) RunCmdArgsForCall(i int) string {
	fake.runCmdMutex.RLock()
	defer fake.runCmdMutex.RUnlock()
	argsForCall := fake.runCmdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeICmdRunner) RunCmdReturns(result1 error) {
	fake.runCmdMutex.Lock()
	defer fake.runCmdMutex.Unlock()
	fake.RunCmdStub = nil
	fake.runCmdReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeICmdRunner) RunCmdReturnsOnCall(i int, result1 error) {
	fake.runCmdMutex.Lock()
	defer fake.runCmdMutex.Unlock()
	fake.RunCmdStub = nil
	if fake.runCmdReturnsOnCall == nil {
		fake.runCmdReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runCmdReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeICmdRunner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runCmdMutex.RLock()
	defer fake.runCmdMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeICmdRunner) recordInvocation(key string, args []interface{}) {
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

var _ cloudinit.ICmdRunner = new(FakeICmdRunner)
