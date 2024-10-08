package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/nqxcode/chat_microservice/internal/repository.ChatToUserRepository -o ./mocks\chat_to_user_repository_minimock.go -n ChatToUserRepositoryMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/platform_common/pagination"
)

// ChatToUserRepositoryMock implements repository.ChatToUserRepository
type ChatToUserRepositoryMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, model *model.ChatToUser) (i1 int64, err error)
	inspectFuncCreate   func(ctx context.Context, model *model.ChatToUser)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mChatToUserRepositoryMockCreate

	funcDeleteByChatID          func(ctx context.Context, chatID int64) (err error)
	inspectFuncDeleteByChatID   func(ctx context.Context, chatID int64)
	afterDeleteByChatIDCounter  uint64
	beforeDeleteByChatIDCounter uint64
	DeleteByChatIDMock          mChatToUserRepositoryMockDeleteByChatID

	funcGet          func(ctx context.Context, chatID int64, limit pagination.Limit) (ca1 []model.ChatToUser, err error)
	inspectFuncGet   func(ctx context.Context, chatID int64, limit pagination.Limit)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mChatToUserRepositoryMockGet
}

// NewChatToUserRepositoryMock returns a mock for repository.ChatToUserRepository
func NewChatToUserRepositoryMock(t minimock.Tester) *ChatToUserRepositoryMock {
	m := &ChatToUserRepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mChatToUserRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*ChatToUserRepositoryMockCreateParams{}

	m.DeleteByChatIDMock = mChatToUserRepositoryMockDeleteByChatID{mock: m}
	m.DeleteByChatIDMock.callArgs = []*ChatToUserRepositoryMockDeleteByChatIDParams{}

	m.GetMock = mChatToUserRepositoryMockGet{mock: m}
	m.GetMock.callArgs = []*ChatToUserRepositoryMockGetParams{}

	return m
}

type mChatToUserRepositoryMockCreate struct {
	mock               *ChatToUserRepositoryMock
	defaultExpectation *ChatToUserRepositoryMockCreateExpectation
	expectations       []*ChatToUserRepositoryMockCreateExpectation

	callArgs []*ChatToUserRepositoryMockCreateParams
	mutex    sync.RWMutex
}

// ChatToUserRepositoryMockCreateExpectation specifies expectation struct of the ChatToUserRepository.Create
type ChatToUserRepositoryMockCreateExpectation struct {
	mock    *ChatToUserRepositoryMock
	params  *ChatToUserRepositoryMockCreateParams
	results *ChatToUserRepositoryMockCreateResults
	Counter uint64
}

// ChatToUserRepositoryMockCreateParams contains parameters of the ChatToUserRepository.Create
type ChatToUserRepositoryMockCreateParams struct {
	ctx   context.Context
	model *model.ChatToUser
}

// ChatToUserRepositoryMockCreateResults contains results of the ChatToUserRepository.Create
type ChatToUserRepositoryMockCreateResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for ChatToUserRepository.Create
func (mmCreate *mChatToUserRepositoryMockCreate) Expect(ctx context.Context, model *model.ChatToUser) *mChatToUserRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatToUserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatToUserRepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &ChatToUserRepositoryMockCreateParams{ctx, model}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the ChatToUserRepository.Create
func (mmCreate *mChatToUserRepositoryMockCreate) Inspect(f func(ctx context.Context, model *model.ChatToUser)) *mChatToUserRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for ChatToUserRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by ChatToUserRepository.Create
func (mmCreate *mChatToUserRepositoryMockCreate) Return(i1 int64, err error) *ChatToUserRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatToUserRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &ChatToUserRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &ChatToUserRepositoryMockCreateResults{i1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the ChatToUserRepository.Create method
func (mmCreate *mChatToUserRepositoryMockCreate) Set(f func(ctx context.Context, model *model.ChatToUser) (i1 int64, err error)) *ChatToUserRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the ChatToUserRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the ChatToUserRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the ChatToUserRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mChatToUserRepositoryMockCreate) When(ctx context.Context, model *model.ChatToUser) *ChatToUserRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("ChatToUserRepositoryMock.Create mock is already set by Set")
	}

	expectation := &ChatToUserRepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &ChatToUserRepositoryMockCreateParams{ctx, model},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up ChatToUserRepository.Create return parameters for the expectation previously defined by the When method
