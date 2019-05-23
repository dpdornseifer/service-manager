// Code generated by counterfeiter. DO NOT EDIT.
package storagefakes

import (
	"context"
	"sync"

	"github.com/Peripli/service-manager/pkg/query"
	"github.com/Peripli/service-manager/pkg/types"
	"github.com/Peripli/service-manager/storage"
)

type FakeSecuredTransactionalRepository struct {
	CreateStub        func(context.Context, types.Object) (types.Object, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 types.Object
	}
	createReturns struct {
		result1 types.Object
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 types.Object
		result2 error
	}
	DeleteStub        func(context.Context, types.ObjectType, ...query.Criterion) (types.ObjectList, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 types.ObjectType
		arg3 []query.Criterion
	}
	deleteReturns struct {
		result1 types.ObjectList
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 types.ObjectList
		result2 error
	}
	GetStub        func(context.Context, types.ObjectType, string) (types.Object, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 types.ObjectType
		arg3 string
	}
	getReturns struct {
		result1 types.Object
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 types.Object
		result2 error
	}
	GetEncryptionKeyStub        func(context.Context, func(context.Context, []byte, []byte) ([]byte, error)) ([]byte, error)
	getEncryptionKeyMutex       sync.RWMutex
	getEncryptionKeyArgsForCall []struct {
		arg1 context.Context
		arg2 func(context.Context, []byte, []byte) ([]byte, error)
	}
	getEncryptionKeyReturns struct {
		result1 []byte
		result2 error
	}
	getEncryptionKeyReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	InTransactionStub        func(context.Context, func(ctx context.Context, storage storage.Repository) error) error
	inTransactionMutex       sync.RWMutex
	inTransactionArgsForCall []struct {
		arg1 context.Context
		arg2 func(ctx context.Context, storage storage.Repository) error
	}
	inTransactionReturns struct {
		result1 error
	}
	inTransactionReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func(context.Context, types.ObjectType, ...query.Criterion) (types.ObjectList, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
		arg2 types.ObjectType
		arg3 []query.Criterion
	}
	listReturns struct {
		result1 types.ObjectList
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 types.ObjectList
		result2 error
	}
	LockStub        func(context.Context) error
	lockMutex       sync.RWMutex
	lockArgsForCall []struct {
		arg1 context.Context
	}
	lockReturns struct {
		result1 error
	}
	lockReturnsOnCall map[int]struct {
		result1 error
	}
	SetEncryptionKeyStub        func(context.Context, []byte, func(context.Context, []byte, []byte) ([]byte, error)) error
	setEncryptionKeyMutex       sync.RWMutex
	setEncryptionKeyArgsForCall []struct {
		arg1 context.Context
		arg2 []byte
		arg3 func(context.Context, []byte, []byte) ([]byte, error)
	}
	setEncryptionKeyReturns struct {
		result1 error
	}
	setEncryptionKeyReturnsOnCall map[int]struct {
		result1 error
	}
	UnlockStub        func(context.Context) error
	unlockMutex       sync.RWMutex
	unlockArgsForCall []struct {
		arg1 context.Context
	}
	unlockReturns struct {
		result1 error
	}
	unlockReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStub        func(context.Context, types.Object, ...*query.LabelChange) (types.Object, error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 context.Context
		arg2 types.Object
		arg3 []*query.LabelChange
	}
	updateReturns struct {
		result1 types.Object
		result2 error
	}
	updateReturnsOnCall map[int]struct {
		result1 types.Object
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecuredTransactionalRepository) Create(arg1 context.Context, arg2 types.Object) (types.Object, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 types.Object
	}{arg1, arg2})
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecuredTransactionalRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeSecuredTransactionalRepository) CreateCalls(stub func(context.Context, types.Object) (types.Object, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeSecuredTransactionalRepository) CreateArgsForCall(i int) (context.Context, types.Object) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSecuredTransactionalRepository) CreateReturns(result1 types.Object, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 types.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeSecuredTransactionalRepository) CreateReturnsOnCall(i int, result1 types.Object, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 types.Object
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 types.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeSecuredTransactionalRepository) Delete(arg1 context.Context, arg2 types.ObjectType, arg3 ...query.Criterion) (types.ObjectList, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 types.ObjectType
		arg3 []query.Criterion
	}{arg1, arg2, arg3})
	fake.recordInvocation("Delete", []interface{}{arg1, arg2, arg3})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecuredTransactionalRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeSecuredTransactionalRepository) DeleteCalls(stub func(context.Context, types.ObjectType, ...query.Criterion) (types.ObjectList, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeSecuredTransactionalRepository) DeleteArgsForCall(i int) (context.Context, types.ObjectType, []query.Criterion) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecuredTransactionalRepository) DeleteReturns(result1 types.ObjectList, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 types.ObjectList
		result2 error
	}{result1, result2}
}

