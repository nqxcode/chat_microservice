package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/nqxcode/chat_microservice/internal/repository.ChatRepository -o ./mocks\chat_repository_minimock.go -n ChatRepositoryMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/nqxcode/chat_microservice/internal/model"
)

// ChatRepositoryMock implements repository.ChatRepository
type ChatRepositoryMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, model *model.ChatInfo) (i1 int64, err error)
	inspectFuncCreate   func(ctx context.Context, model *model.ChatInfo)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mChatRepositoryMockCreate

	funcDelete          func(ctx context.Context, id int64) (err error)
	inspectFuncDelete   func(ctx context.Context, id int64)
	afterDeleteCounter  uint64
	beforeDeleteCounter uint64
	DeleteMock          mChatRepositoryMockDelete

	funcGet          func(ctx context.Context, id int64) (cp1 *model.Chat, err error)
	inspectFuncGet   func(ctx context.Context, id int64)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mChatRepositoryMockGet
}

// NewChatRepositoryMock returns a mock for repository.ChatRepository
func NewChatRepositoryMock(t minimock.Tester) *ChatRepositoryMock {
	m := &ChatRepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mChatRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*ChatRepositoryMockCreateParams{}

	m.DeleteMock = mChatRepositoryMockDelete{mock: m}
	m.DeleteMock.callArgs = []*ChatRepositoryMockDeleteParams{}

	m.GetMock = mChatRepositoryMockGet{mock: m}
	m.GetMock.callArgs = []*ChatRepositoryMockGetParams{}

	return m
}

type mChatRepositoryMockCreate struct {
	mock               *ChatRepositoryMock
	defaultExpectation *ChatRepositoryMockCreateExpectation
	expectations       []*ChatRepositoryMockCreateExpectation

	callArgs []*ChatRepositoryMockCreateParams
	mutex    sync.RWMutex
}

// ChatRepositoryMockCreateExpectation specifies expectation struct of the ChatRepository.Create
type ChatRepositoryMockCreateExpectation struct {
	mock    *ChatRepositoryMock
	params  *ChatRepositoryMockCreateParams
	results *ChatRepositoryMockCreateResults
	Counter uint64
}

// ChatRepositoryMockCreateParams contains parameters of the ChatRepository.Create
type ChatRepositoryMockCreateParams struct {
	ctx   context.Context
	model *model.ChatInfo
}

// ChatRepositoryMockCreateResults contains results of the ChatRepository.Create
type ChatRepositoryMockCreateResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for ChatRepository.Create
func (mmCreate *mChatRepositoryMockCreate) Expect(ctx context.Context, model *model.ChatInfo) *mChatRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatRepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &ChatRepositoryMockCreateParams{ctx, model}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the ChatRepository.Create
func (mmCreate *mChatRepositoryMockCreate) Inspect(f func(ctx context.Context, model *model.ChatInfo)) *mChatRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for ChatRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by ChatRepository.Create
func (mmCreate *mChatRepositoryMockCreate) Return(i1 int64, err error) *ChatRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &ChatRepositoryMockCreateResults{i1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the ChatRepository.Create method
func (mmCreate *mChatRepositoryMockCreate) Set(f func(ctx context.Context, model *model.ChatInfo) (i1 int64, err error)) *ChatRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the ChatRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the ChatRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the ChatRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mChatRepositoryMockCreate) When(ctx context.Context, model *model.ChatInfo) *ChatRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatRepositoryMock.Create mock is already set by Set")
	}

	expectation := &ChatRepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &ChatRepositoryMockCreateParams{ctx, model},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up ChatRepository.Create return parameters for the expectation previously defined by the When method
func (e *ChatRepositoryMockCreateExpectation) Then(i1 int64, err error) *ChatRepositoryMock {
	e.results = &ChatRepositoryMockCreateResults{i1, err}
	return e.mock
}

