// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ledgerwatch/erigon/consensus (interfaces: Engine)
//
// Generated by this command:
//
//	mockgen -typed=true -destination=./engine_mock.go -package=consensus . Engine
//

// Package consensus is a generated GoMock package.
package consensus

import (
	big "math/big"
	reflect "reflect"

	chain "github.com/ledgerwatch/erigon-lib/chain"
	common "github.com/ledgerwatch/erigon-lib/common"
	state "github.com/ledgerwatch/erigon/core/state"
	types "github.com/ledgerwatch/erigon/core/types"
	rpc "github.com/ledgerwatch/erigon/rpc"
	log "github.com/ledgerwatch/log/v3"
	gomock "go.uber.org/mock/gomock"
)

// MockEngine is a mock of Engine interface.
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine.
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance.
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// APIs mocks base method.
func (m *MockEngine) APIs(arg0 ChainHeaderReader) []rpc.API {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APIs", arg0)
	ret0, _ := ret[0].([]rpc.API)
	return ret0
}

// APIs indicates an expected call of APIs.
func (mr *MockEngineMockRecorder) APIs(arg0 any) *MockEngineAPIsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIs", reflect.TypeOf((*MockEngine)(nil).APIs), arg0)
	return &MockEngineAPIsCall{Call: call}
}

// MockEngineAPIsCall wrap *gomock.Call
type MockEngineAPIsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineAPIsCall) Return(arg0 []rpc.API) *MockEngineAPIsCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineAPIsCall) Do(f func(ChainHeaderReader) []rpc.API) *MockEngineAPIsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineAPIsCall) DoAndReturn(f func(ChainHeaderReader) []rpc.API) *MockEngineAPIsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Author mocks base method.
func (m *MockEngine) Author(arg0 *types.Header) (common.Address, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Author", arg0)
	ret0, _ := ret[0].(common.Address)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Author indicates an expected call of Author.
func (mr *MockEngineMockRecorder) Author(arg0 any) *MockEngineAuthorCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Author", reflect.TypeOf((*MockEngine)(nil).Author), arg0)
	return &MockEngineAuthorCall{Call: call}
}

// MockEngineAuthorCall wrap *gomock.Call
type MockEngineAuthorCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineAuthorCall) Return(arg0 common.Address, arg1 error) *MockEngineAuthorCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineAuthorCall) Do(f func(*types.Header) (common.Address, error)) *MockEngineAuthorCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineAuthorCall) DoAndReturn(f func(*types.Header) (common.Address, error)) *MockEngineAuthorCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CalcDifficulty mocks base method.
func (m *MockEngine) CalcDifficulty(arg0 ChainHeaderReader, arg1, arg2 uint64, arg3 *big.Int, arg4 uint64, arg5, arg6 common.Hash, arg7 uint64) *big.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalcDifficulty", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// CalcDifficulty indicates an expected call of CalcDifficulty.
func (mr *MockEngineMockRecorder) CalcDifficulty(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 any) *MockEngineCalcDifficultyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalcDifficulty", reflect.TypeOf((*MockEngine)(nil).CalcDifficulty), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	return &MockEngineCalcDifficultyCall{Call: call}
}

// MockEngineCalcDifficultyCall wrap *gomock.Call
type MockEngineCalcDifficultyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineCalcDifficultyCall) Return(arg0 *big.Int) *MockEngineCalcDifficultyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineCalcDifficultyCall) Do(f func(ChainHeaderReader, uint64, uint64, *big.Int, uint64, common.Hash, common.Hash, uint64) *big.Int) *MockEngineCalcDifficultyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineCalcDifficultyCall) DoAndReturn(f func(ChainHeaderReader, uint64, uint64, *big.Int, uint64, common.Hash, common.Hash, uint64) *big.Int) *MockEngineCalcDifficultyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// CalculateRewards mocks base method.
func (m *MockEngine) CalculateRewards(arg0 *chain.Config, arg1 *types.Header, arg2 []*types.Header, arg3 SystemCall) ([]Reward, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateRewards", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]Reward)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CalculateRewards indicates an expected call of CalculateRewards.
func (mr *MockEngineMockRecorder) CalculateRewards(arg0, arg1, arg2, arg3 any) *MockEngineCalculateRewardsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateRewards", reflect.TypeOf((*MockEngine)(nil).CalculateRewards), arg0, arg1, arg2, arg3)
	return &MockEngineCalculateRewardsCall{Call: call}
}