func (fake *FakeSecuredTransactionalRepository) DeleteReturnsOnCall(i int, result1 types.ObjectList, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 types.ObjectList
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 types.ObjectList
		result2 error
	}{result1, result2}
}

func (fake *FakeSecuredTransactionalRepository) Get(arg1 context.Context, arg2 types.ObjectType, arg3 string) (types.Object, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 types.ObjectType
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecuredTransactionalRepository) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeSecuredTransactionalRepository) GetCalls(stub func(context.Context, types.ObjectType, string) (types.Object, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeSecuredTransactionalRepository) GetArgsForCall(i int) (context.Context, types.ObjectType, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecuredTransactionalRepository) GetReturns(result1 types.Object, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 types.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeSecuredTransactionalRepository) GetReturnsOnCall(i int, result1 types.Object, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 types.Object
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 types.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeSecuredTransactionalRepository) GetEncryptionKey(arg1 context.Context, arg2 func(context.Context, []byte, []byte) ([]byte, error)) ([]byte, error) {
	fake.getEncryptionKeyMutex.Lock()
	ret, specificReturn := fake.getEncryptionKeyReturnsOnCall[len(fake.getEncryptionKeyArgsForCall)]
	fake.getEncryptionKeyArgsForCall = append(fake.getEncryptionKeyArgsForCall, struct {
		arg1 context.Context
		arg2 func(context.Context, []byte, []byte) ([]byte, error)
	}{arg1, arg2})
	fake.recordInvocation("GetEncryptionKey", []interface{}{arg1, arg2})
	fake.getEncryptionKeyMutex.Unlock()
	if fake.GetEncryptionKeyStub != nil {
		return fake.GetEncryptionKeyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getEncryptionKeyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecuredTransactionalRepository) GetEncryptionKeyCallCount() int {
	fake.getEncryptionKeyMutex.RLock()
	defer fake.getEncryptionKeyMutex.RUnlock()
	return len(fake.getEncryptionKeyArgsForCall)
}

func (fake *FakeSecuredTransactionalRepository) GetEncryptionKeyCalls(stub func(context.Context, func(context.Context, []byte, []byte) ([]byte, error)) ([]byte, error)) {
	fake.getEncryptionKeyMutex.Lock()
	defer fake.getEncryptionKeyMutex.Unlock()
	fake.GetEncryptionKeyStub = stub
}

func (fake *FakeSecuredTransactionalRepository) GetEncryptionKeyArgsForCall(i int) (context.Context, func(context.Context, []byte, []byte) ([]byte, error)) {
	fake.getEncryptionKeyMutex.RLock()
	defer fake.getEncryptionKeyMutex.RUnlock()
	argsForCall := fake.getEncryptionKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSecuredTransactionalRepository) GetEncryptionKeyReturns(result1 []byte, result2 error) {
	fake.getEncryptionKeyMutex.Lock()
	defer fake.getEncryptionKeyMutex.Unlock()
	fake.GetEncryptionKeyStub = nil
	fake.getEncryptionKeyReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeSecuredTransactionalRepository) GetEncryptionKeyReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.getEncryptionKeyMutex.Lock()
	defer fake.getEncryptionKeyMutex.Unlock()
	fake.GetEncryptionKeyStub = nil
	if fake.getEncryptionKeyReturnsOnCall == nil {
		fake.getEncryptionKeyReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.getEncryptionKeyReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeSecuredTransactionalRepository) InTransaction(arg1 context.Context, arg2 func(ctx context.Context, storage storage.Repository) error) error {
	fake.inTransactionMutex.Lock()
	ret, specificReturn := fake.inTransactionReturnsOnCall[len(fake.inTransactionArgsForCall)]
	fake.inTransactionArgsForCall = append(fake.inTransactionArgsForCall, struct {
		arg1 context.Context
		arg2 func(ctx context.Context, storage storage.Repository) error
	}{arg1, arg2})
	fake.recordInvocation("InTransaction", []interface{}{arg1, arg2})
	fake.inTransactionMutex.Unlock()
	if fake.InTransactionStub != nil {
		return fake.InTransactionStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.inTransactionReturns
	return fakeReturns.result1
}

func (fake *FakeSecuredTransactionalRepository) InTransactionCallCount() int {
	fake.inTransactionMutex.RLock()
	defer fake.inTransactionMutex.RUnlock()
	return len(fake.inTransactionArgsForCall)
}

func (fake *FakeSecuredTransactionalRepository) InTransactionCalls(stub func(context.Context, func(ctx context.Context, storage storage.Repository) error) error) {
	fake.inTransactionMutex.Lock()
	defer fake.inTransactionMutex.Unlock()
	fake.InTransactionStub = stub
}

func (fake *FakeSecuredTransactionalRepository) InTransactionArgsForCall(i int) (context.Context, func(ctx context.Context, storage storage.Repository) error) {
	fake.inTransactionMutex.RLock()
	defer fake.inTransactionMutex.RUnlock()
	argsForCall := fake.inTransactionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSecuredTransactionalRepository) InTransactionReturns(result1 error) {
	fake.inTransactionMutex.Lock()
	defer fake.inTransactionMutex.Unlock()
	fake.InTransactionStub = nil
	fake.inTransactionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecuredTransactionalRepository) InTransactionReturnsOnCall(i int, result1 error) {
	fake.inTransactionMutex.Lock()
	defer fake.inTransactionMutex.Unlock()
	fake.InTransactionStub = nil
	if fake.inTransactionReturnsOnCall == nil {
		fake.inTransactionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.inTransactionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecuredTransactionalRepository) List(arg1 context.Context, arg2 types.ObjectType, arg3 ...query.Criterion) (types.ObjectList, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
		arg2 types.ObjectType
		arg3 []query.Criterion
	}{arg1, arg2, arg3})
	fake.recordInvocation("List", []interface{}{arg1, arg2, arg3})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecuredTransactionalRepository) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeSecuredTransactionalRepository) ListCalls(stub func(context.Context, types.ObjectType, ...query.Criterion) (types.ObjectList, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeSecuredTransactionalRepository) ListArgsForCall(i int) (context.Context, types.ObjectType, []query.Criterion) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecuredTransactionalRepository) ListReturns(result1 types.ObjectList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 types.ObjectList
		result2 error
	}{result1, result2}
}

