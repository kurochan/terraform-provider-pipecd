// Code generated by MockGen. DO NOT EDIT.
// Source: provider.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	apiservice "github.com/pipe-cd/pipecd/pkg/app/server/service/apiservice"
	grpc "google.golang.org/grpc"
)

// MockAPIClient is a mock of APIClient interface.
type MockAPIClient struct {
	ctrl     *gomock.Controller
	recorder *MockAPIClientMockRecorder
}

// MockAPIClientMockRecorder is the mock recorder for MockAPIClient.
type MockAPIClientMockRecorder struct {
	mock *MockAPIClient
}

// NewMockAPIClient creates a new mock instance.
func NewMockAPIClient(ctrl *gomock.Controller) *MockAPIClient {
	mock := &MockAPIClient{ctrl: ctrl}
	mock.recorder = &MockAPIClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIClient) EXPECT() *MockAPIClientMockRecorder {
	return m.recorder
}

// AddApplication mocks base method.
func (m *MockAPIClient) AddApplication(ctx context.Context, in *apiservice.AddApplicationRequest, opts ...grpc.CallOption) (*apiservice.AddApplicationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddApplication", varargs...)
	ret0, _ := ret[0].(*apiservice.AddApplicationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddApplication indicates an expected call of AddApplication.
func (mr *MockAPIClientMockRecorder) AddApplication(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddApplication", reflect.TypeOf((*MockAPIClient)(nil).AddApplication), varargs...)
}

// DeleteApplication mocks base method.
func (m *MockAPIClient) DeleteApplication(ctx context.Context, in *apiservice.DeleteApplicationRequest, opts ...grpc.CallOption) (*apiservice.DeleteApplicationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteApplication", varargs...)
	ret0, _ := ret[0].(*apiservice.DeleteApplicationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteApplication indicates an expected call of DeleteApplication.
func (mr *MockAPIClientMockRecorder) DeleteApplication(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteApplication", reflect.TypeOf((*MockAPIClient)(nil).DeleteApplication), varargs...)
}

// DisableApplication mocks base method.
func (m *MockAPIClient) DisableApplication(ctx context.Context, in *apiservice.DisableApplicationRequest, opts ...grpc.CallOption) (*apiservice.DisableApplicationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisableApplication", varargs...)
	ret0, _ := ret[0].(*apiservice.DisableApplicationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableApplication indicates an expected call of DisableApplication.
func (mr *MockAPIClientMockRecorder) DisableApplication(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableApplication", reflect.TypeOf((*MockAPIClient)(nil).DisableApplication), varargs...)
}

// DisablePiped mocks base method.
func (m *MockAPIClient) DisablePiped(ctx context.Context, in *apiservice.DisablePipedRequest, opts ...grpc.CallOption) (*apiservice.DisablePipedResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisablePiped", varargs...)
	ret0, _ := ret[0].(*apiservice.DisablePipedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisablePiped indicates an expected call of DisablePiped.
func (mr *MockAPIClientMockRecorder) DisablePiped(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisablePiped", reflect.TypeOf((*MockAPIClient)(nil).DisablePiped), varargs...)
}

// EnableApplication mocks base method.
func (m *MockAPIClient) EnableApplication(ctx context.Context, in *apiservice.EnableApplicationRequest, opts ...grpc.CallOption) (*apiservice.EnableApplicationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableApplication", varargs...)
	ret0, _ := ret[0].(*apiservice.EnableApplicationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableApplication indicates an expected call of EnableApplication.
func (mr *MockAPIClientMockRecorder) EnableApplication(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableApplication", reflect.TypeOf((*MockAPIClient)(nil).EnableApplication), varargs...)
}

// EnablePiped mocks base method.
func (m *MockAPIClient) EnablePiped(ctx context.Context, in *apiservice.EnablePipedRequest, opts ...grpc.CallOption) (*apiservice.EnablePipedResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnablePiped", varargs...)
	ret0, _ := ret[0].(*apiservice.EnablePipedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnablePiped indicates an expected call of EnablePiped.
func (mr *MockAPIClientMockRecorder) EnablePiped(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnablePiped", reflect.TypeOf((*MockAPIClient)(nil).EnablePiped), varargs...)
}

// Encrypt mocks base method.
func (m *MockAPIClient) Encrypt(ctx context.Context, in *apiservice.EncryptRequest, opts ...grpc.CallOption) (*apiservice.EncryptResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Encrypt", varargs...)
	ret0, _ := ret[0].(*apiservice.EncryptResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encrypt indicates an expected call of Encrypt.
func (mr *MockAPIClientMockRecorder) Encrypt(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encrypt", reflect.TypeOf((*MockAPIClient)(nil).Encrypt), varargs...)
}

// GetApplication mocks base method.
func (m *MockAPIClient) GetApplication(ctx context.Context, in *apiservice.GetApplicationRequest, opts ...grpc.CallOption) (*apiservice.GetApplicationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetApplication", varargs...)
	ret0, _ := ret[0].(*apiservice.GetApplicationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApplication indicates an expected call of GetApplication.
func (mr *MockAPIClientMockRecorder) GetApplication(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApplication", reflect.TypeOf((*MockAPIClient)(nil).GetApplication), varargs...)
}

// GetCommand mocks base method.
func (m *MockAPIClient) GetCommand(ctx context.Context, in *apiservice.GetCommandRequest, opts ...grpc.CallOption) (*apiservice.GetCommandResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCommand", varargs...)
	ret0, _ := ret[0].(*apiservice.GetCommandResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommand indicates an expected call of GetCommand.
func (mr *MockAPIClientMockRecorder) GetCommand(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommand", reflect.TypeOf((*MockAPIClient)(nil).GetCommand), varargs...)
}

// GetDeployment mocks base method.
func (m *MockAPIClient) GetDeployment(ctx context.Context, in *apiservice.GetDeploymentRequest, opts ...grpc.CallOption) (*apiservice.GetDeploymentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDeployment", varargs...)
	ret0, _ := ret[0].(*apiservice.GetDeploymentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployment indicates an expected call of GetDeployment.
func (mr *MockAPIClientMockRecorder) GetDeployment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployment", reflect.TypeOf((*MockAPIClient)(nil).GetDeployment), varargs...)
}

// GetPiped mocks base method.
func (m *MockAPIClient) GetPiped(ctx context.Context, in *apiservice.GetPipedRequest, opts ...grpc.CallOption) (*apiservice.GetPipedResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPiped", varargs...)
	ret0, _ := ret[0].(*apiservice.GetPipedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPiped indicates an expected call of GetPiped.
func (mr *MockAPIClientMockRecorder) GetPiped(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPiped", reflect.TypeOf((*MockAPIClient)(nil).GetPiped), varargs...)
}

// GetPlanPreviewResults mocks base method.
func (m *MockAPIClient) GetPlanPreviewResults(ctx context.Context, in *apiservice.GetPlanPreviewResultsRequest, opts ...grpc.CallOption) (*apiservice.GetPlanPreviewResultsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPlanPreviewResults", varargs...)
	ret0, _ := ret[0].(*apiservice.GetPlanPreviewResultsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlanPreviewResults indicates an expected call of GetPlanPreviewResults.
func (mr *MockAPIClientMockRecorder) GetPlanPreviewResults(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlanPreviewResults", reflect.TypeOf((*MockAPIClient)(nil).GetPlanPreviewResults), varargs...)
}

// ListApplications mocks base method.
func (m *MockAPIClient) ListApplications(ctx context.Context, in *apiservice.ListApplicationsRequest, opts ...grpc.CallOption) (*apiservice.ListApplicationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListApplications", varargs...)
	ret0, _ := ret[0].(*apiservice.ListApplicationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListApplications indicates an expected call of ListApplications.
func (mr *MockAPIClientMockRecorder) ListApplications(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListApplications", reflect.TypeOf((*MockAPIClient)(nil).ListApplications), varargs...)
}

// ListDeployments mocks base method.
func (m *MockAPIClient) ListDeployments(ctx context.Context, in *apiservice.ListDeploymentsRequest, opts ...grpc.CallOption) (*apiservice.ListDeploymentsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDeployments", varargs...)
	ret0, _ := ret[0].(*apiservice.ListDeploymentsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeployments indicates an expected call of ListDeployments.
func (mr *MockAPIClientMockRecorder) ListDeployments(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeployments", reflect.TypeOf((*MockAPIClient)(nil).ListDeployments), varargs...)
}

// ListStageLogs mocks base method.
func (m *MockAPIClient) ListStageLogs(ctx context.Context, in *apiservice.ListStageLogsRequest, opts ...grpc.CallOption) (*apiservice.ListStageLogsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStageLogs", varargs...)
	ret0, _ := ret[0].(*apiservice.ListStageLogsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStageLogs indicates an expected call of ListStageLogs.
func (mr *MockAPIClientMockRecorder) ListStageLogs(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStageLogs", reflect.TypeOf((*MockAPIClient)(nil).ListStageLogs), varargs...)
}

// RegisterEvent mocks base method.
func (m *MockAPIClient) RegisterEvent(ctx context.Context, in *apiservice.RegisterEventRequest, opts ...grpc.CallOption) (*apiservice.RegisterEventResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterEvent", varargs...)
	ret0, _ := ret[0].(*apiservice.RegisterEventResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterEvent indicates an expected call of RegisterEvent.
func (mr *MockAPIClientMockRecorder) RegisterEvent(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterEvent", reflect.TypeOf((*MockAPIClient)(nil).RegisterEvent), varargs...)
}

// RegisterPiped mocks base method.
func (m *MockAPIClient) RegisterPiped(ctx context.Context, in *apiservice.RegisterPipedRequest, opts ...grpc.CallOption) (*apiservice.RegisterPipedResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RegisterPiped", varargs...)
	ret0, _ := ret[0].(*apiservice.RegisterPipedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterPiped indicates an expected call of RegisterPiped.
func (mr *MockAPIClientMockRecorder) RegisterPiped(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterPiped", reflect.TypeOf((*MockAPIClient)(nil).RegisterPiped), varargs...)
}

// RenameApplicationConfigFile mocks base method.
func (m *MockAPIClient) RenameApplicationConfigFile(ctx context.Context, in *apiservice.RenameApplicationConfigFileRequest, opts ...grpc.CallOption) (*apiservice.RenameApplicationConfigFileResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RenameApplicationConfigFile", varargs...)
	ret0, _ := ret[0].(*apiservice.RenameApplicationConfigFileResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenameApplicationConfigFile indicates an expected call of RenameApplicationConfigFile.
func (mr *MockAPIClientMockRecorder) RenameApplicationConfigFile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenameApplicationConfigFile", reflect.TypeOf((*MockAPIClient)(nil).RenameApplicationConfigFile), varargs...)
}

// RequestPlanPreview mocks base method.
func (m *MockAPIClient) RequestPlanPreview(ctx context.Context, in *apiservice.RequestPlanPreviewRequest, opts ...grpc.CallOption) (*apiservice.RequestPlanPreviewResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RequestPlanPreview", varargs...)
	ret0, _ := ret[0].(*apiservice.RequestPlanPreviewResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestPlanPreview indicates an expected call of RequestPlanPreview.
func (mr *MockAPIClientMockRecorder) RequestPlanPreview(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestPlanPreview", reflect.TypeOf((*MockAPIClient)(nil).RequestPlanPreview), varargs...)
}

// SyncApplication mocks base method.
func (m *MockAPIClient) SyncApplication(ctx context.Context, in *apiservice.SyncApplicationRequest, opts ...grpc.CallOption) (*apiservice.SyncApplicationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncApplication", varargs...)
	ret0, _ := ret[0].(*apiservice.SyncApplicationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncApplication indicates an expected call of SyncApplication.
func (mr *MockAPIClientMockRecorder) SyncApplication(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncApplication", reflect.TypeOf((*MockAPIClient)(nil).SyncApplication), varargs...)
}

// UpdateApplication mocks base method.
func (m *MockAPIClient) UpdateApplication(ctx context.Context, in *apiservice.UpdateApplicationRequest, opts ...grpc.CallOption) (*apiservice.UpdateApplicationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateApplication", varargs...)
	ret0, _ := ret[0].(*apiservice.UpdateApplicationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateApplication indicates an expected call of UpdateApplication.
func (mr *MockAPIClientMockRecorder) UpdateApplication(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateApplication", reflect.TypeOf((*MockAPIClient)(nil).UpdateApplication), varargs...)
}

// UpdatePiped mocks base method.
func (m *MockAPIClient) UpdatePiped(ctx context.Context, in *apiservice.UpdatePipedRequest, opts ...grpc.CallOption) (*apiservice.UpdatePipedResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdatePiped", varargs...)
	ret0, _ := ret[0].(*apiservice.UpdatePipedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePiped indicates an expected call of UpdatePiped.
func (mr *MockAPIClientMockRecorder) UpdatePiped(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePiped", reflect.TypeOf((*MockAPIClient)(nil).UpdatePiped), varargs...)
}
