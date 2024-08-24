package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/nqxcode/chat_microservice/internal/repository.MessageRepository -o ./mocks\message_repository_minimock.go -n MessageRepositoryMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
	"github.com/nqxcode/chat_microservice/internal/model"
	"github.com/nqxcode/platform_common/pagination"
)

// MessageRepositoryMock implements repository.MessageRepository
type MessageRepositoryMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, model *model.Message) (i1 int64, err error)
	inspectFuncCreate   func(ctx context.Context, model *model.Message)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mMessageRepositoryMockCreate

	funcDeleteByChatID          func(ctx context.Context, chatID int64) (err error)
	inspectFuncDeleteByChatID   func(ctx context.Context, chatID int64)
	afterDeleteByChatIDCounter  uint64
	beforeDeleteByChatIDCounter uint64
	DeleteByChatIDMock          mMessageRepositoryMockDeleteByChatID

	funcGet          func(ctx context.Context, chatID int64, limit *pagination.Limit) (ma1 []model.Message, err error)
	inspectFuncGet   func(ctx context.Context, chatID int64, limit *pagination.Limit)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mMessageRepositoryMockGet
}

// NewMessageRepositoryMock returns a mock for repository.MessageRepository
func NewMessageRepositoryMock(t minimock.Tester) *MessageRepositoryMock {
	m := &MessageRepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mMessageRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*MessageRepositoryMockCreateParams{}

	m.DeleteByChatIDMock = mMessageRepositoryMockDeleteByChatID{mock: m}
	m.DeleteByChatIDMock.callArgs = []*MessageRepositoryMockDeleteByChatIDParams{}

	m.GetMock = mMessageRepositoryMockGet{mock: m}
	m.GetMock.callArgs = []*MessageRepositoryMockGetParams{}

	return m
}

type mMessageRepositoryMockCreate struct {
	mock               *MessageRepositoryMock
	defaultExpectation *MessageRepositoryMockCreateExpectation
	expectations       []*MessageRepositoryMockCreateExpectation

	callArgs []*MessageRepositoryMockCreateParams
	mutex    sync.RWMutex
}

// MessageRepositoryMockCreateExpectation specifies expectation struct of the MessageRepository.Create
type MessageRepositoryMockCreateExpectation struct {
	mock    *MessageRepositoryMock
	params  *MessageRepositoryMockCreateParams
	results *MessageRepositoryMockCreateResults
	Counter uint64
}

// MessageRepositoryMockCreateParams contains parameters of the MessageRepository.Create
type MessageRepositoryMockCreateParams struct {
	ctx   context.Context
	model *model.Message
}

// MessageRepositoryMockCreateResults contains results of the MessageRepository.Create
type MessageRepositoryMockCreateResults struct {
	i1  int64
	err error
}

// Expect sets up expected params for MessageRepository.Create
func (mmCreate *mMessageRepositoryMockCreate) Expect(ctx context.Context, model *model.Message) *mMessageRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("MessageRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &MessageRepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &MessageRepositoryMockCreateParams{ctx, model}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the MessageRepository.Create
func (mmCreate *mMessageRepositoryMockCreate) Inspect(f func(ctx context.Context, model *model.Message)) *mMessageRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for MessageRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by MessageRepository.Create
func (mmCreate *mMessageRepositoryMockCreate) Return(i1 int64, err error) *MessageRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("MessageRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &MessageRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &MessageRepositoryMockCreateResults{i1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the MessageRepository.Create method
func (mmCreate *mMessageRepositoryMockCreate) Set(f func(ctx context.Context, model *model.Message) (i1 int64, err error)) *MessageRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the MessageRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the MessageRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the MessageRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mMessageRepositoryMockCreate) When(ctx context.Context, model *model.Message) *MessageRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("MessageRepositoryMock.Create mock is already set by Set")
	}

	expectation := &MessageRepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &MessageRepositoryMockCreateParams{ctx, model},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up MessageRepository.Create return parameters for the expectation previously defined by the When method
func (e *MessageRepositoryMockCreateExpectation) Then(i1 int64, err error) *MessageRepositoryMock {
	e.results = &MessageRepositoryMockCreateResults{i1, err}
	return e.mock
}

// Create implements repository.MessageRepository
func (mmCreate *MessageRepositoryMock) Create(ctx context.Context, model *model.Message) (i1 int64, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, model)
	}

	mm_params := &MessageRepositoryMockCreateParams{ctx, model}

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
		mm_got := MessageRepositoryMockCreateParams{ctx, model}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("MessageRepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the MessageRepositoryMock.Create")
		}
		return (*mm_results).i1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, model)
	}
	mmCreate.t.Fatalf("Unexpected call to MessageRepositoryMock.Create. %v %v", ctx, model)
	return
}