func (fake *FakeSecuredTransactionalRepository) ListReturnsOnCall(i int, result1 types.ObjectList, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 types.ObjectList
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 types.ObjectList
		result2 error
	}{result1, result2}
}

func (fake *FakeSecuredTransactionalRepository) Lock(arg1 context.Context) error {
	fake.lockMutex.Lock()
	ret, specificReturn := fake.lockReturnsOnCall[len(fake.lockArgsForCall)]
	fake.lockArgsForCall = append(fake.lockArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Lock", []interface{}{arg1})
	fake.lockMutex.Unlock()
	if fake.LockStub != nil {
		return fake.LockStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.lockReturns
	return fakeReturns.result1
}

func (fake *FakeSecuredTransactionalRepository) LockCallCount() int {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return len(fake.lockArgsForCall)
}

func (fake *FakeSecuredTransactionalRepository) LockCalls(stub func(context.Context) error) {
	fake.lockMutex.Lock()
	defer fake.lockMutex.Unlock()
	fake.LockStub = stub
}

func (fake *FakeSecuredTransactionalRepository) LockArgsForCall(i int) context.Context {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	argsForCall := fake.lockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSecuredTransactionalRepository) LockReturns(result1 error) {
	fake.lockMutex.Lock()
	defer fake.lockMutex.Unlock()
	fake.LockStub = nil
	fake.lockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecuredTransactionalRepository) LockReturnsOnCall(i int, result1 error) {
	fake.lockMutex.Lock()
	defer fake.lockMutex.Unlock()
	fake.LockStub = nil
	if fake.lockReturnsOnCall == nil {
		fake.lockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.lockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecuredTransactionalRepository) SetEncryptionKey(arg1 context.Context, arg2 []byte, arg3 func(context.Context, []byte, []byte) ([]byte, error)) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.setEncryptionKeyMutex.Lock()
	ret, specificReturn := fake.setEncryptionKeyReturnsOnCall[len(fake.setEncryptionKeyArgsForCall)]
	fake.setEncryptionKeyArgsForCall = append(fake.setEncryptionKeyArgsForCall, struct {
		arg1 context.Context
		arg2 []byte
		arg3 func(context.Context, []byte, []byte) ([]byte, error)
	}{arg1, arg2Copy, arg3})
	fake.recordInvocation("SetEncryptionKey", []interface{}{arg1, arg2Copy, arg3})
	fake.setEncryptionKeyMutex.Unlock()
	if fake.SetEncryptionKeyStub != nil {
		return fake.SetEncryptionKeyStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.setEncryptionKeyReturns
	return fakeReturns.result1
}

func (fake *FakeSecuredTransactionalRepository) SetEncryptionKeyCallCount() int {
	fake.setEncryptionKeyMutex.RLock()
	defer fake.setEncryptionKeyMutex.RUnlock()
	return len(fake.setEncryptionKeyArgsForCall)
}

func (fake *FakeSecuredTransactionalRepository) SetEncryptionKeyCalls(stub func(context.Context, []byte, func(context.Context, []byte, []byte) ([]byte, error)) error) {
	fake.setEncryptionKeyMutex.Lock()
	defer fake.setEncryptionKeyMutex.Unlock()
	fake.SetEncryptionKeyStub = stub
}

func (fake *FakeSecuredTransactionalRepository) SetEncryptionKeyArgsForCall(i int) (context.Context, []byte, func(context.Context, []byte, []byte) ([]byte, error)) {
	fake.setEncryptionKeyMutex.RLock()
	defer fake.setEncryptionKeyMutex.RUnlock()
	argsForCall := fake.setEncryptionKeyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecuredTransactionalRepository) SetEncryptionKeyReturns(result1 error) {
	fake.setEncryptionKeyMutex.Lock()
	defer fake.setEncryptionKeyMutex.Unlock()
	fake.SetEncryptionKeyStub = nil
	fake.setEncryptionKeyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecuredTransactionalRepository) SetEncryptionKeyReturnsOnCall(i int, result1 error) {
	fake.setEncryptionKeyMutex.Lock()
	defer fake.setEncryptionKeyMutex.Unlock()
	fake.SetEncryptionKeyStub = nil
	if fake.setEncryptionKeyReturnsOnCall == nil {
		fake.setEncryptionKeyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setEncryptionKeyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecuredTransactionalRepository) Unlock(arg1 context.Context) error {
	fake.unlockMutex.Lock()
	ret, specificReturn := fake.unlockReturnsOnCall[len(fake.unlockArgsForCall)]
	fake.unlockArgsForCall = append(fake.unlockArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	fake.recordInvocation("Unlock", []interface{}{arg1})
	fake.unlockMutex.Unlock()
	if fake.UnlockStub != nil {
		return fake.UnlockStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.unlockReturns
	return fakeReturns.result1
}

func (fake *FakeSecuredTransactionalRepository) UnlockCallCount() int {
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	return len(fake.unlockArgsForCall)
}

func (fake *FakeSecuredTransactionalRepository) UnlockCalls(stub func(context.Context) error) {
	fake.unlockMutex.Lock()
	defer fake.unlockMutex.Unlock()
	fake.UnlockStub = stub
}

func (fake *FakeSecuredTransactionalRepository) UnlockArgsForCall(i int) context.Context {
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	argsForCall := fake.unlockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSecuredTransactionalRepository) UnlockReturns(result1 error) {
	fake.unlockMutex.Lock()
	defer fake.unlockMutex.Unlock()
	fake.UnlockStub = nil
	fake.unlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecuredTransactionalRepository) UnlockReturnsOnCall(i int, result1 error) {
	fake.unlockMutex.Lock()
	defer fake.unlockMutex.Unlock()
	fake.UnlockStub = nil
	if fake.unlockReturnsOnCall == nil {
		fake.unlockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.unlockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSecuredTransactionalRepository) Update(arg1 context.Context, arg2 types.Object, arg3 ...*query.LabelChange) (types.Object, error) {
	fake.updateMutex.Lock()
	ret, specificReturn := fake.updateReturnsOnCall[len(fake.updateArgsForCall)]
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 context.Context
		arg2 types.Object
		arg3 []*query.LabelChange
	}{arg1, arg2, arg3})
	fake.recordInvocation("Update", []interface{}{arg1, arg2, arg3})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeSecuredTransactionalRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeSecuredTransactionalRepository) UpdateCalls(stub func(context.Context, types.Object, ...*query.LabelChange) (types.Object, error)) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = stub
}

func (fake *FakeSecuredTransactionalRepository) UpdateArgsForCall(i int) (context.Context, types.Object, []*query.LabelChange) {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	argsForCall := fake.updateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeSecuredTransactionalRepository) UpdateReturns(result1 types.Object, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 types.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeSecuredTransactionalRepository) UpdateReturnsOnCall(i int, result1 types.Object, result2 error) {
	fake.updateMutex.Lock()
	defer fake.updateMutex.Unlock()
	fake.UpdateStub = nil
	if fake.updateReturnsOnCall == nil {
		fake.updateReturnsOnCall = make(map[int]struct {
			result1 types.Object
			result2 error
		})
	}
	fake.updateReturnsOnCall[i] = struct {
		result1 types.Object
		result2 error
	}{result1, result2}
}

func (fake *FakeSecuredTransactionalRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.getEncryptionKeyMutex.RLock()
	defer fake.getEncryptionKeyMutex.RUnlock()
	fake.inTransactionMutex.RLock()
	defer fake.inTransactionMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	fake.setEncryptionKeyMutex.RLock()
	defer fake.setEncryptionKeyMutex.RUnlock()
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecuredTransactionalRepository) recordInvocation(key string, args []interface{}) {
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