func (e *ChatToUserRepositoryMockCreateExpectation) Then(i1 int64, err error) *ChatToUserRepositoryMock {
	e.results = &ChatToUserRepositoryMockCreateResults{i1, err}
	return e.mock
}

// Create implements repository.ChatToUserRepository
func (mmCreate *ChatToUserRepositoryMock) Create(ctx context.Context, model *model.ChatToUser) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, model)
	}

	mm_params := &ChatToUserRepositoryMockCreateParams{ctx, model}

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
		mm_got := ChatToUserRepositoryMockCreateParams{ctx, model}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("ChatToUserRepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the ChatToUserRepositoryMock.Create")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, model)
	}
	mmCreate.t.Fatalf("Unexpected call to ChatToUserRepositoryMock.Create. %v %v", ctx, model)
	return
}

// CreateAfterCounter returns a count of finished ChatToUserRepositoryMock.Create invocations
func (mmCreate *ChatToUserRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of ChatToUserRepositoryMock.Create invocations
func (mmCreate *ChatToUserRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to ChatToUserRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mChatToUserRepositoryMockCreate) Calls() []*ChatToUserRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*ChatToUserRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *ChatToUserRepositoryMock) MinimockCreateDone() bool {
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
func (m *ChatToUserRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatToUserRepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatToUserRepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to ChatToUserRepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to ChatToUserRepositoryMock.Create")
	}
}

type mChatToUserRepositoryMockDeleteByChatID struct {
	mock               *ChatToUserRepositoryMock
	defaultExpectation *ChatToUserRepositoryMockDeleteByChatIDExpectation
	expectations       []*ChatToUserRepositoryMockDeleteByChatIDExpectation

	callArgs []*ChatToUserRepositoryMockDeleteByChatIDParams
	mutex    sync.RWMutex
}

// ChatToUserRepositoryMockDeleteByChatIDExpectation specifies expectation struct of the ChatToUserRepository.DeleteByChatID
type ChatToUserRepositoryMockDeleteByChatIDExpectation struct {
	mock    *ChatToUserRepositoryMock
	params  *ChatToUserRepositoryMockDeleteByChatIDParams
	results *ChatToUserRepositoryMockDeleteByChatIDResults
	Counter uint64
}

// ChatToUserRepositoryMockDeleteByChatIDParams contains parameters of the ChatToUserRepository.DeleteByChatID
type ChatToUserRepositoryMockDeleteByChatIDParams struct {
	ctx    context.Context
	chatID int64
}

// ChatToUserRepositoryMockDeleteByChatIDResults contains results of the ChatToUserRepository.DeleteByChatID
type ChatToUserRepositoryMockDeleteByChatIDResults struct {
	err error
}

// Expect sets up expected params for ChatToUserRepository.DeleteByChatID
func (mmDeleteByChatID *mChatToUserRepositoryMockDeleteByChatID) Expect(ctx context.Context, chatID int64) *mChatToUserRepositoryMockDeleteByChatID {
	if mmDeleteByChatID.mock.funcDeleteByChatID != nil {
		mmDeleteByChatID.mock.t.Fatalf("ChatToUserRepositoryMock.DeleteByChatID mock is already set by Set")
	}

	if mmDeleteByChatID.defaultExpectation == nil {
		mmDeleteByChatID.defaultExpectation = &ChatToUserRepositoryMockDeleteByChatIDExpectation{}
	}

	mmDeleteByChatID.defaultExpectation.params = &ChatToUserRepositoryMockDeleteByChatIDParams{ctx, chatID}
	for _, e := range mmDeleteByChatID.expectations {
		if minimock.Equal(e.params, mmDeleteByChatID.defaultExpectation.params) {
			mmDeleteByChatID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeleteByChatID.defaultExpectation.params)
		}
	}

	return mmDeleteByChatID
}