// MockEngineCalculateRewardsCall wrap *gomock.Call
type MockEngineCalculateRewardsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineCalculateRewardsCall) Return(arg0 []Reward, arg1 error) *MockEngineCalculateRewardsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineCalculateRewardsCall) Do(f func(*chain.Config, *types.Header, []*types.Header, SystemCall) ([]Reward, error)) *MockEngineCalculateRewardsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineCalculateRewardsCall) DoAndReturn(f func(*chain.Config, *types.Header, []*types.Header, SystemCall) ([]Reward, error)) *MockEngineCalculateRewardsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Close mocks base method.
func (m *MockEngine) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockEngineMockRecorder) Close() *MockEngineCloseCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEngine)(nil).Close))
	return &MockEngineCloseCall{Call: call}
}

// MockEngineCloseCall wrap *gomock.Call
type MockEngineCloseCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineCloseCall) Return(arg0 error) *MockEngineCloseCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineCloseCall) Do(f func() error) *MockEngineCloseCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineCloseCall) DoAndReturn(f func() error) *MockEngineCloseCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Finalize mocks base method.
func (m *MockEngine) Finalize(arg0 *chain.Config, arg1 *types.Header, arg2 *state.IntraBlockState, arg3 types.Transactions, arg4 []*types.Header, arg5 types.Receipts, arg6 []*types.Withdrawal, arg7 ChainReader, arg8 SystemCall, arg9 log.Logger) (types.Transactions, types.Receipts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Finalize", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	ret0, _ := ret[0].(types.Transactions)
	ret1, _ := ret[1].(types.Receipts)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Finalize indicates an expected call of Finalize.
func (mr *MockEngineMockRecorder) Finalize(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9 any) *MockEngineFinalizeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Finalize", reflect.TypeOf((*MockEngine)(nil).Finalize), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9)
	return &MockEngineFinalizeCall{Call: call}
}

// MockEngineFinalizeCall wrap *gomock.Call
type MockEngineFinalizeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineFinalizeCall) Return(arg0 types.Transactions, arg1 types.Receipts, arg2 error) *MockEngineFinalizeCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineFinalizeCall) Do(f func(*chain.Config, *types.Header, *state.IntraBlockState, types.Transactions, []*types.Header, types.Receipts, []*types.Withdrawal, ChainReader, SystemCall, log.Logger) (types.Transactions, types.Receipts, error)) *MockEngineFinalizeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineFinalizeCall) DoAndReturn(f func(*chain.Config, *types.Header, *state.IntraBlockState, types.Transactions, []*types.Header, types.Receipts, []*types.Withdrawal, ChainReader, SystemCall, log.Logger) (types.Transactions, types.Receipts, error)) *MockEngineFinalizeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// FinalizeAndAssemble mocks base method.
func (m *MockEngine) FinalizeAndAssemble(arg0 *chain.Config, arg1 *types.Header, arg2 *state.IntraBlockState, arg3 types.Transactions, arg4 []*types.Header, arg5 types.Receipts, arg6 []*types.Withdrawal, arg7 ChainReader, arg8 SystemCall, arg9 Call, arg10 log.Logger) (*types.Block, types.Transactions, types.Receipts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinalizeAndAssemble", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	ret0, _ := ret[0].(*types.Block)
	ret1, _ := ret[1].(types.Transactions)
	ret2, _ := ret[2].(types.Receipts)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// FinalizeAndAssemble indicates an expected call of FinalizeAndAssemble.
func (mr *MockEngineMockRecorder) FinalizeAndAssemble(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10 any) *MockEngineFinalizeAndAssembleCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinalizeAndAssemble", reflect.TypeOf((*MockEngine)(nil).FinalizeAndAssemble), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
	return &MockEngineFinalizeAndAssembleCall{Call: call}
}

// MockEngineFinalizeAndAssembleCall wrap *gomock.Call
type MockEngineFinalizeAndAssembleCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineFinalizeAndAssembleCall) Return(arg0 *types.Block, arg1 types.Transactions, arg2 types.Receipts, arg3 error) *MockEngineFinalizeAndAssembleCall {
	c.Call = c.Call.Return(arg0, arg1, arg2, arg3)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineFinalizeAndAssembleCall) Do(f func(*chain.Config, *types.Header, *state.IntraBlockState, types.Transactions, []*types.Header, types.Receipts, []*types.Withdrawal, ChainReader, SystemCall, Call, log.Logger) (*types.Block, types.Transactions, types.Receipts, error)) *MockEngineFinalizeAndAssembleCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineFinalizeAndAssembleCall) DoAndReturn(f func(*chain.Config, *types.Header, *state.IntraBlockState, types.Transactions, []*types.Header, types.Receipts, []*types.Withdrawal, ChainReader, SystemCall, Call, log.Logger) (*types.Block, types.Transactions, types.Receipts, error)) *MockEngineFinalizeAndAssembleCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GenerateSeal mocks base method.