// CreateAfterCounter returns a count of finished MessageRepositoryMock.Create invocations
func (mmCreate *MessageRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of MessageRepositoryMock.Create invocations
func (mmCreate *MessageRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to MessageRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mMessageRepositoryMockCreate) Calls() []*MessageRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*MessageRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *MessageRepositoryMock) MinimockCreateDone() bool {
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
func (m *MessageRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageRepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageRepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to MessageRepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to MessageRepositoryMock.Create")
	}
}

type mMessageRepositoryMockDeleteByChatID struct {
	mock               *MessageRepositoryMock
	defaultExpectation *MessageRepositoryMockDeleteByChatIDExpectation
	expectations       []*MessageRepositoryMockDeleteByChatIDExpectation

	callArgs []*MessageRepositoryMockDeleteByChatIDParams
	mutex    sync.RWMutex
}

// MessageRepositoryMockDeleteByChatIDExpectation specifies expectation struct of the MessageRepository.DeleteByChatID
type MessageRepositoryMockDeleteByChatIDExpectation struct {
	mock    *MessageRepositoryMock
	params  *MessageRepositoryMockDeleteByChatIDParams
	results *MessageRepositoryMockDeleteByChatIDResults
	Counter uint64
}

// MessageRepositoryMockDeleteByChatIDParams contains parameters of the MessageRepository.DeleteByChatID
type MessageRepositoryMockDeleteByChatIDParams struct {
	ctx    context.Context
	chatID int64
}

// MessageRepositoryMockDeleteByChatIDResults contains results of the MessageRepository.DeleteByChatID
type MessageRepositoryMockDeleteByChatIDResults struct {
	err error
}

// Expect sets up expected params for MessageRepository.DeleteByChatID
func (mmDeleteByChatID *mMessageRepositoryMockDeleteByChatID) Expect(ctx context.Context, chatID int64) *mMessageRepositoryMockDeleteByChatID {
	if mmDeleteByChatID.mock.funcDeleteByChatID != nil {
		mmDeleteByChatID.mock.t.Fatalf("MessageRepositoryMock.DeleteByChatID mock is already set by Set")
	}

	if mmDeleteByChatID.defaultExpectation == nil {
		mmDeleteByChatID.defaultExpectation = &MessageRepositoryMockDeleteByChatIDExpectation{}
	}

	mmDeleteByChatID.defaultExpectation.params = &MessageRepositoryMockDeleteByChatIDParams{ctx, chatID}
	for _, e := range mmDeleteByChatID.expectations {
		if minimock.Equal(e.params, mmDeleteByChatID.defaultExpectation.params) {
			mmDeleteByChatID.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmDeleteByChatID.defaultExpectation.params)
		}
	}

	return mmDeleteByChatID
}

// Inspect accepts an inspector function that has same arguments as the MessageRepository.DeleteByChatID
func (mmDeleteByChatID *mMessageRepositoryMockDeleteByChatID) Inspect(f func(ctx context.Context, chatID int64)) *mMessageRepositoryMockDeleteByChatID {
	if mmDeleteByChatID.mock.inspectFuncDeleteByChatID != nil {
		mmDeleteByChatID.mock.t.Fatalf("Inspect function is already set for MessageRepositoryMock.DeleteByChatID")
	}

	mmDeleteByChatID.mock.inspectFuncDeleteByChatID = f

	return mmDeleteByChatID
}

// Return sets up results that will be returned by MessageRepository.DeleteByChatID
func (mmDeleteByChatID *mMessageRepositoryMockDeleteByChatID) Return(err error) *MessageRepositoryMock {
	if mmDeleteByChatID.mock.funcDeleteByChatID != nil {
		mmDeleteByChatID.mock.t.Fatalf("MessageRepositoryMock.DeleteByChatID mock is already set by Set")
	}

	if mmDeleteByChatID.defaultExpectation == nil {
		mmDeleteByChatID.defaultExpectation = &MessageRepositoryMockDeleteByChatIDExpectation{mock: mmDeleteByChatID.mock}
	}
	mmDeleteByChatID.defaultExpectation.results = &MessageRepositoryMockDeleteByChatIDResults{err}
	return mmDeleteByChatID.mock
}

// Set uses given function f to mock the MessageRepository.DeleteByChatID method
func (mmDeleteByChatID *mMessageRepositoryMockDeleteByChatID) Set(f func(ctx context.Context, chatID int64) (err error)) *MessageRepositoryMock {
	if mmDeleteByChatID.defaultExpectation != nil {
		mmDeleteByChatID.mock.t.Fatalf("Default expectation is already set for the MessageRepository.DeleteByChatID method")
	}

	if len(mmDeleteByChatID.expectations) > 0 {
		mmDeleteByChatID.mock.t.Fatalf("Some expectations are already set for the MessageRepository.DeleteByChatID method")
	}

	mmDeleteByChatID.mock.funcDeleteByChatID = f
	return mmDeleteByChatID.mock
}