// Inspect accepts an inspector function that has same arguments as the ChatToUserRepository.DeleteByChatID
func (mmDeleteByChatID *mChatToUserRepositoryMockDeleteByChatID) Inspect(f func(ctx context.Context, chatID int64)) *mChatToUserRepositoryMockDeleteByChatID {
	if mmDeleteByChatID.mock.inspectFuncDeleteByChatID != nil {
		mmDeleteByChatID.mock.t.Fatalf("Inspect function is already set for ChatToUserRepositoryMock.DeleteByChatID")
	}

	mmDeleteByChatID.mock.inspectFuncDeleteByChatID = f

	return mmDeleteByChatID
}

// Return sets up results that will be returned by ChatToUserRepository.DeleteByChatID
func (mmDeleteByChatID *mChatToUserRepositoryMockDeleteByChatID) Return(err error) *ChatToUserRepositoryMock {
	if mmDeleteByChatID.mock.funcDeleteByChatID != nil {
		mmDeleteByChatID.mock.t.Fatalf("ChatToUserRepositoryMock.DeleteByChatID mock is already set by Set")
	}

	if mmDeleteByChatID.defaultExpectation == nil {
		mmDeleteByChatID.defaultExpectation = &ChatToUserRepositoryMockDeleteByChatIDExpectation{mock: mmDeleteByChatID.mock}
	}
	mmDeleteByChatID.defaultExpectation.results = &ChatToUserRepositoryMockDeleteByChatIDResults{err}
	return mmDeleteByChatID.mock
}

// Set uses given function f to mock the ChatToUserRepository.DeleteByChatID method
func (mmDeleteByChatID *mChatToUserRepositoryMockDeleteByChatID) Set(f func(ctx context.Context, chatID int64) (err error)) *ChatToUserRepositoryMock {
	if mmDeleteByChatID.defaultExpectation != nil {
		mmDeleteByChatID.mock.t.Fatalf("Default expectation is already set for the ChatToUserRepository.DeleteByChatID method")
	}

	if len(mmDeleteByChatID.expectations) > 0 {
		mmDeleteByChatID.mock.t.Fatalf("Some expectations are already set for the ChatToUserRepository.DeleteByChatID method")
	}

	mmDeleteByChatID.mock.funcDeleteByChatID = f
	return mmDeleteByChatID.mock
}

// When sets expectation for the ChatToUserRepository.DeleteByChatID which will trigger the result defined by the following
// Then helper
func (mmDeleteByChatID *mChatToUserRepositoryMockDeleteByChatID) When(ctx context.Context, chatID int64) *ChatToUserRepositoryMockDeleteByChatIDExpectation {
	if mmDeleteByChatID.mock.funcDeleteByChatID != nil {
		mmDeleteByChatID.mock.t.Fatalf("ChatToUserRepositoryMock.DeleteByChatID mock is already set by Set")
	}

	expectation := &ChatToUserRepositoryMockDeleteByChatIDExpectation{
		mock:   mmDeleteByChatID.mock,
		params: &ChatToUserRepositoryMockDeleteByChatIDParams{ctx, chatID},
	}
	mmDeleteByChatID.expectations = append(mmDeleteByChatID.expectations, expectation)
	return expectation
}

// Then sets up ChatToUserRepository.DeleteByChatID return parameters for the expectation previously defined by the When method
func (e *ChatToUserRepositoryMockDeleteByChatIDExpectation) Then(err error) *ChatToUserRepositoryMock {
	e.results = &ChatToUserRepositoryMockDeleteByChatIDResults{err}
	return e.mock
}