func (m *MockEngine) GenerateSeal(arg0 ChainHeaderReader, arg1, arg2 *types.Header, arg3 Call) []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateSeal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GenerateSeal indicates an expected call of GenerateSeal.
func (mr *MockEngineMockRecorder) GenerateSeal(arg0, arg1, arg2, arg3 any) *MockEngineGenerateSealCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateSeal", reflect.TypeOf((*MockEngine)(nil).GenerateSeal), arg0, arg1, arg2, arg3)
	return &MockEngineGenerateSealCall{Call: call}
}

// MockEngineGenerateSealCall wrap *gomock.Call
type MockEngineGenerateSealCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineGenerateSealCall) Return(arg0 []byte) *MockEngineGenerateSealCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineGenerateSealCall) Do(f func(ChainHeaderReader, *types.Header, *types.Header, Call) []byte) *MockEngineGenerateSealCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineGenerateSealCall) DoAndReturn(f func(ChainHeaderReader, *types.Header, *types.Header, Call) []byte) *MockEngineGenerateSealCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Initialize mocks base method.
func (m *MockEngine) Initialize(arg0 *chain.Config, arg1 ChainHeaderReader, arg2 *types.Header, arg3 *state.IntraBlockState, arg4 SysCallCustom, arg5 log.Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Initialize", arg0, arg1, arg2, arg3, arg4, arg5)
}

// Initialize indicates an expected call of Initialize.
func (mr *MockEngineMockRecorder) Initialize(arg0, arg1, arg2, arg3, arg4, arg5 any) *MockEngineInitializeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockEngine)(nil).Initialize), arg0, arg1, arg2, arg3, arg4, arg5)
	return &MockEngineInitializeCall{Call: call}
}

// MockEngineInitializeCall wrap *gomock.Call
type MockEngineInitializeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineInitializeCall) Return() *MockEngineInitializeCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineInitializeCall) Do(f func(*chain.Config, ChainHeaderReader, *types.Header, *state.IntraBlockState, SysCallCustom, log.Logger)) *MockEngineInitializeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineInitializeCall) DoAndReturn(f func(*chain.Config, ChainHeaderReader, *types.Header, *state.IntraBlockState, SysCallCustom, log.Logger)) *MockEngineInitializeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// IsServiceTransaction mocks base method.
func (m *MockEngine) IsServiceTransaction(arg0 common.Address, arg1 SystemCall) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsServiceTransaction", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsServiceTransaction indicates an expected call of IsServiceTransaction.
func (mr *MockEngineMockRecorder) IsServiceTransaction(arg0, arg1 any) *MockEngineIsServiceTransactionCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsServiceTransaction", reflect.TypeOf((*MockEngine)(nil).IsServiceTransaction), arg0, arg1)
	return &MockEngineIsServiceTransactionCall{Call: call}
}

// MockEngineIsServiceTransactionCall wrap *gomock.Call
type MockEngineIsServiceTransactionCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineIsServiceTransactionCall) Return(arg0 bool) *MockEngineIsServiceTransactionCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineIsServiceTransactionCall) Do(f func(common.Address, SystemCall) bool) *MockEngineIsServiceTransactionCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineIsServiceTransactionCall) DoAndReturn(f func(common.Address, SystemCall) bool) *MockEngineIsServiceTransactionCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Prepare mocks base method.
func (m *MockEngine) Prepare(arg0 ChainHeaderReader, arg1 *types.Header, arg2 *state.IntraBlockState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prepare", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prepare indicates an expected call of Prepare.
func (mr *MockEngineMockRecorder) Prepare(arg0, arg1, arg2 any) *MockEnginePrepareCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prepare", reflect.TypeOf((*MockEngine)(nil).Prepare), arg0, arg1, arg2)
	return &MockEnginePrepareCall{Call: call}
}

