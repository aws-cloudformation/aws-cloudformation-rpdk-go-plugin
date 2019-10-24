package callback

import (
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
)

//MockedEvents mocks the call to AWS CloudWatch Events
type MockedCallback struct {
	cloudformationiface.CloudFormationAPI
	errCount int
}

func NewMockedCallback(errCount int) *MockedCallback {
	return &MockedCallback{
		errCount: errCount,
	}
}

func (m *MockedCallback) RecordHandlerProgress(in *cloudformation.RecordHandlerProgressInput) (*cloudformation.RecordHandlerProgressOutput, error) {

	if m.errCount > 0 {
		m.errCount--
		return nil, errors.New("error")
	}

	return nil, nil
}

func TestTranslateOperationStatus(t *testing.T) {
	type args struct {
		operationStatus string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestSUCCESS", args{"SUCCESS"}, cloudformation.OperationStatusSuccess},
		{"TestFAILED", args{"FAILED"}, cloudformation.OperationStatusFailed},
		{"TestIN_PROGRESS", args{"IN_PROGRESS"}, cloudformation.OperationStatusInProgress},
		{"TestFoo", args{"Foo"}, cloudformation.OperationStatusFailed},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TranslateOperationStatus(tt.args.operationStatus); got != tt.want {
				t.Errorf("TranslateOperationStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranslateErrorCode(t *testing.T) {
	type args struct {
		errorCode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestNotUpdatable", args{"NotUpdatable"}, cloudformation.HandlerErrorCodeNotUpdatable},
		{"TestInvalidRequest", args{"InvalidRequest"}, cloudformation.HandlerErrorCodeInvalidRequest},
		{"AccessDenied", args{"AccessDenied"}, cloudformation.HandlerErrorCodeAccessDenied},
		{"TestInvalidCredentials", args{"InvalidCredentials"}, cloudformation.HandlerErrorCodeInvalidCredentials},
		{"TestAlreadyExists", args{"AlreadyExists"}, cloudformation.HandlerErrorCodeAlreadyExists},
		{"TestNotFound", args{"NotFound"}, cloudformation.HandlerErrorCodeNotFound},
		{"TestResourceConflict", args{"ResourceConflict"}, cloudformation.HandlerErrorCodeResourceConflict},
		{"TestThrottling", args{"Throttling"}, cloudformation.HandlerErrorCodeThrottling},
		{"TestServiceLimitExceeded", args{"ServiceLimitExceeded"}, cloudformation.HandlerErrorCodeServiceLimitExceeded},
		{"TestGeneralServiceException", args{"GeneralServiceException"}, cloudformation.HandlerErrorCodeGeneralServiceException},
		{"TestServiceInternalError", args{"ServiceInternalError"}, cloudformation.HandlerErrorCodeServiceInternalError},
		{"TestNetworkFailure", args{"NetworkFailure"}, cloudformation.HandlerErrorCodeNetworkFailure},
		{"TestFoo", args{"foo"}, cloudformation.HandlerErrorCodeInternalFailure},
		{"TestInternalFailure", args{"InternalFailure"}, cloudformation.HandlerErrorCodeInternalFailure},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TranslateErrorCode(tt.args.errorCode); got != tt.want {
				t.Errorf("TranslateErrorCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCloudFormationCallbackAdapter_ReportProgress(t *testing.T) {
	type fields struct {
		client cloudformationiface.CloudFormationAPI
	}
	type args struct {
		bearerToken   string
		code          string
		status        string
		resourceModel interface{}
		statusMessage string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"TestRetryMaxReturnErr", fields{NewMockedCallback(6)}, args{"123456", "ACCESSDENIED", "FAILED", "test", "retry"}, true},
		{"TestRetryReturnNoErr", fields{NewMockedCallback(0)}, args{"123456", "ACCESSDENIED", "FAILED", "test", "retry"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CloudFormationCallbackAdapter{
				client: tt.fields.client,
			}
			if err := c.ReportProgress(tt.args.bearerToken, tt.args.code, tt.args.status, tt.args.resourceModel, tt.args.statusMessage); (err != nil) != tt.wantErr {
				t.Errorf("CloudFormationCallbackAdapter.ReportProgress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