// DeleteByChatID implements repository.ChatToUserRepository
func (mmDeleteByChatID *ChatToUserRepositoryMock) DeleteByChatID(ctx context.Context, chatID int64) (err error) {
	mm_atomic.AddUint64(&mmDeleteByChatID.beforeDeleteByChatIDCounter, 1)
	defer mm_atomic.AddUint64(&mmDeleteByChatID.afterDeleteByChatIDCounter, 1)

	if mmDeleteByChatID.inspectFuncDeleteByChatID != nil {
		mmDeleteByChatID.inspectFuncDeleteByChatID(ctx, chatID)
	}

	mm_params := &ChatToUserRepositoryMockDeleteByChatIDParams{ctx, chatID}

	// Record call args
	mmDeleteByChatID.DeleteByChatIDMock.mutex.Lock()
	mmDeleteByChatID.DeleteByChatIDMock.callArgs = append(mmDeleteByChatID.DeleteByChatIDMock.callArgs, mm_params)
	mmDeleteByChatID.DeleteByChatIDMock.mutex.Unlock()

	for _, e := range mmDeleteByChatID.DeleteByChatIDMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmDeleteByChatID.DeleteByChatIDMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmDeleteByChatID.DeleteByChatIDMock.defaultExpectation.Counter, 1)
		mm_want := mmDeleteByChatID.DeleteByChatIDMock.defaultExpectation.params
		mm_got := ChatToUserRepositoryMockDeleteByChatIDParams{ctx, chatID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeleteByChatID.t.Errorf("ChatToUserRepositoryMock.DeleteByChatID got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeleteByChatID.DeleteByChatIDMock.defaultExpectation.results
		if mm_results == nil {
			mmDeleteByChatID.t.Fatal("No results are set for the ChatToUserRepositoryMock.DeleteByChatID")
		}
		return (*mm_results).err
	}
	if mmDeleteByChatID.funcDeleteByChatID != nil {
		return mmDeleteByChatID.funcDeleteByChatID(ctx, chatID)
	}
	mmDeleteByChatID.t.Fatalf("Unexpected call to ChatToUserRepositoryMock.DeleteByChatID. %v %v", ctx, chatID)
	return
}

// DeleteByChatIDAfterCounter returns a count of finished ChatToUserRepositoryMock.DeleteByChatID invocations
func (mmDeleteByChatID *ChatToUserRepositoryMock) DeleteByChatIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteByChatID.afterDeleteByChatIDCounter)
}

// DeleteByChatIDBeforeCounter returns a count of ChatToUserRepositoryMock.DeleteByChatID invocations
func (mmDeleteByChatID *ChatToUserRepositoryMock) DeleteByChatIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteByChatID.beforeDeleteByChatIDCounter)
}

// Calls returns a list of arguments used in each call to ChatToUserRepositoryMock.DeleteByChatID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeleteByChatID *mChatToUserRepositoryMockDeleteByChatID) Calls() []*ChatToUserRepositoryMockDeleteByChatIDParams {
	mmDeleteByChatID.mutex.RLock()

	argCopy := make([]*ChatToUserRepositoryMockDeleteByChatIDParams, len(mmDeleteByChatID.callArgs))
	copy(argCopy, mmDeleteByChatID.callArgs)

	mmDeleteByChatID.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteByChatIDDone returns true if the count of the DeleteByChatID invocations corresponds
// the number of defined expectations
func (m *ChatToUserRepositoryMock) MinimockDeleteByChatIDDone() bool {
	for _, e := range m.DeleteByChatIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteByChatIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteByChatIDCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteByChatID != nil && mm_atomic.LoadUint64(&m.afterDeleteByChatIDCounter) < 1 {
		return false
	}
	return true
}

