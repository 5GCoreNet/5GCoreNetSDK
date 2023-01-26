package mock

import (
	"context"
	"github.com/5GCoreNet/5GCoreNetSDK/fivegc"
	openapinlmflocation "github.com/5GCoreNet/server-openapi/Nlmf_Location"
)

// LocationMock is a mock of the Location interface
type LocationMock struct {
	locationData     openapinlmflocation.LocationData
	problemDetails   openapinlmflocation.ProblemDetails
	redirectResponse fivegc.RedirectResponse
	statusCode       fivegc.StatusCode
}

// NewLocationMock creates a new mock of the Location interface
func NewLocationMock(
	locationData openapinlmflocation.LocationData,
	problemDetails openapinlmflocation.ProblemDetails,
	redirectResponse fivegc.RedirectResponse,
	statusCode fivegc.StatusCode,
) *LocationMock {
	return &LocationMock{
		locationData:     locationData,
		problemDetails:   problemDetails,
		redirectResponse: redirectResponse,
		statusCode:       statusCode,
	}
}

// ProblemDetails allows to set the problem details of the mock
func (l *LocationMock) ProblemDetails(problemDetails openapinlmflocation.ProblemDetails) {
	l.problemDetails = problemDetails
}

// LocationData allows to set the location data of the mock
func (l *LocationMock) LocationData(locationData openapinlmflocation.LocationData) {
	l.locationData = locationData
}

// RedirectResponse allows to set the redirect response of the mock
func (l *LocationMock) RedirectResponse(redirectResponse fivegc.RedirectResponse) {
	l.redirectResponse = redirectResponse
}

// StatusCode allows to set the status code of the mock
func (l *LocationMock) StatusCode(statusCode fivegc.StatusCode) {
	l.statusCode = statusCode
}

// Error returns the problem details of the mock
func (l *LocationMock) Error(ctx context.Context, err error) openapinlmflocation.ProblemDetails {
	return l.problemDetails
}

// CancelLocation returns the problem details, the redirect response and the status code of the mock
func (l *LocationMock) CancelLocation(ctx context.Context, data openapinlmflocation.CancelLocData) (openapinlmflocation.ProblemDetails, fivegc.RedirectResponse, fivegc.StatusCode) {
	return l.problemDetails, l.redirectResponse, l.statusCode
}

// DetermineLocation returns the location data, the problem details, the redirect response and the status code of the mock
func (l *LocationMock) DetermineLocation(ctx context.Context, data openapinlmflocation.InputData) (openapinlmflocation.LocationData, openapinlmflocation.ProblemDetails, fivegc.RedirectResponse, fivegc.StatusCode) {
	return l.locationData, l.problemDetails, l.redirectResponse, l.statusCode
}

// LocationContextTransfer returns the problem details, the redirect response and the status code of the mock
func (l *LocationMock) LocationContextTransfer(ctx context.Context, data openapinlmflocation.LocContextData) (openapinlmflocation.ProblemDetails, fivegc.RedirectResponse, fivegc.StatusCode) {
	return l.problemDetails, l.redirectResponse, l.statusCode
}
