// Code generated by counterfeiter. DO NOT EDIT.
package storagefakes

import (
	"sync"

	"github.com/Peripli/service-manager/storage"
)

type FakeDeleteInterceptor struct {
	OnAPIDeleteStub        func(h storage.InterceptDeleteOnAPI) storage.InterceptDeleteOnAPI
	onAPIDeleteMutex       sync.RWMutex
	onAPIDeleteArgsForCall []struct {
		h storage.InterceptDeleteOnAPI
	}
	onAPIDeleteReturns struct {
		result1 storage.InterceptDeleteOnAPI
	}
	onAPIDeleteReturnsOnCall map[int]struct {
		result1 storage.InterceptDeleteOnAPI
	}
	OnTxDeleteStub        func(f storage.InterceptDeleteOnTx) storage.InterceptDeleteOnTx
	onTxDeleteMutex       sync.RWMutex
	onTxDeleteArgsForCall []struct {
		f storage.InterceptDeleteOnTx
	}
	onTxDeleteReturns struct {
		result1 storage.InterceptDeleteOnTx
	}
	onTxDeleteReturnsOnCall map[int]struct {
		result1 storage.InterceptDeleteOnTx
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDeleteInterceptor) OnAPIDelete(h storage.InterceptDeleteOnAPI) storage.InterceptDeleteOnAPI {
	fake.onAPIDeleteMutex.Lock()
	ret, specificReturn := fake.onAPIDeleteReturnsOnCall[len(fake.onAPIDeleteArgsForCall)]
	fake.onAPIDeleteArgsForCall = append(fake.onAPIDeleteArgsForCall, struct {
		h storage.InterceptDeleteOnAPI
	}{h})
	fake.recordInvocation("OnAPIDelete", []interface{}{h})
	fake.onAPIDeleteMutex.Unlock()
	if fake.OnAPIDeleteStub != nil {
		return fake.OnAPIDeleteStub(h)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.onAPIDeleteReturns.result1
}

func (fake *FakeDeleteInterceptor) OnAPIDeleteCallCount() int {
	fake.onAPIDeleteMutex.RLock()
	defer fake.onAPIDeleteMutex.RUnlock()
	return len(fake.onAPIDeleteArgsForCall)
}

func (fake *FakeDeleteInterceptor) OnAPIDeleteArgsForCall(i int) storage.InterceptDeleteOnAPI {
	fake.onAPIDeleteMutex.RLock()
	defer fake.onAPIDeleteMutex.RUnlock()
	return fake.onAPIDeleteArgsForCall[i].h
}

func (fake *FakeDeleteInterceptor) OnAPIDeleteReturns(result1 storage.InterceptDeleteOnAPI) {
	fake.OnAPIDeleteStub = nil
	fake.onAPIDeleteReturns = struct {
		result1 storage.InterceptDeleteOnAPI
	}{result1}
}

func (fake *FakeDeleteInterceptor) OnAPIDeleteReturnsOnCall(i int, result1 storage.InterceptDeleteOnAPI) {
	fake.OnAPIDeleteStub = nil
	if fake.onAPIDeleteReturnsOnCall == nil {
		fake.onAPIDeleteReturnsOnCall = make(map[int]struct {
			result1 storage.InterceptDeleteOnAPI
		})
	}
	fake.onAPIDeleteReturnsOnCall[i] = struct {
		result1 storage.InterceptDeleteOnAPI
	}{result1}
}

func (fake *FakeDeleteInterceptor) OnTxDelete(f storage.InterceptDeleteOnTx) storage.InterceptDeleteOnTx {
	fake.onTxDeleteMutex.Lock()
	ret, specificReturn := fake.onTxDeleteReturnsOnCall[len(fake.onTxDeleteArgsForCall)]
	fake.onTxDeleteArgsForCall = append(fake.onTxDeleteArgsForCall, struct {
		f storage.InterceptDeleteOnTx
	}{f})
	fake.recordInvocation("OnTxDelete", []interface{}{f})
	fake.onTxDeleteMutex.Unlock()
	if fake.OnTxDeleteStub != nil {
		return fake.OnTxDeleteStub(f)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.onTxDeleteReturns.result1
}

func (fake *FakeDeleteInterceptor) OnTxDeleteCallCount() int {
	fake.onTxDeleteMutex.RLock()
	defer fake.onTxDeleteMutex.RUnlock()
	return len(fake.onTxDeleteArgsForCall)
}

func (fake *FakeDeleteInterceptor) OnTxDeleteArgsForCall(i int) storage.InterceptDeleteOnTx {
	fake.onTxDeleteMutex.RLock()
	defer fake.onTxDeleteMutex.RUnlock()
	return fake.onTxDeleteArgsForCall[i].f
}

func (fake *FakeDeleteInterceptor) OnTxDeleteReturns(result1 storage.InterceptDeleteOnTx) {
	fake.OnTxDeleteStub = nil
	fake.onTxDeleteReturns = struct {
		result1 storage.InterceptDeleteOnTx
	}{result1}
}

func (fake *FakeDeleteInterceptor) OnTxDeleteReturnsOnCall(i int, result1 storage.InterceptDeleteOnTx) {
	fake.OnTxDeleteStub = nil
	if fake.onTxDeleteReturnsOnCall == nil {
		fake.onTxDeleteReturnsOnCall = make(map[int]struct {
			result1 storage.InterceptDeleteOnTx
		})
	}
	fake.onTxDeleteReturnsOnCall[i] = struct {
		result1 storage.InterceptDeleteOnTx
	}{result1}
}

func (fake *FakeDeleteInterceptor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.onAPIDeleteMutex.RLock()
	defer fake.onAPIDeleteMutex.RUnlock()
	fake.onTxDeleteMutex.RLock()
	defer fake.onTxDeleteMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDeleteInterceptor) recordInvocation(key string, args []interface{}) {
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

var _ storage.DeleteInterceptor = new(FakeDeleteInterceptor)