// MinimockDeleteByChatIDInspect logs each unmet expectation
func (m *ChatToUserRepositoryMock) MinimockDeleteByChatIDInspect() {
	for _, e := range m.DeleteByChatIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatToUserRepositoryMock.DeleteByChatID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteByChatIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteByChatIDCounter) < 1 {
		if m.DeleteByChatIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatToUserRepositoryMock.DeleteByChatID")
		} else {
			m.t.Errorf("Expected call to ChatToUserRepositoryMock.DeleteByChatID with params: %#v", *m.DeleteByChatIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteByChatID != nil && mm_atomic.LoadUint64(&m.afterDeleteByChatIDCounter) < 1 {
		m.t.Error("Expected call to ChatToUserRepositoryMock.DeleteByChatID")
	}
}

type mChatToUserRepositoryMockGet struct {
	mock               *ChatToUserRepositoryMock
	defaultExpectation *ChatToUserRepositoryMockGetExpectation
	expectations       []*ChatToUserRepositoryMockGetExpectation

	callArgs []*ChatToUserRepositoryMockGetParams
	mutex    sync.RWMutex
}

// ChatToUserRepositoryMockGetExpectation specifies expectation struct of the ChatToUserRepository.Get
type ChatToUserRepositoryMockGetExpectation struct {
	mock    *ChatToUserRepositoryMock
	params  *ChatToUserRepositoryMockGetParams
	results *ChatToUserRepositoryMockGetResults
	Counter uint64
}

// ChatToUserRepositoryMockGetParams contains parameters of the ChatToUserRepository.Get
type ChatToUserRepositoryMockGetParams struct {
	ctx    context.Context
	chatID int64
	limit  pagination.Limit
}

// ChatToUserRepositoryMockGetResults contains results of the ChatToUserRepository.Get
type ChatToUserRepositoryMockGetResults struct {
	ca1 []model.ChatToUser
	err error
}

// Expect sets up expected params for ChatToUserRepository.Get
func (mmGet *mChatToUserRepositoryMockGet) Expect(ctx context.Context, chatID int64, limit pagination.Limit) *mChatToUserRepositoryMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("ChatToUserRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &ChatToUserRepositoryMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &ChatToUserRepositoryMockGetParams{ctx, chatID, limit}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the ChatToUserRepository.Get
func (mmGet *mChatToUserRepositoryMockGet) Inspect(f func(ctx context.Context, chatID int64, limit pagination.Limit)) *mChatToUserRepositoryMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for ChatToUserRepositoryMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by ChatToUserRepository.Get
func (mmGet *mChatToUserRepositoryMockGet) Return(ca1 []model.ChatToUser, err error) *ChatToUserRepositoryMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("ChatToUserRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &ChatToUserRepositoryMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &ChatToUserRepositoryMockGetResults{ca1, err}
	return mmGet.mock
}

// Set uses given function f to mock the ChatToUserRepository.Get method
func (mmGet *mChatToUserRepositoryMockGet) Set(f func(ctx context.Context, chatID int64, limit pagination.Limit) (ca1 []model.ChatToUser, err error)) *ChatToUserRepositoryMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the ChatToUserRepository.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the ChatToUserRepository.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the ChatToUserRepository.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mChatToUserRepositoryMockGet) When(ctx context.Context, chatID int64, limit pagination.Limit) *ChatToUserRepositoryMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("ChatToUserRepositoryMock.Get mock is already set by Set")
	}

	expectation := &ChatToUserRepositoryMockGetExpectation{
		mock:   mmGet.mock,
		params: &ChatToUserRepositoryMockGetParams{ctx, chatID, limit},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up ChatToUserRepository.Get return parameters for the expectation previously defined by the When method
func (e *ChatToUserRepositoryMockGetExpectation) Then(ca1 []model.ChatToUser, err error) *ChatToUserRepositoryMock {
	e.results = &ChatToUserRepositoryMockGetResults{ca1, err}
	return e.mock
}

// Get implements repository.ChatToUserRepository
func (mmGet *ChatToUserRepositoryMock) Get(ctx context.Context, chatID int64, limit pagination.Limit) (ca1 []model.ChatToUser, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, chatID, limit)
	}

	mm_params := &ChatToUserRepositoryMockGetParams{ctx, chatID, limit}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ca1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := ChatToUserRepositoryMockGetParams{ctx, chatID, limit}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("ChatToUserRepositoryMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the ChatToUserRepositoryMock.Get")
		}
		return (*mm_results).ca1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, chatID, limit)
	}
	mmGet.t.Fatalf("Unexpected call to ChatToUserRepositoryMock.Get. %v %v %v", ctx, chatID, limit)
	return
}

// GetAfterCounter returns a count of finished ChatToUserRepositoryMock.Get invocations
func (mmGet *ChatToUserRepositoryMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of ChatToUserRepositoryMock.Get invocations
func (mmGet *ChatToUserRepositoryMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to ChatToUserRepositoryMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mChatToUserRepositoryMockGet) Calls() []*ChatToUserRepositoryMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*ChatToUserRepositoryMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *ChatToUserRepositoryMock) MinimockGetDone() bool {
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
func (m *ChatToUserRepositoryMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to ChatToUserRepositoryMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to ChatToUserRepositoryMock.Get")
		} else {
			m.t.Errorf("Expected call to ChatToUserRepositoryMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to ChatToUserRepositoryMock.Get")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *ChatToUserRepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockDeleteByChatIDInspect()

		m.MinimockGetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *ChatToUserRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *ChatToUserRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteByChatIDDone() &&
		m.MinimockGetDone()
}
