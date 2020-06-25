package cfn

import (
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
)

// EmptyHandler is a implementation of Handler
//
// This implementation of the handlers is only used for testing.
type EmptyHandler struct{}

func (h *EmptyHandler) Create(request handler.Request) handler.ProgressEvent {
	return handler.ProgressEvent{}
}

func (h *EmptyHandler) Read(request handler.Request) handler.ProgressEvent {
	return handler.ProgressEvent{}
}

func (h *EmptyHandler) Update(request handler.Request) handler.ProgressEvent {
	return handler.ProgressEvent{}
}

func (h *EmptyHandler) Delete(request handler.Request) handler.ProgressEvent {
	return handler.ProgressEvent{}
}

func (h *EmptyHandler) List(request handler.Request) handler.ProgressEvent {
	return handler.ProgressEvent{}
}

// MockHandler is a implementation of Handler
//
// This implementation of the handlers is only used for testing.
type MockHandler struct {
	fn func(callback map[string]interface{}, s *session.Session) handler.ProgressEvent
}

func (m *MockHandler) Create(request handler.Request) handler.ProgressEvent {
	return m.fn(request.CallbackContext, request.Session)
}

func (m *MockHandler) Read(request handler.Request) handler.ProgressEvent {
	return m.fn(request.CallbackContext, request.Session)
}

func (m *MockHandler) Update(request handler.Request) handler.ProgressEvent {
	return m.fn(request.CallbackContext, request.Session)
}

func (m *MockHandler) Delete(request handler.Request) handler.ProgressEvent {
	return m.fn(request.CallbackContext, request.Session)
}

func (m *MockHandler) List(request handler.Request) handler.ProgressEvent {
	return m.fn(request.CallbackContext, request.Session)
}

//MockedMetrics mocks the call to AWS CloudWatch Metrics
//
// This implementation of the handlers is only used for testing.
type MockedMetrics struct {
	cloudwatchiface.CloudWatchAPI
	ResourceTypeName               string
	HandlerExceptionCount          int
	HandlerInvocationDurationCount int
	HandlerInvocationCount         int
}

//NewMockedMetrics is a factory function that returns a new MockedMetrics.
//
// This implementation of the handlers is only used for testing.
func NewMockedMetrics() *MockedMetrics {
	return &MockedMetrics{}
}

//PutMetricData mocks the PutMetricData method.
//
// This implementation of the handlers is only used for testing.
func (m *MockedMetrics) PutMetricData(in *cloudwatch.PutMetricDataInput) (*cloudwatch.PutMetricDataOutput, error) {
	m.ResourceTypeName = *in.Namespace
	d := in.MetricData[0].MetricName
	switch *d {
	case "HandlerException":
		m.HandlerExceptionCount++
	case "HandlerInvocationDuration":
		m.HandlerInvocationDurationCount++
	case "HandlerInvocationCount":
		m.HandlerInvocationCount++
	}

	return nil, nil
}

// MockModel mocks a resource model
//
// This implementation of the handlers is only used for testing.
type MockModel struct {
	Property1 *string `json:"property1,omitempty"`
	Property2 *string `json:"property2,omitempty"`
}

// MockModelHandler is a implementation of Handler
//
// This implementation of the handlers is only used for testing.
type MockModelHandler struct {
	fn func(r handler.Request) handler.ProgressEvent
}

func (m *MockModelHandler) Create(request handler.Request) handler.ProgressEvent {
	return m.fn(request)
}

func (m *MockModelHandler) Read(request handler.Request) handler.ProgressEvent {
	return m.fn(request)
}

func (m *MockModelHandler) Update(request handler.Request) handler.ProgressEvent {
	return m.fn(request)
}

func (m *MockModelHandler) Delete(request handler.Request) handler.ProgressEvent {
	return m.fn(request)
}

func (m *MockModelHandler) List(request handler.Request) handler.ProgressEvent {
	return m.fn(request)
}