// When sets expectation for the MessageRepository.DeleteByChatID which will trigger the result defined by the following
// Then helper
func (mmDeleteByChatID *mMessageRepositoryMockDeleteByChatID) When(ctx context.Context, chatID int64) *MessageRepositoryMockDeleteByChatIDExpectation {
	if mmDeleteByChatID.mock.funcDeleteByChatID != nil {
		mmDeleteByChatID.mock.t.Fatalf("MessageRepositoryMock.DeleteByChatID mock is already set by Set")
	}

	expectation := &MessageRepositoryMockDeleteByChatIDExpectation{
		mock:   mmDeleteByChatID.mock,
		params: &MessageRepositoryMockDeleteByChatIDParams{ctx, chatID},
	}
	mmDeleteByChatID.expectations = append(mmDeleteByChatID.expectations, expectation)
	return expectation
}

// Then sets up MessageRepository.DeleteByChatID return parameters for the expectation previously defined by the When method
func (e *MessageRepositoryMockDeleteByChatIDExpectation) Then(err error) *MessageRepositoryMock {
	e.results = &MessageRepositoryMockDeleteByChatIDResults{err}
	return e.mock
}

// DeleteByChatID implements repository.MessageRepository
func (mmDeleteByChatID *MessageRepositoryMock) DeleteByChatID(ctx context.Context, chatID int64) (err error) {
	mm_atomic.AddUint64(&mmDeleteByChatID.beforeDeleteByChatIDCounter, 1)
	defer mm_atomic.AddUint64(&mmDeleteByChatID.afterDeleteByChatIDCounter, 1)

	if mmDeleteByChatID.inspectFuncDeleteByChatID != nil {
		mmDeleteByChatID.inspectFuncDeleteByChatID(ctx, chatID)
	}

	mm_params := &MessageRepositoryMockDeleteByChatIDParams{ctx, chatID}

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
		mm_got := MessageRepositoryMockDeleteByChatIDParams{ctx, chatID}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmDeleteByChatID.t.Errorf("MessageRepositoryMock.DeleteByChatID got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmDeleteByChatID.DeleteByChatIDMock.defaultExpectation.results
		if mm_results == nil {
			mmDeleteByChatID.t.Fatal("No results are set for the MessageRepositoryMock.DeleteByChatID")
		}
		return (*mm_results).err
	}
	if mmDeleteByChatID.funcDeleteByChatID != nil {
		return mmDeleteByChatID.funcDeleteByChatID(ctx, chatID)
	}
	mmDeleteByChatID.t.Fatalf("Unexpected call to MessageRepositoryMock.DeleteByChatID. %v %v", ctx, chatID)
	return
}

// DeleteByChatIDAfterCounter returns a count of finished MessageRepositoryMock.DeleteByChatID invocations
func (mmDeleteByChatID *MessageRepositoryMock) DeleteByChatIDAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteByChatID.afterDeleteByChatIDCounter)
}

// DeleteByChatIDBeforeCounter returns a count of MessageRepositoryMock.DeleteByChatID invocations
func (mmDeleteByChatID *MessageRepositoryMock) DeleteByChatIDBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmDeleteByChatID.beforeDeleteByChatIDCounter)
}

// Calls returns a list of arguments used in each call to MessageRepositoryMock.DeleteByChatID.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmDeleteByChatID *mMessageRepositoryMockDeleteByChatID) Calls() []*MessageRepositoryMockDeleteByChatIDParams {
	mmDeleteByChatID.mutex.RLock()

	argCopy := make([]*MessageRepositoryMockDeleteByChatIDParams, len(mmDeleteByChatID.callArgs))
	copy(argCopy, mmDeleteByChatID.callArgs)

	mmDeleteByChatID.mutex.RUnlock()

	return argCopy
}

