package mock

import (
	"context"
	"github.com/5GCoreNet/5GCoreNetSDK/fivegc"
	openapinlmfbroadcast "github.com/5GCoreNet/server-openapi/Nlmf_Broadcast"
)

// BroadcastMock is a mock of the Broadcast interface
type BroadcastMock struct {
	cipherResponseData openapinlmfbroadcast.CipherResponseData
	problemDetails     openapinlmfbroadcast.ProblemDetails
	redirectResponse   fivegc.RedirectResponse
	statusCode         fivegc.StatusCode
}

// NewBroadcastMock creates a new mock of the Broadcast interface
func NewBroadcastMock(
	cipherResponseData openapinlmfbroadcast.CipherResponseData,
	problemDetails openapinlmfbroadcast.ProblemDetails,
	redirectResponse fivegc.RedirectResponse,
	statusCode fivegc.StatusCode,
) *BroadcastMock {
	return &BroadcastMock{
		cipherResponseData: cipherResponseData,
		problemDetails:     problemDetails,
		redirectResponse:   redirectResponse,
		statusCode:         statusCode,
	}
}

// ProblemDetails allows to set the problem details of the mock
func (b *BroadcastMock) ProblemDetails(problemDetails openapinlmfbroadcast.ProblemDetails) {
	b.problemDetails = problemDetails
}

// CipherResponseData allows to set the cipher response data of the mock
func (b *BroadcastMock) CipherResponseData(cipherResponseData openapinlmfbroadcast.CipherResponseData) {
	b.cipherResponseData = cipherResponseData
}

// RedirectResponse allows to set the redirect response of the mock
func (b *BroadcastMock) RedirectResponse(redirectResponse fivegc.RedirectResponse) {
	b.redirectResponse = redirectResponse
}

// StatusCode allows to set the status code of the mock
func (b *BroadcastMock) StatusCode(statusCode fivegc.StatusCode) {
	b.statusCode = statusCode
}

// Error returns the problem details of the mock
func (b *BroadcastMock) Error(ctx context.Context, err error) openapinlmfbroadcast.ProblemDetails {
	return b.problemDetails
}

// CipherKeyData returns the cipher response data, the problem details, the redirect response and the status code of the mock
func (b *BroadcastMock) CipherKeyData(ctx context.Context, data openapinlmfbroadcast.CipherRequestData) (openapinlmfbroadcast.CipherResponseData, openapinlmfbroadcast.ProblemDetails, fivegc.RedirectResponse, fivegc.StatusCode) {
	return b.cipherResponseData, b.problemDetails, b.redirectResponse, b.statusCode
}