// Create implements repository.ChatRepository
func (mmCreate *ChatRepositoryMock) Create(ctx context.Context, model *model.ChatInfo) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, model)
	}

	mm_params := &ChatRepositoryMockCreateParams{ctx, model}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.i1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := ChatRepositoryMockCreateParams{ctx, model}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("ChatRepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the ChatRepositoryMock.Create")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, model)
	}
	mmCreate.t.Fatalf("Unexpected call to ChatRepositoryMock.Create. %v %v", ctx, model)
	return
}

// CreateAfterCounter returns a count of finished ChatRepositoryMock.Create invocations
func (mmCreate *ChatRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of ChatRepositoryMock.Create invocations
func (mmCreate *ChatRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to ChatRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mChatRepositoryMockCreate) Calls() []*ChatRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*ChatRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *ChatRepositoryMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *ChatRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatRepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatRepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to ChatRepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to ChatRepositoryMock.Create")
	}
}

type mChatRepositoryMockDelete struct {
	mock               *ChatRepositoryMock
	defaultExpectation *ChatRepositoryMockDeleteExpectation
	expectations       []*ChatRepositoryMockDeleteExpectation

	callArgs []*ChatRepositoryMockDeleteParams
	mutex    sync.RWMutex
}

// ChatRepositoryMockDeleteExpectation specifies expectation struct of the ChatRepository.Delete
type ChatRepositoryMockDeleteExpectation struct {
	mock    *ChatRepositoryMock
	params  *ChatRepositoryMockDeleteParams
	results *ChatRepositoryMockDeleteResults
	Counter uint64
}

// ChatRepositoryMockDeleteParams contains parameters of the ChatRepository.Delete
type ChatRepositoryMockDeleteParams struct {
	ctx context.Context
	id  int64
}

// ChatRepositoryMockDeleteResults contains results of the ChatRepository.Delete
type ChatRepositoryMockDeleteResults struct {
	err error
}

// Expect sets up expected params for ChatRepository.Delete
func (mmDelete *mChatRepositoryMockDelete) Expect(ctx context.Context, id int64) *mChatRepositoryMockDelete {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ChatRepositoryMockDeleteExpectation{}
	}

	mmDelete.defaultExpectation.params = &ChatRepositoryMockDeleteParams{ctx, id}
	for _, e := range mmDelete.expectations {
		if minimock.Equal(e.params, mmDelete.defaultExpectation.params) {
			mmDelete.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDelete.defaultExpectation.params)
		}
	}

	return mmDelete
}

// Inspect accepts an inspector function that has same arguments as the ChatRepository.Delete
func (mmDelete *mChatRepositoryMockDelete) Inspect(f func(ctx context.Context, id int64)) *mChatRepositoryMockDelete {
	if mmDelete.mock.inspectFuncDelete != nil {
		mmDelete.mock.t.Fatalf("Inspect function is already set for ChatRepositoryMock.Delete")
	}

	mmDelete.mock.inspectFuncDelete = f

	return mmDelete
}

// Return sets up results that will be returned by ChatRepository.Delete
func (mmDelete *mChatRepositoryMockDelete) Return(err error) *ChatRepositoryMock {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatRepositoryMock.Delete mock is already set by Set")
	}

	if mmDelete.defaultExpectation == nil {
		mmDelete.defaultExpectation = &ChatRepositoryMockDeleteExpectation{mock: mmDelete.mock}
	}
	mmDelete.defaultExpectation.results = &ChatRepositoryMockDeleteResults{err}
	return mmDelete.mock
}

// Set uses given function f to mock the ChatRepository.Delete method
func (mmDelete *mChatRepositoryMockDelete) Set(f func(ctx context.Context, id int64) (err error)) *ChatRepositoryMock {
	if mmDelete.defaultExpectation != nil {
		mmDelete.mock.t.Fatalf("Default expectation is already set for the ChatRepository.Delete method")
	}

	if len(mmDelete.expectations) > 0 {
		mmDelete.mock.t.Fatalf("Some expectations are already set for the ChatRepository.Delete method")
	}

	mmDelete.mock.funcDelete = f
	return mmDelete.mock
}