// MinimockDeleteByChatIDDone returns true if the count of the DeleteByChatID invocations corresponds
// the number of defined expectations
func (m *MessageRepositoryMock) MinimockDeleteByChatIDDone() bool {
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
func (m *MessageRepositoryMock) MinimockDeleteByChatIDInspect() {
	for _, e := range m.DeleteByChatIDMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageRepositoryMock.DeleteByChatID with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.DeleteByChatIDMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterDeleteByChatIDCounter) < 1 {
		if m.DeleteByChatIDMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageRepositoryMock.DeleteByChatID")
		} else {
			m.t.Errorf("Expected call to MessageRepositoryMock.DeleteByChatID with params: %#v", *m.DeleteByChatIDMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcDeleteByChatID != nil && mm_atomic.LoadUint64(&m.afterDeleteByChatIDCounter) < 1 {
		m.t.Error("Expected call to MessageRepositoryMock.DeleteByChatID")
	}
}

type mMessageRepositoryMockGet struct {
	mock               *MessageRepositoryMock
	defaultExpectation *MessageRepositoryMockGetExpectation
	expectations       []*MessageRepositoryMockGetExpectation

	callArgs []*MessageRepositoryMockGetParams
	mutex    sync.RWMutex
}

// MessageRepositoryMockGetExpectation specifies expectation struct of the MessageRepository.Get
type MessageRepositoryMockGetExpectation struct {
	mock    *MessageRepositoryMock
	params  *MessageRepositoryMockGetParams
	results *MessageRepositoryMockGetResults
	Counter uint64
}

// MessageRepositoryMockGetParams contains parameters of the MessageRepository.Get
type MessageRepositoryMockGetParams struct {
	ctx    context.Context
	chatID int64
	limit  *pagination.Limit
}

// MessageRepositoryMockGetResults contains results of the MessageRepository.Get
type MessageRepositoryMockGetResults struct {
	ma1 []model.Message
	err error
}

// Expect sets up expected params for MessageRepository.Get
func (mmGet *mMessageRepositoryMockGet) Expect(ctx context.Context, chatID int64, limit *pagination.Limit) *mMessageRepositoryMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("MessageRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &MessageRepositoryMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &MessageRepositoryMockGetParams{ctx, chatID, limit}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the MessageRepository.Get
func (mmGet *mMessageRepositoryMockGet) Inspect(f func(ctx context.Context, chatID int64, limit *pagination.Limit)) *mMessageRepositoryMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for MessageRepositoryMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by MessageRepository.Get
func (mmGet *mMessageRepositoryMockGet) Return(ma1 []model.Message, err error) *MessageRepositoryMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("MessageRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &MessageRepositoryMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &MessageRepositoryMockGetResults{ma1, err}
	return mmGet.mock
}

// Set uses given function f to mock the MessageRepository.Get method
func (mmGet *mMessageRepositoryMockGet) Set(f func(ctx context.Context, chatID int64, limit *pagination.Limit) (ma1 []model.Message, err error)) *MessageRepositoryMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the MessageRepository.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the MessageRepository.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the MessageRepository.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mMessageRepositoryMockGet) When(ctx context.Context, chatID int64, limit *pagination.Limit) *MessageRepositoryMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("MessageRepositoryMock.Get mock is already set by Set")
	}

	expectation := &MessageRepositoryMockGetExpectation{
		mock:   mmGet.mock,
		params: &MessageRepositoryMockGetParams{ctx, chatID, limit},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up MessageRepository.Get return parameters for the expectation previously defined by the When method
func (e *MessageRepositoryMockGetExpectation) Then(ma1 []model.Message, err error) *MessageRepositoryMock {
	e.results = &MessageRepositoryMockGetResults{ma1, err}
	return e.mock
}

// Get implements repository.MessageRepository
func (mmGet *MessageRepositoryMock) Get(ctx context.Context, chatID int64, limit *pagination.Limit) (ma1 []model.Message, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, chatID, limit)
	}

	mm_params := &MessageRepositoryMockGetParams{ctx, chatID, limit}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.ma1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := MessageRepositoryMockGetParams{ctx, chatID, limit}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("MessageRepositoryMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the MessageRepositoryMock.Get")
		}
		return (*mm_results).ma1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, chatID, limit)
	}
	mmGet.t.Fatalf("Unexpected call to MessageRepositoryMock.Get. %v %v %v", ctx, chatID, limit)
	return
}

// GetAfterCounter returns a count of finished MessageRepositoryMock.Get invocations
func (mmGet *MessageRepositoryMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of MessageRepositoryMock.Get invocations
func (mmGet *MessageRepositoryMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to MessageRepositoryMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mMessageRepositoryMockGet) Calls() []*MessageRepositoryMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*MessageRepositoryMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *MessageRepositoryMock) MinimockGetDone() bool {
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
func (m *MessageRepositoryMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MessageRepositoryMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MessageRepositoryMock.Get")
		} else {
			m.t.Errorf("Expected call to MessageRepositoryMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to MessageRepositoryMock.Get")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MessageRepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockDeleteByChatIDInspect()

		m.MinimockGetInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MessageRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *MessageRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockDeleteByChatIDDone() &&
		m.MinimockGetDone()
}