// MockEnginePrepareCall wrap *gomock.Call
type MockEnginePrepareCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEnginePrepareCall) Return(arg0 error) *MockEnginePrepareCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEnginePrepareCall) Do(f func(ChainHeaderReader, *types.Header, *state.IntraBlockState) error) *MockEnginePrepareCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEnginePrepareCall) DoAndReturn(f func(ChainHeaderReader, *types.Header, *state.IntraBlockState) error) *MockEnginePrepareCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Seal mocks base method.
func (m *MockEngine) Seal(arg0 ChainHeaderReader, arg1 *types.Block, arg2 chan<- *types.Block, arg3 <-chan struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Seal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Seal indicates an expected call of Seal.
func (mr *MockEngineMockRecorder) Seal(arg0, arg1, arg2, arg3 any) *MockEngineSealCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Seal", reflect.TypeOf((*MockEngine)(nil).Seal), arg0, arg1, arg2, arg3)
	return &MockEngineSealCall{Call: call}
}

// MockEngineSealCall wrap *gomock.Call
type MockEngineSealCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineSealCall) Return(arg0 error) *MockEngineSealCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineSealCall) Do(f func(ChainHeaderReader, *types.Block, chan<- *types.Block, <-chan struct{}) error) *MockEngineSealCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineSealCall) DoAndReturn(f func(ChainHeaderReader, *types.Block, chan<- *types.Block, <-chan struct{}) error) *MockEngineSealCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SealHash mocks base method.
func (m *MockEngine) SealHash(arg0 *types.Header) common.Hash {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SealHash", arg0)
	ret0, _ := ret[0].(common.Hash)
	return ret0
}

// SealHash indicates an expected call of SealHash.
func (mr *MockEngineMockRecorder) SealHash(arg0 any) *MockEngineSealHashCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SealHash", reflect.TypeOf((*MockEngine)(nil).SealHash), arg0)
	return &MockEngineSealHashCall{Call: call}
}

// MockEngineSealHashCall wrap *gomock.Call
type MockEngineSealHashCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineSealHashCall) Return(arg0 common.Hash) *MockEngineSealHashCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineSealHashCall) Do(f func(*types.Header) common.Hash) *MockEngineSealHashCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineSealHashCall) DoAndReturn(f func(*types.Header) common.Hash) *MockEngineSealHashCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Type mocks base method.
func (m *MockEngine) Type() chain.ConsensusName {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(chain.ConsensusName)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockEngineMockRecorder) Type() *MockEngineTypeCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockEngine)(nil).Type))
	return &MockEngineTypeCall{Call: call}
}

// MockEngineTypeCall wrap *gomock.Call
type MockEngineTypeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineTypeCall) Return(arg0 chain.ConsensusName) *MockEngineTypeCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineTypeCall) Do(f func() chain.ConsensusName) *MockEngineTypeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineTypeCall) DoAndReturn(f func() chain.ConsensusName) *MockEngineTypeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// VerifyHeader mocks base method.
func (m *MockEngine) VerifyHeader(arg0 ChainHeaderReader, arg1 *types.Header, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyHeader", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyHeader indicates an expected call of VerifyHeader.
func (mr *MockEngineMockRecorder) VerifyHeader(arg0, arg1, arg2 any) *MockEngineVerifyHeaderCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyHeader", reflect.TypeOf((*MockEngine)(nil).VerifyHeader), arg0, arg1, arg2)
	return &MockEngineVerifyHeaderCall{Call: call}
}

// MockEngineVerifyHeaderCall wrap *gomock.Call
type MockEngineVerifyHeaderCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineVerifyHeaderCall) Return(arg0 error) *MockEngineVerifyHeaderCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineVerifyHeaderCall) Do(f func(ChainHeaderReader, *types.Header, bool) error) *MockEngineVerifyHeaderCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineVerifyHeaderCall) DoAndReturn(f func(ChainHeaderReader, *types.Header, bool) error) *MockEngineVerifyHeaderCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// VerifyUncles mocks base method.
func (m *MockEngine) VerifyUncles(arg0 ChainReader, arg1 *types.Header, arg2 []*types.Header) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUncles", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUncles indicates an expected call of VerifyUncles.
func (mr *MockEngineMockRecorder) VerifyUncles(arg0, arg1, arg2 any) *MockEngineVerifyUnclesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUncles", reflect.TypeOf((*MockEngine)(nil).VerifyUncles), arg0, arg1, arg2)
	return &MockEngineVerifyUnclesCall{Call: call}
}

// MockEngineVerifyUnclesCall wrap *gomock.Call
type MockEngineVerifyUnclesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockEngineVerifyUnclesCall) Return(arg0 error) *MockEngineVerifyUnclesCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockEngineVerifyUnclesCall) Do(f func(ChainReader, *types.Header, []*types.Header) error) *MockEngineVerifyUnclesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockEngineVerifyUnclesCall) DoAndReturn(f func(ChainReader, *types.Header, []*types.Header) error) *MockEngineVerifyUnclesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