// When sets expectation for the ChatRepository.Delete which will trigger the result defined by the following
// Then helper
func (mmDelete *mChatRepositoryMockDelete) When(ctx context.Context, id int64) *ChatRepositoryMockDeleteExpectation {
	if mmDelete.mock.funcDelete != nil {
		mmDelete.mock.t.Fatalf("ChatRepositoryMock.Delete mock is already set by Set")
	}

	expectation := &ChatRepositoryMockDeleteExpectation{
		mock:   mmDelete.mock,
		params: &ChatRepositoryMockDeleteParams{ctx, id},
	}
	mmDelete.expectations = append(mmDelete.expectations, expectation)
	return expectation
}

// Then sets up ChatRepository.Delete return parameters for the expectation previously defined by the When method
func (e *ChatRepositoryMockDeleteExpectation) Then(err error) *ChatRepositoryMock {
	e.results = &ChatRepositoryMockDeleteResults{err}
	return e.mock
}

// Delete implements repository.ChatRepository
func (mmDelete *ChatRepositoryMock) Delete(ctx context.Context, id int64) (err error) {
	mm_atomic.AddUint64(&mmDelete.beforeDeleteCounter, 1)
	defer mm_atomic.AddUint64(&mmDelete.afterDeleteCounter, 1)

	if mmDelete.inspectFuncDelete != nil {
		mmDelete.inspectFuncDelete(ctx, id)
	}

	mm_params := &ChatRepositoryMockDeleteParams{ctx, id}

	// Record call args
	mmDelete.DeleteMock.mutex.Lock()
	mmDelete.DeleteMock.callArgs = append(mmDelete.DeleteMock.callArgs, mm_params)
	mmDelete.DeleteMock.mutex.Unlock()

	for _, e := range mmDelete.DeleteMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDelete.DeleteMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDelete.DeleteMock.defaultExpectation.Counter, 1)
		mm_want := mmDelete.DeleteMock.defaultExpectation.params
		mm_got := ChatRepositoryMockDeleteParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDelete.t.Errorf("ChatRepositoryMock.Delete got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDelete.DeleteMock.defaultExpectation.results
		if mm_results == nil {
			mmDelete.t.Fatal("No results are set for the ChatRepositoryMock.Delete")
		}
		return (*mm_results).err
	}
	if mmDelete.funcDelete != nil {
		return mmDelete.funcDelete(ctx, id)
	}
	mmDelete.t.Fatalf("Unexpected call to ChatRepositoryMock.Delete. %v %v", ctx, id)
	return
}

// DeleteAfterCounter returns a count of finished ChatRepositoryMock.Delete invocations
func (mmDelete *ChatRepositoryMock) DeleteAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.afterDeleteCounter)
}

// DeleteBeforeCounter returns a count of ChatRepositoryMock.Delete invocations
func (mmDelete *ChatRepositoryMock) DeleteBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDelete.beforeDeleteCounter)
}

// Calls returns a list of arguments used in each call to ChatRepositoryMock.Delete.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDelete *mChatRepositoryMockDelete) Calls() []*ChatRepositoryMockDeleteParams {
	mmDelete.mutex.RLock()

	argCopy := make([]*ChatRepositoryMockDeleteParams, len(mmDelete.callArgs))
	copy(argCopy, mmDelete.callArgs)

	mmDelete.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteDone returns true if the count of the Delete invocations corresponds
// the number of defined expectations
func (m *ChatRepositoryMock) MinimockDeleteDone() bool {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteInspect logs each unmet expectation
func (m *ChatRepositoryMock) MinimockDeleteInspect() {
	for _, e := range m.DeleteMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatRepositoryMock.Delete with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		if m.DeleteMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatRepositoryMock.Delete")
		} else {
			m.t.Errorf("Expected call to ChatRepositoryMock.Delete with params: %#v", *m.DeleteMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDelete != nil && mm_atomic.LoadUint64(&m.afterDeleteCounter) < 1 {
		m.t.Error("Expected call to ChatRepositoryMock.Delete")
	}
}

type mChatRepositoryMockGet struct {
	mock               *ChatRepositoryMock
	defaultExpectation *ChatRepositoryMockGetExpectation
	expectations       []*ChatRepositoryMockGetExpectation

	callArgs []*ChatRepositoryMockGetParams
	mutex    sync.RWMutex
}

// ChatRepositoryMockGetExpectation specifies expectation struct of the ChatRepository.Get
type ChatRepositoryMockGetExpectation struct {
	mock    *ChatRepositoryMock
	params  *ChatRepositoryMockGetParams
	results *ChatRepositoryMockGetResults
	Counter uint64
}

// ChatRepositoryMockGetParams contains parameters of the ChatRepository.Get
type ChatRepositoryMockGetParams struct {
	ctx context.Context
	id  int64
}

// ChatRepositoryMockGetResults contains results of the ChatRepository.Get
type ChatRepositoryMockGetResults struct {
	cp1 *model.Chat
	err error
}

// Expect sets up expected params for ChatRepository.Get
func (mmGet *mChatRepositoryMockGet) Expect(ctx context.Context, id int64) *mChatRepositoryMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("ChatRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &ChatRepositoryMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &ChatRepositoryMockGetParams{ctx, id}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the ChatRepository.Get
func (mmGet *mChatRepositoryMockGet) Inspect(f func(ctx context.Context, id int64)) *mChatRepositoryMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for ChatRepositoryMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by ChatRepository.Get
func (mmGet *mChatRepositoryMockGet) Return(cp1 *model.Chat, err error) *ChatRepositoryMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("ChatRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &ChatRepositoryMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &ChatRepositoryMockGetResults{cp1, err}
	return mmGet.mock
}

// Set uses given function f to mock the ChatRepository.Get method
func (mmGet *mChatRepositoryMockGet) Set(f func(ctx context.Context, id int64) (cp1 *model.Chat, err error)) *ChatRepositoryMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the ChatRepository.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the ChatRepository.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the ChatRepository.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mChatRepositoryMockGet) When(ctx context.Context, id int64) *ChatRepositoryMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("ChatRepositoryMock.Get mock is already set by Set")
	}

	expectation := &ChatRepositoryMockGetExpectation{
		mock:   mmGet.mock,
		params: &ChatRepositoryMockGetParams{ctx, id},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up ChatRepository.Get return parameters for the expectation previously defined by the When method
func (e *ChatRepositoryMockGetExpectation) Then(cp1 *model.Chat, err error) *ChatRepositoryMock {
	e.results = &ChatRepositoryMockGetResults{cp1, err}
	return e.mock
}

// Get implements repository.ChatRepository
func (mmGet *ChatRepositoryMock) Get(ctx context.Context, id int64) (cp1 *model.Chat, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, id)
	}

	mm_params := &ChatRepositoryMockGetParams{ctx, id}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.cp1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := ChatRepositoryMockGetParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("ChatRepositoryMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the ChatRepositoryMock.Get")
		}
		return (*mm_results).cp1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, id)
	}
	mmGet.t.Fatalf("Unexpected call to ChatRepositoryMock.Get. %v %v", ctx, id)
	return
}

// GetAfterCounter returns a count of finished ChatRepositoryMock.Get invocations
func (mmGet *ChatRepositoryMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of ChatRepositoryMock.Get invocations
func (mmGet *ChatRepositoryMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to ChatRepositoryMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mChatRepositoryMockGet) Calls() []*ChatRepositoryMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*ChatRepositoryMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *ChatRepositoryMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *ChatRepositoryMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatRepositoryMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatRepositoryMock.Get")
		} else {
			m.t.Errorf("Expected call to ChatRepositoryMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to ChatRepositoryMock.Get")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ChatRepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockDeleteInspect()

		m.MinimockGetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ChatRepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *ChatRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteDone() &&
		m.MinimockGetDone()
}
